package jwttool

import (
	"time"

	"github.com/lestrrat-go/jwx/jwt"
)

// GenerateJWT is to generate JWT.
// using ES256.
// However, this function cannot generate the whole JWT(include head, payload and sign).
// Only head and payload. Sign part need jwt.Sign to append.
// After calling jwt.Sign, Sign function return []byte(have processed through base64).
// Just need to string(what jwt.Sign return).
func GenerateJWT(info map[string]interface{}) (token jwt.Token) {
	token = jwt.New()

	// set payload in RFC-7519
	token.Set(jwt.IssuerKey, "arapgp")
	token.Set(jwt.IssuedAtKey, time.Now())
	token.Set(jwt.ExpirationKey, time.Now().Add(time.Hour))

	// set extra payload k-v pair
	// unsafe operation
	for k, v := range info {
		token.Set(k, v)
	}

	return token
}
