package dao

import (
	"context"
	"errors"

	"log"
	"oceanbolt.com/obapi/rpc/iam"
	"strings"

	"cloud.google.com/go/firestore"
)

const APIKEY_COLLECTION_NAME = "apikeys"
const USER_COLLECTION_NAME = "users"
const PUBLIC_KEY_COLLECTION_NAME = "public_keys"

// NewFireStoreDatabase initializes a new firestore db client
func NewFireStoreDatabase(ctx context.Context, projectID string) (*firestore.Client, error) {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return client, err
	}

	return client, nil

}

// ListKeysFS gets all the keys for a given user from the firestore db
func (fs *IamDAO) ListKeysFS(user *iam.User) (*iam.UserKeys, error) {

	r, err := fs.Fs.Collection(USER_COLLECTION_NAME).Doc(user.UserId).Collection(APIKEY_COLLECTION_NAME).Documents(fs.Ctx).GetAll()
	if err != nil {
		return &iam.UserKeys{}, err
	}
	structResult := make([]*iam.UserKey, len(r))
	for k, v := range r {
		err = v.DataTo(&structResult[k])
		if err != nil {
			return &iam.UserKeys{}, err
		}
	}
	return &iam.UserKeys{NumberOfKeys: int64(len(structResult)), Keys: structResult}, nil
}

// InsertKeyFS inserts a new apikey the firestore db
func (fs *IamDAO) InsertKeyFS(key *iam.UserKey) error {
	_, err := fs.Fs.Collection(USER_COLLECTION_NAME).Doc(key.UserId).Collection(APIKEY_COLLECTION_NAME).Doc(key.ApikeyId).Create(fs.Ctx, key)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

// CheckIfApikeyExistsFS deletes an apikey if it exists in the firestore db
func (fs *IamDAO) DeleteKeyFS(key *iam.DeleteKeyRequest) error {

	r, err := fs.Fs.Collection(USER_COLLECTION_NAME).Doc(key.UserId).Collection(APIKEY_COLLECTION_NAME).Doc(key.ApikeyId).Get(fs.Ctx)

	if !r.Exists() {
		return errors.New("No key exists with apikey_id '" + key.ApikeyId + "'")
	}

	_, err = fs.Fs.Collection(USER_COLLECTION_NAME).Doc(key.UserId).Collection(APIKEY_COLLECTION_NAME).Doc(key.ApikeyId).Delete(fs.Ctx)
	if err != nil {
		return err
	}
	return nil
}

// CheckIfApikeyExistsFS checks if an apikey exists in the firestore db
func (fs *IamDAO) CheckIfApikeyExistsFS(key *iam.UserKey) (bool, error) {
	_, err := fs.Fs.Collection(USER_COLLECTION_NAME).Doc(key.UserId).Collection(APIKEY_COLLECTION_NAME).Doc(key.ApikeyId).Get(fs.Ctx)
	if err != nil {
		if strings.HasPrefix(err.Error(), "rpc error: code = NotFound") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}


// GetPublicKeyFS gets the public key for a given privatekey id from the firestore db
func (fs *IamDAO) GetPublicKeyFS(user *iam.PrivateKey) (pub *iam.PublicKey, err error) {

	r, err := fs.Fs.Collection(PUBLIC_KEY_COLLECTION_NAME).Doc(user.Kid).Get(fs.Ctx)
	if err != nil {
		return pub, err
	}
	err = r.DataTo(&pub)
	return pub, err
}

// CheckIfPublicKeyExistsFS checks if a public key exists in the firestore db
func (fs *IamDAO) CheckIfPublicKeyExistsFS(kid string) bool {
	r, _ := fs.Fs.Collection(PUBLIC_KEY_COLLECTION_NAME).Doc(kid).Get(fs.Ctx)
	if !r.Exists() {
		return false
	}
	return true
}

// InsertPublicKey inserts a Oceanbolt public in the firestore db
func (fs *IamDAO) InsertPublicKeyFS(pub *iam.PublicKey) error {
	_, err := fs.Fs.Collection(PUBLIC_KEY_COLLECTION_NAME).Doc(pub.Kid).Create(fs.Ctx, pub)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
