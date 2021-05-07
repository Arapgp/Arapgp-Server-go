package session

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/Arapgp/Arapgp-Server-go/pkg/shatool"
)

// GenerateSession is a tool function
// Alg Now: sha256(input+tm+rand(Int))
func GenerateSession(input string) (session string) {
	tm := time.Now().String()
	rd := strconv.Itoa(rand.Int())
	return shatool.Sha256String(input + tm + rd)
}
