package jwt

import (
	"auth/biz/util/random"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {

	secret := "secret"
	tokenId := random.RandStr(10)
	sessId := random.RandStr(10)

	jwtStr, err := generateToken(Payload{
		UserID: random.RandStr(32),
	}, time.Second, tokenId, sessId, secret, "go test")
	assert.Nil(t, err)
	t.Log(jwtStr)

	t.Run("success", func(t *testing.T) {
		claims, err := validateToken(jwtStr, secret)
		assert.Nil(t, err)
		assert.Equal(t, sessId, claims.Subject)

		assert.True(t, claims.CheckSum(sessId))
	})

	t.Run("secret key invalid", func(t *testing.T) {
		_, err := validateToken(jwtStr, secret+"123")
		assert.ErrorIs(t, ErrJwtInvalid, err)
	})

	t.Run("expired", func(t *testing.T) {
		time.Sleep(time.Second * 2)
		_, err := validateToken(jwtStr, secret)
		assert.ErrorIs(t, ErrJwtExpired, err)
	})

}
