package iamserver

import (
	"context"
	"oceanbolt.com/iamservice/internal/iam/dao"
	"oceanbolt.com/iamservice/rpc/iam"
)

func (s *Server) DeleteKey(ctx context.Context, req *iam.DeleteKeyRequest) (resp *iam.KeyDeletedResponse, err error) {
	db := dao.MgoDao{Ctx: ctx, Db: s.Db}

	err = db.DeleteKey(req)
	if err != nil {
		return resp, err
	}

	resp = &iam.KeyDeletedResponse{
		Message: "Key '" + req.ApikeyId + "' successfully deleted",
	}
	return resp, nil
}
