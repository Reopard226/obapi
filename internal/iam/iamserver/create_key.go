package iamserver

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"math/rand"
	"oceanbolt.com/iamservice/internal/iam/dao"
	"oceanbolt.com/iamservice/rpc/iam"
	"strings"
	"time"
)

func (s *Server) CreateKey(ctx context.Context, req *iam.CreateKeyRequest) (key *iam.UserKeyWithSecret, err error) {
	db := dao.MgoDao{Ctx: ctx, Db: s.Db}

	parsedKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(s.Config.JWKS_RS256_PRIVATE_KEY))
	if err != nil {
		return key, err
	}

	obkid := getKeyId(16)
	signKeyBytePrint := md5.Sum([]byte(s.Config.JWKS_RS256_PRIVATE_KEY))
	signKeyStringPrint := fmt.Sprintf("%x", signKeyBytePrint)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"aud":   "https://api.oceanbolt.com",
		"iss":   "https://oceanbolt.eu.auth0.com/",
		"ktype": "apikey",
		"obkid": obkid,
		"sub":   req.UserId,
		"iat":   time.Now().Unix(),
		"exp":   time.Unix(req.Expires, 0),
	})
	token.Header["kid"] = signKeyStringPrint

	signedToken, err := token.SignedString(parsedKey)
	if err != nil {
		log.Fatal("Token could not be signed")
	}

	err = db.InsertKey(&iam.UserKey{
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
