package iamserver

import (
	"context"
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"math/rand"
	"oceanbolt.com/obapi/internal/iam/dao"
	"oceanbolt.com/obapi/rpc/iam"
	"strings"
	"time"
)

// CreateKey inserts new apikey in backend db
func (s *Server) CreateKey(ctx context.Context, req *iam.CreateKeyRequest) (key *iam.UserKeyWithSecret, err error) {
	db := dao.IamDAO{Ctx: ctx, Fs: s.Fs}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(s.Config.JWKS_RS256_PRIVATE_KEY))
	if err != nil {
		return key, err
	}

	obkid := getKeyId(16)
	signKeyBytePrint := md5.Sum([]byte(s.Config.JWKS_RS256_PRIVATE_KEY))
	signKeyStringPrint := fmt.Sprintf("%x", signKeyBytePrint)

	log.Println("Checking if keys exists")
	if ok := db.CheckIfPublicKeyExistsFS(signKeyStringPrint); !ok {
		log.Println("Key did not exist. inserting...")
		publicKey, _ := extractPublicKeyFromPrivate(privateKey)
		db.InsertPublicKeyFS(&iam.PublicKey{
			Kid:       signKeyStringPrint,
			KeyEnv:    s.Config.OBENV,
			PublicKey: publicKey,
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"aud":   "https://api.oceanbolt.com",
		"iss":   "https://oceanbolt.eu.auth0.com/",
		"ktype": "apikey",
		"obkid": obkid,
		"sub":   req.UserId,
		"iat":   time.Now().Unix(),
		"exp":   time.Unix(req.Expires, 0),
		"kid":   signKeyStringPrint,
	})
	token.Header["kid"] = signKeyStringPrint

	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("Token could not be signed")
	}
	log.Printf("Inserting key")
	err = db.InsertKeyFS(&iam.UserKey{
		ApikeyId:     obkid,
		Expires:      req.Expires,
		UserId:       req.UserId,
		SigningKeyId: signKeyStringPrint,
		KeyTag:       req.KeyTag,
	})
	if err != nil {
		return key, err
	}

	return &iam.UserKeyWithSecret{
		Expires:      req.Expires,
		ApikeyId:     obkid,
		KeyTag:       req.KeyTag,
		ApikeySecret: signedToken,
		UserId:       req.UserId,
	}, nil
}

// getKeyId generates random key id string
func getKeyId(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")

	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

// extractPublicKeyFromPrivate extracts the public key from an rsa private key
func extractPublicKeyFromPrivate(pk *rsa.PrivateKey) ([]byte, error) {

	publicKeyDer, _ := x509.MarshalPKIXPublicKey(pk.Public())
	publicKeyBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyDer,
	}
	return pem.EncodeToMemory(&publicKeyBlock), nil

}
