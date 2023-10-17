//go:build wireinject
// +build wireinject

package main

import (
	"tic4303-mini-proj/api"
	"tic4303-mini-proj/dao"
	"tic4303-mini-proj/middleware"
	"tic4303-mini-proj/router"
	"tic4303-mini-proj/service"

	"github.com/google/wire"
)

func initServerRouter() *router.ServerRouter {
	wire.Build(dao.DaoSet, service.ServiceSet, api.ApiSet, middleware.MiddlewareSet, router.RouterSet)
	return nil
}
