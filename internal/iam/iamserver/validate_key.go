package iamserver

import (
	"context"
	"oceanbolt.com/iamservice/internal/iam/dao"
	pb "oceanbolt.com/iamservice/rpc/iam"
	"time"
)

type PermissionCache struct {
	Valid       bool
	Permissions []string
	Expires     time.Time
}

//var PermissionDict map[string]PermissionCache

func (s *Server) ValidateKey(ctx context.Context, key *pb.UserKey) (resp *pb.ValidationResponse, err error) {
	db := dao.MgoDao{Ctx: ctx, Db: s.Db}

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

	keyExists, err := db.CheckKey(key)
	if err != nil {
		return nil, err
	}

	if !keyExists {
		return resp, nil
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
