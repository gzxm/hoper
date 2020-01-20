package iris_gateway

import (
	"context"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/handlerconv"
	"github.com/liov/hoper/go/v2/utils/http/debug"
	"github.com/liov/hoper/go/v2/utils/http/gateway"
	"github.com/liov/hoper/go/v2/utils/log"
)

func Http(irisHandle func(*iris.Application), gatewayHandle func(context.Context, *runtime.ServeMux)) http.Handler {
	gwmux := gateway.Gateway(gatewayHandle)
	//openapi
	mux := iris.New()
	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		mux.Shutdown(ctx)
	})
	mux.Any("/{grpc:path}", handlerconv.FromStd(gwmux))
	mux.Any("/debug/{path:path}", handlerconv.FromStd(debug.Debug()))
	if irisHandle != nil {
		irisHandle(mux)
	}
	if err := mux.Build(); err != nil {
		log.Fatal(err)
	}
	return mux
}