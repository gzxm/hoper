package service

import (
	"context"
	"log"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	goauth "github.com/liov/hoper/go/v2/protobuf/utils/oauth"
	"github.com/liov/hoper/go/v2/protobuf/utils/response"
	"github.com/liov/hoper/go/v2/user/internal/config"
	"github.com/liov/hoper/go/v2/user/internal/dao"
	ijwt "github.com/liov/hoper/go/v2/utils/net/http/auth/jwt"
	"github.com/liov/hoper/go/v2/utils/net/http/auth/oauth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
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
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte(config.Conf.Customize.TokenSecret), jwt.SigningMethodHS512))

	clientStore := oauth.NewClientStore(dao.Dao.GORMDB)

	manager.MapClientStorage(clientStore)

	srv := oauth.NewServer(server.NewConfig(), manager)

	srv.UserAuthorizationHandler = func(token string) (userID, loginUri string) {
		loginUri = "loginUri"
		if token == "" {
			return
		}
		claims, err := ijwt.ParseToken(token, config.Conf.Customize.TokenSecret)
		if err != nil {
			return
		}
		return strconv.FormatUint(claims.UserID, 10), ""
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
}

func (u *OauthService) OauthAuthorize(ctx context.Context, req *goauth.OauthReq) (*response.HttpResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	tokens := md.Get("auth")
	tokens = append(tokens, "")
	req.AccessTokenExp = config.Conf.Customize.TokenMaxAge
	req.LoginUri = "/login"
	res := u.Server.HandleAuthorizeRequest(req, tokens[0])
	return res, nil
}

func (*OauthService) OauthToken(ctx context.Context, req *goauth.OauthReq) (*response.HttpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OauthToken not implemented")
}
