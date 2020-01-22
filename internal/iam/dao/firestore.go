package dao

import (
	"context"
	"errors"

	"log"
	"oceanbolt.com/obapi/rpc/iam"
	"strings"

	"cloud.google.com/go/firestore"
)

func NewFireStoreDatabase(ctx context.Context, projectID string) (*firestore.Client, error) {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return client, err
	}

	return client, nil

}

func (fs *IamDAO) ListKeysFS(user *iam.User) (*iam.UserKeys, error) {

	r, err := fs.Fs.Collection(APIKEY_COLLECTION_NAME).Where("user_id", "==", user.UserId).Documents(fs.Ctx).GetAll()
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

func (fs *IamDAO) InsertKeyFS(key *iam.UserKey) error {
	_, err := fs.Fs.Collection(APIKEY_COLLECTION_NAME).Doc(key.UserId+"|"+key.ApikeyId).Create(fs.Ctx, key)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (fs *IamDAO) DeleteKeyFS(key *iam.DeleteKeyRequest) error {

	r, err := fs.Fs.Collection(APIKEY_COLLECTION_NAME).Doc(key.UserId + "|" + key.ApikeyId).Get(fs.Ctx)

	if !r.Exists() {
		return errors.New("No key exists with apikey_id '" + key.ApikeyId + "'")
	}

	_, err = fs.Fs.Collection(APIKEY_COLLECTION_NAME).Doc(key.UserId + "|" + key.ApikeyId).Delete(fs.Ctx)
	if err != nil {
		return err
	}
	return nil
}

func (fs *IamDAO) CheckKeyFS(key *iam.UserKey) (bool, error) {
	_, err := fs.Fs.Collection(APIKEY_COLLECTION_NAME).Doc(key.UserId + "|" + key.ApikeyId).Get(fs.Ctx)
	if err != nil {
		if strings.HasPrefix(err.Error(), "rpc error: code = NotFound") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
