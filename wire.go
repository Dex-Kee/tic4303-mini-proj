//go:build wireinject
// +build wireinject

package main

import (
	"tic4303-mini-proj/api"
	"tic4303-mini-proj/dao"
	"tic4303-mini-proj/middleware"
	"tic4303-mini-proj/page"
	"tic4303-mini-proj/router"
	"tic4303-mini-proj/service"

	"github.com/dzhcool/sven/setting"
	"github.com/google/wire"
)

func initServerRouter() *router.ServerRouter {
	wire.Build(dao.DaoSet, service.ServiceSet, api.ApiSet, page.PageSet, middleware.MiddlewareSet, router.RouterSet, initJwtSigningKey, initDigestKey)
	return nil
}

func initJwtSigningKey() []byte {
	return []byte(setting.Config.MustString("auth.signing.key", "tic4303-mini-proj"))
}

func initDigestKey() string {
	return setting.Config.MustString("auth.digest.key", "digestKey256")
}
