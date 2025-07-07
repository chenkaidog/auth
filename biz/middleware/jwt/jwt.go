package jwt

import (
	"auth/biz/config"
	rediscli "auth/biz/db/redis"
	"auth/biz/model/dto"
	"auth/biz/model/errs"
	"auth/biz/util/encode"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/hertz-contrib/sessions"
	"github.com/redis/go-redis/v9"
)

var (
	ErrUnexpectedJwtMethod = errors.New("unexpected jwt method")
	ErrJwtInvalid          = errors.New("jwt is invalid")
	ErrJwtExpired          = errors.New("jwt is expired")
)

func ValidateMW() app.HandlerFunc {
	jwtConf := config.GetJWTConfig()

	return func(ctx context.Context, c *app.RequestContext) {
		jwtStr := exactJWT(c)
		if jwtStr == "" {
			hlog.CtxInfof(ctx, "authorization failed, token is empty")
			dto.AbortWithErr(c, errs.Unauthorized, http.StatusUnauthorized)
			return
		}

		// 0. basic validation
		claims, err := validateToken(jwtStr, jwtConf.SecretKey)
		if err != nil {
			hlog.CtxInfof(ctx, "jwt invalid: %v", err)
			dto.AbortWithErr(c, errs.Unauthorized, http.StatusUnauthorized)
			return
		}

		// 1. check the summary of session id
		sess := sessions.Default(c)
		if !claims.CheckSum(sess.ID()) {
			hlog.CtxInfof(ctx, "session not match")
			dto.AbortWithErr(c, errs.Unauthorized, http.StatusUnauthorized)
			return
		}

		// 2. check the existance of token id
		if exist, err := rediscli.GetRedisClient().
			Get(ctx, tokenExistKey(claims.ID)).Bool(); err != nil && !errors.Is(err, redis.Nil) {
			hlog.CtxErrorf(ctx, "redis get err: %v", err)
			dto.AbortWithErr(c, errs.ServerError, http.StatusInternalServerError)
			return
		} else if !exist {
			hlog.CtxInfof(ctx, "jwt token invalid or expired")
			dto.AbortWithErr(c, errs.Unauthorized, http.StatusUnauthorized)
			return
		}

		// set claims
		ctx = context.WithValue(ctx, Payload{}, claims)

		c.Next(ctx)
	}
}

type Payload struct {
	UserID string `json:"user_id,omitempty"`
	Name   string `json:"name,omitempty"`
}

type Claims struct {
	jwt.RegisteredClaims
	Payload

	Sum string `json:"sum,omitempty"`
}

func (c *Claims) SessSum() {
	c.Sum = encode.EncodePassword(c.ID, c.Subject)
}

func (c *Claims) CheckSum(sessID string) bool {
	return encode.EncodePassword(c.ID, sessID) == c.Sum
}

func GenerateToken(ctx context.Context, c *app.RequestContext, payload Payload, sessID string) (string, int64, error) {
	tokenID := uuid.New().String()

	jwtConf := config.GetJWTConfig()
	exp := time.Duration(jwtConf.Expiration) * time.Second
	expAt := time.Now().Add(exp).Unix()

	jwtStr, err := generateToken(payload, exp, tokenID, sessID, jwtConf.SecretKey, jwtConf.Issuer)
	if err != nil {
		return "", 0, err
	}

	if err := rediscli.GetRedisClient().
		Set(ctx, tokenExistKey(tokenID), true, exp).Err(); err != nil {
		hlog.CtxErrorf(ctx, "cache token id err: %v", err)
		return "", 0, err
	}

	return jwtStr, expAt, nil
}

func GetPayload(ctx context.Context) Payload {
	claims, ok := ctx.Value(Payload{}).(*Claims)
	if ok {
		return claims.Payload
	}
	return Payload{}
}

func RemoveToken(ctx context.Context) error {
	if claims, ok := ctx.Value(Payload{}).(*Claims); ok {
		return rediscli.GetRedisClient().Del(ctx, tokenExistKey(claims.ID)).Err()
	}

	return nil
}

func generateToken(payload Payload, expiration time.Duration, tokenID, sessID, secret, issuer string) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			Issuer:    issuer,
			Subject:   sessID,
			ID:        tokenID,
		},
		Payload: payload,
	}
	claims.SessSum()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func validateToken(tokenStr, secret string) (*Claims, error) {
	var claims Claims
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrHashUnavailable
		}
		return []byte(secret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrHashUnavailable) {
			return nil, ErrUnexpectedJwtMethod
		}
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrJwtExpired
		}
		if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return nil, ErrJwtInvalid
		}
		return nil, err
	}
	if !token.Valid {
		return nil, ErrJwtInvalid
	}

	return &claims, nil
}

func tokenExistKey(tid string) string {
	return fmt.Sprintf("jwt_id_exist:%s", tid)
}

func exactJWT(c *app.RequestContext) string {
	authToken := c.Request.Header.Get("Authorization") // "Bearer token123"
	tokenLs := strings.Split(authToken, " ")
	if len(tokenLs) > 1 {
		return tokenLs[1]
	}

	return ""
}
