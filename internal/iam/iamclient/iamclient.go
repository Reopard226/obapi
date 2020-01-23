package iamclient

import (
	"cloud.google.com/go/compute/metadata"
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	"log"
	"net/http"
	"oceanbolt.com/obapi/rpc/iam"
	"os"
)

type OceanboltIAMClient iam.Apikey

const IAMServiceURL = "https://iamserver-cu5jmh4vyq-ew.a.run.app"

type AddHeaderTransport struct {
	T http.RoundTripper
	ServiceUrl string
}

type Runtime string

const (
	CLOUDRUN Runtime = "CloudRun"
	LOCAL            = "Local"
)

func GetDefaultIamClient(iamUrl string) OceanboltIAMClient {
	client := iam.NewApikeyProtobufClient(iamUrl, &http.Client{Transport: NewAddHeaderTransport(nil, iamUrl)})
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

	token := GetDefaultToken(RUNTIME,adt.ServiceUrl)

	req.Header.Add("Authorization", "Bearer "+token)
	return adt.T.RoundTrip(req)
}

func NewAddHeaderTransport(T http.RoundTripper, iamUrl string) *AddHeaderTransport {
	if T == nil {
		T = http.DefaultTransport
	}
	return &AddHeaderTransport{T,iamUrl}
}

func GetDefaultTokenDEV(serviceUrl string) string {
	ctx := context.Background()
	audience := serviceUrl
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

func GetDefaultTokenPROD(serviceUrl string) string {
	tokenURL := fmt.Sprintf("/instance/service-accounts/default/identity?audience=%s", serviceUrl)
	idToken, err := metadata.Get(tokenURL)
	if err != nil {
		log.Fatal(err)
	}
	return idToken
}

func GetDefaultToken(runtime Runtime, serviceUrl string) string {
	if runtime == CLOUDRUN {
		return GetDefaultTokenPROD(serviceUrl)
	} else {
		return GetDefaultTokenDEV(serviceUrl)
	}
}
