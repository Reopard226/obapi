package iamserver

import (
	"context"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"oceanbolt.com/obapi/internal/iam/dao"
	"oceanbolt.com/obapi/rpc/iam"
	"time"
)

const Auth0KeyUrl = "https://oceanbolt.eu.auth0.com/.well-known/jwks.json"

func (s *Server) GetPublicKey(ctx context.Context, pk *iam.PrivateKey) (pub *iam.PublicKey, err error) {
	db := dao.IamDAO{Ctx: ctx, Db: s.Db, Fs: s.Fs}

	key, err := ParseJWKSAuth0(pk.Kid)
	if err == nil {
		return &iam.PublicKey{
			Kid:       pk.Kid,
			KeyEnv:    "prod",
			PublicKey: PadKeyToByte(key),
		}, nil
	}

	return db.GetPublicKeyFS(pk)

}

func ParseJWKSAuth0(kid string) (key string, err error) {

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	response, err := netClient.Get(Auth0KeyUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	jwtKeys := &JWTKeys{}
	err = jsoniter.Unmarshal(body, jwtKeys)
	if err != nil {
		panic(err)
	}

	for _, v := range jwtKeys.Keys {
		if v.Kid == kid {
			return v.X5c[0], nil
		}
	}

	return "", errors.New("Error - cloud not find public key for the given 'kid'")

}

func PadKeyToByte(key string) []byte {
	bytes := []byte("-----BEGIN CERTIFICATE-----\n" + key + "\n-----END CERTIFICATE-----")
	return bytes
}

type JWTKey struct {
	Alg string   `json:"alg"`
	Kty string   `json:"kty"`
	Use string   `json:"use"`
	X5c []string `json:"x5c"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	Kid string   `json:"kid"`
	X5t string   `json:"x5t"`
}

type JWTKeys struct {
	Keys []JWTKey `json:"keys"`
}
