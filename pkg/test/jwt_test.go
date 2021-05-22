package test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"log"
	"testing"

	ijwt "github.com/Arapgp/Arapgp-Server-go/internal/jwt"
	"github.com/Arapgp/Arapgp-Server-go/pkg/jwttool"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/stretchr/testify/assert"
)

func Test_JWT(t *testing.T) {
	ast := assert.New(t)

	// generate jwt payload
	info := map[string]interface{}{
		"username": "ljg",
		"email":    "ljg@outlook.com",
	}
	token := jwttool.GenerateJWT(info)

	_, err := json.Marshal(token)
	ast.Empty(err)

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	ast.Empty(err)

	// generate
	tokenb, err := jwt.Sign(token, jwa.ES256, *key)
	ast.Empty(err)
	log.Println("tokenb:", string(tokenb))

	buf, err := jws.Verify(tokenb, jwa.ES256, key.PublicKey)
	ast.Empty(err)
	log.Println("buf:", string(buf))

	// json.Unmarshal
	jwtInstance := &ijwt.JwtArab{}
	json.Unmarshal(buf, jwtInstance)
	log.Println("username:", jwtInstance.Username, "; email:", jwtInstance.Email)
}
