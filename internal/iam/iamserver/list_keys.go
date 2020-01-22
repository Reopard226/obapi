package iamserver

import (
	"context"
	"oceanbolt.com/obapi/internal/iam/dao"
	"oceanbolt.com/obapi/rpc/iam"
)

func (s *Server) ListKeys(ctx context.Context, user *iam.User) (keys *iam.UserKeys, err error) {
	db := dao.IamDAO{Ctx: ctx, Db: s.Db, Fs: s.Fs}

	return db.ListKeysFS(user)
}
