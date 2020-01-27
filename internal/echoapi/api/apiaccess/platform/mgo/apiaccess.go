package mgo

import (
	"context"

	"oceanbolt.com/obapi/internal/iam/iamclient"
	"oceanbolt.com/obapi/rpc/iam"
)

// Apiaccess represents the client for Apiaccess table
type Apiaccess struct{}

// NewApiaccess returns a new apiaccess database instance
func NewApiaccess() *Apiaccess {
	return &Apiaccess{}
}

// ListKey returns list of apiaccess keys
func (u *Apiaccess) ListKey(client iamclient.OceanboltIAMClient) (*iam.UserKeys, error) {
	return client.ListKeys(context.Background(), &iam.User{UserId: "auth0|5d89314ad227b10d5b5c1c88"})
}

// CreateKey creates new api access key
func (u *Apiaccess) CreateKey(client iamclient.OceanboltIAMClient, tag string, exp int64) (*iam.UserKeyWithSecret, error) {
	return client.CreateKey(context.Background(), &iam.CreateKeyRequest{
		Expires: exp,
		KeyTag:  tag,
		UserId:  "auth0|5d89314ad227b10d5b5c1c88",
	})
}

// DeleteKey delete one api access key
func (u *Apiaccess) DeleteKey(client iamclient.OceanboltIAMClient, apikeyID string) (*iam.KeyDeletedResponse, error) {
	return client.DeleteKey(context.Background(), &iam.DeleteKeyRequest{
		ApikeyId: apikeyID,
		UserId:   "auth0|5d89314ad227b10d5b5c1c88",
	})
}
