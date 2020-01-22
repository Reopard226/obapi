package iamclient

import (
	"cloud.google.com/go/compute/metadata"
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	"log"
	"net/http"
	"oceanbolt.com/iamservice/rpc/iam"
	"os"
)

type OceanboltIAMClient iam.Apikey

const IAMServiceURL = "https://iamserver-cu5jmh4vyq-ew.a.run.app"

type AddHeaderTransport struct {
	T http.RoundTripper
}

type Runtime string

const (
	CLOUDRUN Runtime = "CloudRun"
	LOCAL            = "Local"
)

func GetDefaultIamClient() OceanboltIAMClient {
	client := iam.NewApikeyProtobufClient(IAMServiceURL, &http.Client{Transport: NewAddHeaderTransport(nil)})
	return client
}

func (adt *AddHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {

	var RUNTIME Runtime
	//RUNTIME := os.Getenv("K_SERVICE")
	if os.Getenv("K_SERVICE") != "" {
		RUNTIME = CLOUDRUN
	} else {
		RUNTIME = LOCAL
	}

	token := GetDefaultToken(RUNTIME)

	req.Header.Add("Authorization", "Bearer "+token)
	return adt.T.RoundTrip(req)
}

func NewAddHeaderTransport(T http.RoundTripper) *AddHeaderTransport {
	if T == nil {
		T = http.DefaultTransport
	}
	return &AddHeaderTransport{T}
}

func GetDefaultTokenDEV() string {
	ctx := context.Background()
	audience := IAMServiceURL
	creds, err := transport.Creds(ctx, option.WithScopes(audience))
	if err != nil {
		log.Fatal(err)
	}
	token, err := creds.TokenSource.Token()
	if err != nil {
		log.Fatal(err)
	}
	return token.Extra("id_token").(string)
}

func GetDefaultTokenPROD() string {
	tokenURL := fmt.Sprintf("/instance/service-accounts/default/identity?audience=%s", IAMServiceURL)
	idToken, err := metadata.Get(tokenURL)
	if err != nil {
		log.Fatal(err)
	}
	return idToken
}

func GetDefaultToken(runtime Runtime) string {
	if runtime == CLOUDRUN {
		return GetDefaultTokenPROD()
	} else {
		return GetDefaultTokenDEV()
	}
}
