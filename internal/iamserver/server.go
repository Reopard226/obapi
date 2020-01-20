package iamserver

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"math/rand"
	"oceanbolt.com/iamservice/internal/config"
	"oceanbolt.com/iamservice/internal/dao"
	pb "oceanbolt.com/iamservice/rpc/iam"
	"strings"
	"time"
)

type Server struct {
	Db *mongo.Database
	Config *config.Config
}

func (s *Server) ListKeys(ctx context.Context, user *pb.User) (keys *pb.UserKeys, err error) {
	db := dao.IamDAO{Ctx: ctx, Db: s.Db,}

	return db.ListKeys(user)
}

func (s *Server) DeleteKey(ctx context.Context, req *pb.DeleteKeyRequest) (resp *pb.KeyDeletedResponse, err error) {
	db := dao.IamDAO{Ctx: ctx, Db: s.Db,}

	err = db.DeleteKey(req)
	if err != nil {
		return resp, err
	}

	resp = &pb.KeyDeletedResponse{
		Message:"Key '" + req.ApikeyId + "' successfully deleted",
	}
	return resp, nil
}

func (s *Server) CreateKey(ctx context.Context, req *pb.CreateKeyRequest) (key *pb.UserKeyWithSecret, err error) {
	db := dao.IamDAO{Ctx: ctx, Db: s.Db,}

	parsedKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(s.Config.JWKS_RS256_PRIVATE_KEY))
	if err != nil {
		return key, err
	}

	obkid := getKeyId(16)
	signKeyBytePrint := md5.Sum([]byte(s.Config.JWKS_RS256_PRIVATE_KEY))
	signKeyStringPrint := fmt.Sprintf("%x",signKeyBytePrint)
	
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"aud":   "https://api.oceanbolt.com",
		"iss":   "https://oceanbolt.eu.auth0.com/",
		"ktype": "apikey",
		"obkid": obkid,
		"sub":   req.UserId,
		"iat":   time.Now().Unix(),
		"exp":   time.Unix(req.Expires,0),
	})
	token.Header["kid"] = signKeyStringPrint

	signedToken, err := token.SignedString(parsedKey)
	if err != nil {
		log.Fatal("Token could not be signed")
	}
	
	err = db.InsertKey(&pb.UserKey{
		ApikeyId:     obkid,
		Expires:      req.Expires,
		UserId:       req.UserId,
		SigningKeyId: signKeyStringPrint,
		KeyTag:       req.KeyTag,
	})
	if err != nil {
		return key, err
	}

	return &pb.UserKeyWithSecret{
		Expires:      req.Expires,
		ApikeyId:     obkid,
		KeyTag:       req.KeyTag,
		ApikeySecret: signedToken,
		UserId:       req.UserId,
	},nil
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