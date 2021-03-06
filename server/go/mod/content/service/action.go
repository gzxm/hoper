package service

import (
	"context"
	"github.com/actliboy/hoper/server/go/lib/protobuf/empty"
	"github.com/actliboy/hoper/server/go/lib/protobuf/errorcode"
	"github.com/actliboy/hoper/server/go/lib/protobuf/request"
	contexti "github.com/actliboy/hoper/server/go/lib/tiga/context"
	"github.com/actliboy/hoper/server/go/lib/utils/log"
	"github.com/actliboy/hoper/server/go/lib/utils/slices"
	"github.com/actliboy/hoper/server/go/mod/content/client"
	"github.com/actliboy/hoper/server/go/mod/content/conf"
	"github.com/actliboy/hoper/server/go/mod/content/dao"
	dbdao "github.com/actliboy/hoper/server/go/mod/content/dao/db"
	"github.com/actliboy/hoper/server/go/mod/content/model"
	"github.com/actliboy/hoper/server/go/mod/protobuf/content"
	"github.com/actliboy/hoper/server/go/mod/protobuf/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ActionService struct {
	content.UnimplementedActionServiceServer
}

func (*ActionService) Like(ctx context.Context, req *content.LikeReq) (*request.Object, error) {
	if req.Action != content.ActionLike && req.Action != content.ActionUnlike && req.Action != content.ActionBrowse {
		return nil, nil
	}
	ctxi, span := contexti.CtxFromContext(ctx).StartSpan("")
	defer span.End()
	auth, err := auth(ctxi, true)
	if err != nil {
		return nil, err
	}
	db := ctxi.NewDB(dao.Dao.GORMDB)

	contentDBDao := dbdao.GetDao(ctxi, db)

	req.UserId = auth.Id

	id, err := contentDBDao.LikeId(req.Type, req.Action, req.RefId, req.UserId)
	if err != nil {
		return &request.Object{Id: id}, err
	}
	if id > 0 {
		return &request.Object{Id: id}, nil
	}

	err = contentDBDao.Transaction(func(tx *gorm.DB) error {
		contenttxDBDao := dbdao.GetDao(ctxi, tx)
		err = tx.Table(model.LikeTableName).Create(req).Error
		if err != nil {
			return ctxi.ErrorLog(errorcode.DBError, err, "Create")
		}

		err = contenttxDBDao.ActionCount(req.Type, req.Action, req.RefId, 1)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	contentRedisDao := dao.GetRedisDao(ctxi, dao.Dao.Redis)
	err = contentRedisDao.HotCount(req.Type, req.RefId, 1)

	if err != nil {
		return nil, ctxi.ErrorLog(errorcode.RedisErr, err, "HotCountRedis")
	}
	return &request.Object{Id: req.Id}, nil
}

func (*ActionService) DelLike(ctx context.Context, req *request.Object) (*empty.Empty, error) {
	ctxi, span := contexti.CtxFromContext(ctx).StartSpan("")
	defer span.End()
	auth, err := auth(ctxi, true)
	if err != nil {
		return nil, err
	}
	db := ctxi.NewDB(dao.Dao.GORMDB)
	contentDBDao := dao.GetDBDao(ctxi, db)

	like, err := contentDBDao.GetLike(req.Id, auth.Id)
	if err != nil {
		return nil, err
	}
	if like.Id == 0 {
		return nil, errorcode.ParamInvalid
	}
	err = contentDBDao.DelByAuth(model.LikeTableName, req.Id, auth.Id)
	if err != nil {
		return nil, err
	}
	err = contentDBDao.ActionCount(like.Type, like.Action, like.RefId, -1)
	if err != nil {
		return nil, err
	}
	contentRedisDao := dao.GetRedisDao(ctxi, dao.Dao.Redis)
	err = contentRedisDao.HotCount(like.Type, like.RefId, -1)
	if err != nil {
		return nil, err
	}
	return new(empty.Empty), err
}

func (*ActionService) Comment(ctx context.Context, req *content.CommentReq) (*request.Object, error) {
	ctxi, span := contexti.CtxFromContext(ctx).StartSpan("")
	defer span.End()
	auth, err := auth(ctxi, true)
	if err != nil {
		return nil, err
	}
	db := ctxi.NewDB(dao.Dao.GORMDB)

	req.UserId = auth.Id
	err = db.Transaction(func(tx *gorm.DB) error {
		contenttxDBDao := dbdao.GetDao(ctxi, tx)
		err = tx.Table(model.CommentTableName).Create(req).Error
		if err != nil {
			return ctxi.ErrorLog(errorcode.DBError, err, "Create")
		}
		err = contenttxDBDao.CreateContextExt(content.ContentComment, req.Id)
		if err != nil {
			return err
		}
		err = contenttxDBDao.ActionCount(req.Type, content.ActionComment, req.RefId, 1)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		if err != errorcode.DBError {
			ctxi.Error(err.Error(), zap.String(log.Position, "Transaction"))
		}
		return nil, err
	}
	contentRedisDao := dao.GetRedisDao(ctxi, dao.Dao.Redis)
	err = contentRedisDao.HotCount(req.Type, req.RefId, 1)
	if err != nil {
		return nil, err
	}
	return &request.Object{Id: req.Id}, nil
}

func (*ActionService) DelComment(ctx context.Context, req *request.Object) (*empty.Empty, error) {
	ctxi, span := contexti.CtxFromContext(ctx).StartSpan("")
	defer span.End()
	auth, err := auth(ctxi, true)
	if err != nil {
		return nil, err
	}
	db := ctxi.NewDB(dao.Dao.GORMDB)
	contentDBDao := dao.GetDBDao(ctxi, db)

	var comment content.Comment
	err = db.Table(model.CommentTableName).First(&comment, "id = ?", req.Id).Error
	if err != nil {
		return nil, ctxi.ErrorLog(errorcode.DBError, err, "Find")
	}
	if comment.UserId != auth.Id {
		var userId uint64
		err = db.Table(model.ContentTableName(comment.Type)).Select("user_id").
			Where(`id = ?`, comment.RefId).Row().Scan(&userId)
		if err != nil {
			return nil, ctxi.ErrorLog(errorcode.DBError, err, "SelectUserId")
		}
		if userId != auth.Id {
			return nil, errorcode.PermissionDenied
		}
	}

	err = contentDBDao.Del(model.CommentTableName, req.Id)
	if err != nil {
		return nil, err
	}
	err = contentDBDao.ActionCount(comment.Type, content.ActionComment, comment.RefId, -1)
	if err != nil {
		return nil, err
	}
	contentRedisDao := dao.GetRedisDao(ctxi, dao.Dao.Redis)
	err = contentRedisDao.HotCount(comment.Type, comment.RefId, -1)
	if err != nil {
		return nil, err
	}
	return new(empty.Empty), nil
}

func (*ActionService) Collect(ctx context.Context, req *content.CollectReq) (*empty.Empty, error) {
	ctxi, span := contexti.CtxFromContext(ctx).StartSpan("")
	defer span.End()
	auth, err := auth(ctxi, true)
	if err != nil {
		return nil, err
	}
	db := ctxi.NewDB(dao.Dao.GORMDB)
	contentDBDao := dao.GetDBDao(ctxi, db)

	req.UserId = auth.Id
	collects, err := contentDBDao.GetCollects(req.Type, []uint64{req.RefId}, auth.Id)
	if err != nil {
		return nil, err
	}
	var origin []uint64
	for _, collect := range collects {
		origin = append(origin, collect.FavId)
	}
	diff := slices.DiffUint64(origin, req.FavIds)
	collect := model.Collect{
		Type:   req.Type,
		RefId:  req.RefId,
		UserId: auth.Id,
		FavId:  0,
	}
	for _, id := range diff {
		collect.FavId = id
		err = db.Table(model.CollectTableName).Create(&collect).Error
		if err != nil {
			return nil, ctxi.ErrorLog(errorcode.DBError, err, "Create")
		}
	}
	if len(origin) == 0 && len(req.FavIds) > 0 {
		err = contentDBDao.ActionCount(req.Type, content.ActionCollect, req.RefId, 1)
		if err != nil {
			return nil, ctxi.ErrorLog(errorcode.DBError, err, "ActionCount")
		}
	}
	err = db.Table(model.CollectTableName).Where(`type = ? AND ref_id = ? AND fav_id NOT IN (?)`, req.Type, req.RefId, req.FavIds).
		Update(`deleted_at`, ctxi.RequestAt.TimeString).Error
	if err != nil {
		return nil, ctxi.ErrorLog(errorcode.DBError, err, "DELETE")
	}
	var hotCount float64
	if len(origin) == 0 && len(req.FavIds) > 0 {
		hotCount = 1
	}
	if len(origin) > 0 && len(req.FavIds) == 0 {
		hotCount = -1
	}
	if hotCount != 0 {
		contentRedisDao := dao.GetRedisDao(ctxi, dao.Dao.Redis)
		err = contentRedisDao.HotCount(req.Type, req.RefId, hotCount)
		if err != nil {
			return nil, err
		}
	}

	return new(empty.Empty), nil
}

func (*ActionService) Report(ctx context.Context, req *content.ReportReq) (*empty.Empty, error) {
	ctxi, span := contexti.CtxFromContext(ctx).StartSpan("")
	defer span.End()
	auth, err := auth(ctxi, true)
	if err != nil {
		return nil, err
	}
	contentRedisDao := dao.GetRedisDao(ctxi, dao.Dao.Redis)
	err = contentRedisDao.Limit(&conf.Conf.Customize.Moment.Limit)
	if err != nil {
		return nil, err
	}
	db := ctxi.NewDB(dao.Dao.GORMDB)
	req.UserId = auth.Id
	err = db.Transaction(func(tx *gorm.DB) error {
		contenttxDBDao := dao.GetDBDao(ctxi, tx)
		err = tx.Table(model.ReportTableName).Create(req).Error
		if err != nil {
			return ctxi.ErrorLog(errorcode.DBError, err, "Create")
		}
		err = contenttxDBDao.ActionCount(req.Type, content.ActionReport, req.RefId, 1)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		if err != errorcode.DBError {
			ctxi.Error(err.Error(), zap.String(log.Position, "Transaction"))
		}
		return nil, errorcode.DBError
	}
	return new(empty.Empty), nil
}

func (*ActionService) CommentList(ctx context.Context, req *content.CommentListReq) (*content.CommentListRep, error) {
	ctxi, span := contexti.CtxFromContext(ctx).StartSpan("")
	defer span.End()
	auth, err := auth(ctxi, true)
	if err != nil {
		return nil, err
	}
	db := ctxi.NewDB(dao.Dao.GORMDB)
	contentDBDao := dao.GetDBDao(ctxi, db)

	total, comments, err := contentDBDao.GetComments(content.ContentMoment, req.RefId, req.RootId, int(req.PageNo), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	var ids, userIds []uint64
	var m = make(map[uint64]*content.Comment)
	for i := range comments {
		ids = append(ids, comments[i].Id)
		m[comments[i].Id] = comments[i]
		userIds = append(userIds, comments[i].UserId)
		userIds = append(userIds, comments[i].RecvId)
		// ????????????
		commentMaskField(comments[i])
	}
	// ext
	exts, err := contentDBDao.GetContentExt(content.ContentComment, ids)
	if err != nil {
		return nil, err
	}
	for i := range exts {
		if comment, ok := m[exts[i].RefId]; ok {
			comment.Ext = exts[i]
		}
	}

	//like
	if auth.Id != 0 {
		likes, err := contentDBDao.GetContentActions(content.ActionLike, content.ContentComment, ids, auth.Id)
		if err != nil {
			return nil, err
		}
		for i := range likes {
			if comment, ok := m[likes[i].RefId]; ok {
				if comment.Action == nil {
					comment.Action = &content.UserAction{}
				}
				if likes[i].Action == content.ActionLike {
					comment.Action.LikeId = likes[i].Id
				}
				if likes[i].Action == content.ActionUnlike {
					comment.Action.UnlikeId = likes[i].Id
				}
			}
		}
		collects, err := contentDBDao.GetCollects(content.ContentComment, ids, auth.Id)
		if err != nil {
			return nil, err
		}
		for i := range collects {
			if comment, ok := m[collects[i].RefId]; ok {
				if comment.Action == nil {
					comment.Action = &content.UserAction{}
				}
				comment.Action.Collects = append(comment.Action.Collects, collects[i].FavId)
			}
		}
	}
	var users []*user.UserBaseInfo
	if len(userIds) > 0 {
		userList, err := client.UserClient.BaseList(ctxi, &user.BaseListReq{Ids: userIds})
		if err != nil {
			return nil, err
		}
		users = userList.List
	}
	return &content.CommentListRep{
		Total: total,
		List:  comments,
		Users: users,
	}, nil
}

// ????????????
func commentMaskField(comment *content.Comment) {
	comment.DeletedAt = ""
	comment.CreatedAt = comment.CreatedAt[:19]
}

func (*ActionService) GetUserAction(ctx context.Context, req *content.ContentReq) (*content.UserAction, error) {
	ctxi, span := contexti.CtxFromContext(ctx).StartSpan("")
	defer span.End()
	auth, err := auth(ctxi, true)
	if err != nil {
		return nil, err
	}

	db := ctxi.NewDB(dao.Dao.GORMDB)
	contentDBDao := dbdao.GetDao(ctxi, db)

	action := &content.UserAction{}
	likes, err := contentDBDao.GetContentActions(content.ActionLike, content.ContentMoment, []uint64{req.RefId}, auth.Id)
	if err != nil {
		return nil, err
	}

	for i := range likes {
		if likes[i].Action == content.ActionLike {
			action.LikeId = likes[i].Id
		}
		if likes[i].Action == content.ActionUnlike {
			action.UnlikeId = likes[i].Id
		}
	}
	collects, err := contentDBDao.GetCollects(content.ContentMoment, []uint64{req.RefId}, auth.Id)
	if err != nil {
		return nil, err
	}
	for i := range collects {
		action.Collects = append(action.Collects, collects[i].FavId)
	}
	return action, nil
}
