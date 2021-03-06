package service

import (
	"context"
	"log"
	"strconv"
	"time"

	goauth "github.com/actliboy/hoper/server/go/lib/protobuf/oauth"
	"github.com/actliboy/hoper/server/go/lib/protobuf/response"
	stringsi "github.com/actliboy/hoper/server/go/lib/utils/strings"
	jwti "github.com/actliboy/hoper/server/go/lib/utils/verification/auth/jwt"
	"github.com/actliboy/hoper/server/go/lib/utils/verification/auth/oauth"
	"github.com/actliboy/hoper/server/go/mod/protobuf/user"
	"github.com/actliboy/hoper/server/go/mod/user/conf"
	"github.com/actliboy/hoper/server/go/mod/user/dao"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"google.golang.org/grpc/metadata"
)

func GetOauthService() *OauthService {
	if oauthSvc != nil {
		return oauthSvc
	}
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", stringsi.ToBytes(conf.Conf.Customize.TokenSecret), jwt.SigningMethodHS512))

	clientStore := oauth.NewClientStore(dao.Dao.GORMDB)

	manager.MapClientStorage(clientStore)

	srv := oauth.NewServer(server.NewConfig(), manager)

	srv.UserAuthorizationHandler = func(token string) (userID string, err error) {
		if token == "" {
			return "", errors.ErrInvalidAccessToken
		}
		claims := new(jwti.Claims)
		if err := jwti.ParseToken(claims, token, conf.Conf.Customize.TokenSecret); err != nil {
			return "", err
		}
		return strconv.FormatUint(claims.UserId, 10), nil
	}

	srv.InternalErrorHandler = func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	}

	srv.ResponseErrorHandler = func(re *errors.Response) {
		log.Println("HttpResponse Error:", re.Error.Error())
	}
	oauthSvc = &OauthService{Server: srv, ClientStore: clientStore}
	return oauthSvc
}

type OauthService struct {
	Server      *oauth.Server
	ClientStore *oauth.ClientStore
	user.UnimplementedOauthServiceServer
}

func (u *OauthService) OauthAuthorize(ctx context.Context, req *goauth.OauthReq) (*response.HttpResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	tokens := md.Get("auth")
	tokens = append(tokens, "")
	req.AccessTokenExp = int64(24 * time.Hour)
	req.LoginURI = "/oauth/login"
	res := u.Server.HandleAuthorizeRequest(ctx, req, tokens[0])
	return res, nil
}

func (u *OauthService) OauthToken(ctx context.Context, req *goauth.OauthReq) (*response.HttpResponse, error) {
	req.GrantType = string(oauth2.AuthorizationCode)
	res, _ := u.Server.HandleTokenRequest(ctx, req)
	return res, nil
}
