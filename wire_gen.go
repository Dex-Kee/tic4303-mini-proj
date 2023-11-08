// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/dzhcool/sven/setting"
	"tic4303-mini-proj/api"
	"tic4303-mini-proj/dao"
	"tic4303-mini-proj/middleware"
	"tic4303-mini-proj/page"
	"tic4303-mini-proj/router"
	"tic4303-mini-proj/service"
	"tic4303-mini-proj/service/validation"
)

// Injectors from wire.go:

func initServerRouter() *router.ServerRouter {
	v := loadJwtSigningKey()
	string2 := loadDigestKey()
	mysqlConfig := loadDBCfg()
	db := middleware.BuildMysqlDB(mysqlConfig)
	userDAO := &dao.UserDAO{
		DB: db,
	}
	redisConfig := loadRedisCfg()
	client := middleware.BuildRedisClient(redisConfig)
	userValidationSvc := &validation.UserValidationSvc{
		UserDAO:     userDAO,
		RedisClient: client,
	}
	userSvc := &service.UserSvc{
		JwtSigningKey:     v,
		DigestKey:         string2,
		UserDAO:           userDAO,
		UserValidationSvc: userValidationSvc,
		RedisClient:       client,
	}
	userApi := &api.UserApi{
		UserSvc: userSvc,
	}
	infoApi := &api.InfoApi{}
	userPage := &page.UserPage{}
	authFilter := &middleware.AuthFilter{
		JwtSigningKey: v,
		DigestKey:     string2,
		UserSvc:       userSvc,
		RedisClient:   client,
	}
	serverRouter := &router.ServerRouter{
		UserApi:    userApi,
		InfoApi:    infoApi,
		UserPage:   userPage,
		AuthFilter: authFilter,
	}
	return serverRouter
}

// wire.go:

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
