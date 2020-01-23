package iamserver

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"regexp"
	"testing"
	"time"
)

func TestGetKeyId(t *testing.T) {
	keyid := getKeyId(10)

	r, _ := regexp.MatchString("[A-z0-9]{10}", keyid)

	if !r {
		t.Errorf("Error getkeyid function not working")
	}
}


func TestExtractPublicKeyFromPrivate(t *testing.T) {

	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(RSATestPrivateKey))
	if err != nil {
		log.Fatal(err)
	}

	testToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"aud":   "test",
		"iss":   "test",
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour*1),
	})

	signedToken, _ := testToken.SignedString(key)

	parsedPublicKeyBytes, _ := extractPublicKeyFromPrivate(key)

	publicKey, _ := jwt.ParseRSAPublicKeyFromPEM(parsedPublicKeyBytes)

	_, err = jwt.Parse(signedToken, func(token *jwt.Token) (i interface{}, err error) {
		return publicKey,nil
	})
	if !(err == nil) {
		t.Errorf("Public certificate failed")
	}

}

const RSATestPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`
