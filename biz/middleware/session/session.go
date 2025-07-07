package session

import (
	"auth/biz/db/redis"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"github.com/rbcervilla/redisstore/v9"
)

const (
	sessionStorePrefix = "auth_session:"
	sessionName        = "auth_session_id"
	cookiePath         = "/"
	cookieDomain       = ""
	cookieMaxAge       = 7 * 24 * 3600
	cookieSecure       = false
	cookieHttpOnly     = true
	cookieSameSite     = http.SameSiteStrictMode
)

func New() app.HandlerFunc {
	store := NewRedisStore()
	store.Options(sessions.Options{
		Path:     cookiePath,
		Domain:   cookieDomain,
		MaxAge:   cookieMaxAge,
		Secure:   cookieSecure,
		HttpOnly: cookieHttpOnly,
		SameSite: cookieSameSite,
	})

	return sessions.New(sessionName, store)
}

func Remove(c *app.RequestContext) error {
	sess := sessions.Default(c)
	sess.Options(sessions.Options{
		Path:     cookiePath,
		Domain:   cookieDomain,
		MaxAge:   -1,
		Secure:   cookieSecure,
		HttpOnly: cookieHttpOnly,
		SameSite: cookieSameSite,
	})
	return sess.Save()
}

type RedisStore struct {
	*redisstore.RedisStore
}

func (r *RedisStore) Options(opts sessions.Options) {
	r.RedisStore.Options(*opts.ToGorillaOptions())
}

func NewRedisStore() *RedisStore {
	redisStore, err := redisstore.NewRedisStore(context.Background(), redis.GetRedisClient())
	if err != nil {
		panic(err)
	}
	redisStore.KeyPrefix(sessionStorePrefix)
	return &RedisStore{
		RedisStore: redisStore,
	}
}
