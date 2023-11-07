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
	"tic4303-mini-proj/service/validation"

	"github.com/dzhcool/sven/setting"
	"github.com/google/wire"
)

func initServerRouter() *router.ServerRouter {
	wire.Build(dao.DaoSet, service.ServiceSet, validation.ValidationSet, api.ApiSet, page.PageSet, middleware.MiddlewareSet, router.RouterSet, loadJwtSigningKey, loadDigestKey,
		loadDBCfg, loadRedisCfg)
	return nil
}

func loadJwtSigningKey() []byte {
	return []byte(setting.Config.MustString("auth.signing.key", "tic4303-mini-proj"))
}

func loadDigestKey() string {
	return setting.Config.MustString("auth.digest.key", "digestKey256")
}

func loadDBCfg() *middleware.MysqlConfig {
	config := new(middleware.MysqlConfig)
	config.Host = setting.Config.MustString("db.host", "")
	config.Name = setting.Config.MustString("db.name", "")
	config.User = setting.Config.MustString("db.user", "")
	config.Passwd = setting.Config.MustString("db.passwd", "")
	config.Port = setting.Config.MustString("db.port", "")
	config.Charset = setting.Config.MustString("db.charset", "utf8")
	config.TablePrefix = setting.Config.MustString("db.prefix", "")
	return config
}

func loadRedisCfg() *middleware.RedisConfig {
	config := new(middleware.RedisConfig)
	config.Host = setting.Config.MustString("redis.host", "127.0.0.1")
	config.Port = setting.Config.MustString("redis.port", "5379")
	config.DB = setting.Config.MustInt("redis.db", 0)
	config.Password = setting.Config.MustString("redis.password", "")
	return config
}
