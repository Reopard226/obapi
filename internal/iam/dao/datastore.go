package dao

import (
	"cloud.google.com/go/datastore"
	"context"
	"errors"
	"fmt"
	"log"
	"oceanbolt.com/obapi/rpc/iam"
)

const APIKEYLIST_ENTITY_NAME = "apikeys"
const USER_KIND = "User"
const USER_APIKEYLIST_KIND = "UserApikeyList"
const PUBLIC_KEY_KIND = "PublicRSAKey"
const APIKEY_KIND = "UserApikey"

type IamDAO struct {
	Ctx context.Context
	Ds  *datastore.Client
}

// NewDataStoreDatabase initializes a new firestore db client
func NewDataStoreDatabase(ctx context.Context, projectID string) (*datastore.Client, error) {
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return client, err
	}

	return client, nil

}

// ListKeysDS gets all the keys for a given user from the firestore db
func (fs *IamDAO) ListKeysDS(user *iam.User) (*iam.UserKeys, error) {
	userKey := datastore.NameKey("User", user.UserId, nil)

	q := datastore.NewQuery("UserApikey").Ancestor(userKey)
	res := make([]*iam.UserKey, 0)
	_, err := fs.Ds.GetAll(fs.Ctx, q, &res)
	if err != nil {
		return &iam.UserKeys{}, err
	}

	return &iam.UserKeys{NumberOfKeys: int64(len(res)), Keys: res}, nil
}

// InsertKeyDS inserts a new apikey the firestore db
func (fs *IamDAO) InsertKeyDS(key *iam.UserKey) error {
	userKey := datastore.NameKey("User", key.UserId, nil)
	keylistKey := datastore.NameKey("UserApikeyList", "apikeys", userKey)
	apikeyKey := datastore.NameKey("UserApikey", key.ApikeyId, keylistKey)
	keyr, err := fs.Ds.Put(fs.Ctx, apikeyKey, key)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("%v\n", keyr.String())

	return nil
}

// DeleteKeyDS deletes an apikey if it exists in the firestore db
func (fs *IamDAO) DeleteKeyDS(key *iam.DeleteKeyRequest) error {
	userKey := datastore.NameKey("User", key.UserId, nil)
	keylistKey := datastore.NameKey("UserApikeyList", "apikeys", userKey)
	apikeyKey := datastore.NameKey("UserApikey", key.ApikeyId, keylistKey)
	err := fs.Ds.Delete(fs.Ctx, apikeyKey)
	if err != nil {
		return errors.New("No key exists with apikey_id '" + key.ApikeyId + "'")
	}

	return nil
}

// CheckIfApikeyExistsDS checks if an apikey exists in the firestore db
func (fs *IamDAO) CheckIfApikeyExistsDS(key *iam.UserKey) bool {
	userKey := datastore.NameKey("User", key.UserId, nil)
	keylistKey := datastore.NameKey("UserApikeyList", "apikeys", userKey)
	apikeyKey := datastore.NameKey("UserApikey", key.ApikeyId, keylistKey)

	//q := datastore.NewQuery("UserApikey").Ancestor(keylistKey).Filter("apikey_id=",key.ApikeyId).KeysOnly()
	keyOut := new(iam.UserKey)
	err := fs.Ds.Get(fs.Ctx, apikeyKey, keyOut)
	if err != nil {
		return false
	}
	//return c != 0, nil
	return true
}

// GetPublicKeyDS gets the public key for a given privatekey id from the firestore db
func (fs *IamDAO) GetPublicKeyDS(pk *iam.PrivateKey) (pub *iam.PublicKey, err error) {
	userKey := datastore.NameKey("PublicRSAKey", pk.Kid, nil)

	err = fs.Ds.Get(fs.Ctx, userKey, &pub)
	if err != nil {
		return pub, err
	}
	return pub, err
}

// CheckIfPublicKeyExistsDS checks if a public key exists in the firestore db
func (fs *IamDAO) CheckIfPublicKeyExistsDS(kid string) bool {
	pubkeyKey := datastore.NameKey("PublicRSAKey", kid, nil)

	keyOut := new(iam.PublicKey)
	err := fs.Ds.Get(fs.Ctx, pubkeyKey, &keyOut)
	if err != nil {
		return false
	}

	return true
}

// InsertPublicKey inserts a Oceanbolt public in the firestore db
func (fs *IamDAO) InsertPublicKeyDS(pub *iam.PublicKey) error {
	pubkeyKey := datastore.NameKey("PublicRSAKey", pub.Kid, nil)

	_, err := fs.Ds.Put(fs.Ctx, pubkeyKey, pub)
	if err != nil {
		return err
	}
	return nil
}
