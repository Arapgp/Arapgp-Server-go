package jwttool

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"

	log "github.com/sirupsen/logrus"
)

// EncodeKey2String can encode pri/pub key to string
func EncodeKey2String(pri *ecdsa.PrivateKey, pub *ecdsa.PublicKey) (priKeyStr, pubKeyStr string) {
	priKey, err := x509.MarshalECPrivateKey(pri)
	if err != nil {
		log.WithFields(log.Fields{
			"priKey": priKey,
			"err":    err.Error(),
		}).Warningln("arapgp.pkg.jwttool EncodeKey2String x509.MarshalECPrivateKey failed")
		return
	}

	pubKey, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		log.WithFields(log.Fields{
			"pubKey": pubKey,
			"err":    err.Error(),
		}).Warningln("arapgp.pkg.jwttool EncodeKey2String x509.MarshalPKIXPublicKey failed")
		return
	}

	priKeyStr = string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: priKey}))
	pubKeyStr = string(pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pubKey}))
	return
}

// DecodeString2Key can decode string to pri/pub key
func DecodeString2Key(pemPri, pemPub string) (priKey *ecdsa.PrivateKey, pubKey *ecdsa.PublicKey) {
	blk, _ := pem.Decode([]byte(pemPri))
	x509Encoded := blk.Bytes
	priKey, err := x509.ParseECPrivateKey(x509Encoded)
	if err != nil {
		log.WithFields(log.Fields{
			"priKey": priKey,
			"err":    err.Error(),
		}).Warningln("arapgp.pkg.jwttool DecodeString2Key x509.ParseECPrivateKey failed")
		return
	}

	blk, _ = pem.Decode([]byte(pemPub))
	x509Encoded = blk.Bytes
	tmp, err := x509.ParsePKIXPublicKey(x509Encoded)
	if err != nil {
		log.WithFields(log.Fields{
			"pubKeyTmp": tmp,
			"err":       err.Error(),
		}).Warningln("arapgp.pkg.jwttool DecodeString2Key x509.ParsePKIXPublicKey failed")
		return
	}
	pubKey = tmp.(*ecdsa.PublicKey)
	return
}
