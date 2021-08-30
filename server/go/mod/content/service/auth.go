package service

import (
	contexti "github.com/liov/hoper/server/go/lib/tiga/context"
	"github.com/liov/hoper/server/go/mod/protobuf/user"
	"github.com/liov/hoper/server/go/mod/user/service"
)

func auth(ctx *contexti.Ctx, update bool) (*user.AuthInfo, error) {
	return service.ExportAuth(ctx, update)
}