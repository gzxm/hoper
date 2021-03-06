package service

import (
	"context"
	"github.com/actliboy/hoper/server/go/lib/protobuf/empty"
	"github.com/actliboy/hoper/server/go/lib/protobuf/errorcode"
	contexti "github.com/actliboy/hoper/server/go/lib/tiga/context"
	dbi "github.com/actliboy/hoper/server/go/lib/utils/dao/db"
	"github.com/actliboy/hoper/server/go/mod/protobuf/user"
	"github.com/actliboy/hoper/server/go/mod/user/dao"
	"github.com/actliboy/hoper/server/go/mod/user/model"
)

// 关注
func (u *UserService) Follow(ctx context.Context, req *user.FollowReq) (*empty.Empty, error) {
	ctxi, span := contexti.CtxFromContext(ctx).StartSpan("")
	defer span.End()
	ctx = ctxi.Context
	auth, err := auth(ctxi, true)
	if err != nil {
		return nil, err
	}
	db := ctxi.NewDB(dao.Dao.GORMDB)
	userDao := dao.GetDao(ctxi)
	exists, err := userDao.FollowExistsDB(db, req.Id, auth.Id)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, nil
	}
	err = db.Table(model.FollowTableName).Create(&user.UserFollow{
		UserId:   req.Id,
		FollowId: auth.Id,
	}).Error
	if err != nil {
		return nil, ctxi.ErrorLog(errorcode.DBError, err, "Create")
	}
	return new(empty.Empty), nil
}

// 取消关注
func (u *UserService) DelFollow(ctx context.Context, req *user.FollowReq) (*user.BaseListRep, error) {
	ctxi, span := contexti.CtxFromContext(ctx).StartSpan("")
	defer span.End()
	ctx = ctxi.Context
	auth, err := auth(ctxi, true)
	if err != nil {
		return nil, err
	}
	db := ctxi.NewDB(dao.Dao.GORMDB)
	userDao := dao.GetDao(ctxi)
	exists, err := userDao.FollowExistsDB(db, req.Id, auth.Id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, nil
	}
	err = db.Table(model.FollowTableName).Where("user_id = ? AND follow_id = ? AND "+dbi.PostgreNotDeleted, req.Id, auth.Id).
		UpdateColumn("deleted_at", ctxi.RequestAt.TimeString).Error
	if err != nil {
		return nil, ctxi.ErrorLog(errorcode.DBError, err, "Create")
	}
	return nil, nil
}
