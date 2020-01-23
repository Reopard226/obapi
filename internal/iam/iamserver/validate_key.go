package iamserver

import (
	"context"
	"oceanbolt.com/obapi/internal/iam/dao"
	pb "oceanbolt.com/obapi/rpc/iam"
	"time"
)

type PermissionCache struct {
	Valid       bool
	Permissions []string
	Expires     time.Time
}

//var PermissionDict map[string]PermissionCache

// ValidateKey validates that an access token exists in backend db and is valid
func (s *Server) ValidateKey(ctx context.Context, key *pb.UserKey) (resp *pb.ValidationResponse, err error) {
	db := dao.IamDAO{Ctx: ctx, Fs: s.Fs}

	/*initMap := make(map[string]PermissionCache)

	if s.PermissionCache == nil {
		s.PermissionCache = &initMap
	}

	if existingPermissions, ok := PermissionDict[key.UserId]; ok {
		if existingPermissions.Expires.After(time.Now()) {
			log.Println("Permission ok: returning cached permissions")
			return &pb.ValidationResponse{
				Valid:       existingPermissions.Valid,
				Permissions: existingPermissions.Permissions,
			}, nil
		}
	}
	log.Println("Fetching new permissions")
	*/

	keyExists, err := db.CheckIfApikeyExistsFS(key)
	if err != nil {
		return nil, err
	}

	if !keyExists {
		return &pb.ValidationResponse{
			Valid:       false,
			Permissions: []string{},
		}, nil
	}

	permissions, err := s.Auth0.User.Permissions(key.UserId)
	if err != nil {
		return nil, err
	}

	permissionsParsed := make([]string, len(permissions))

	for k, v := range permissions {
		permissionsParsed[k] = *v.Name
	}
	/*
		s.PermissionCache = PermissionDict[key.UserId] = PermissionCache{
			Valid:       true,
			Permissions: permissionsParsed,
			Expires:     time.Now(),
		}*/

	resp = &pb.ValidationResponse{
		Valid:       true,
		Permissions: permissionsParsed,
	}

	return resp, nil
}
