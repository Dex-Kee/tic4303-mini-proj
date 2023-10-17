package main

import (
	"flag"
	"tic4303-mini-proj/router"

	"github.com/dzhcool/sven/setting"
	log "github.com/dzhcool/sven/zapkit"
	"github.com/gin-gonic/gin"
)

func main() {
	initConfigFile()
	initLog()

	serverRouter := initServerRouter()
	app := initGinEngine(serverRouter)

	_ = app.Run(":" + setting.Config.MustString("http.port", "8080"))
}

func initConfigFile() {
	var file string
	var env string
	flag.StringVar(&file, "c", "./config/app", "please specify config file")
	flag.StringVar(&env, "e", "", "please specify running environment")
	flag.Parse()
	setting.InitSetting(file, env)
}

func initLog() {
	cfg := &log.ZapkitConfig{
		File:       setting.Config.MustString("zapkit.file", "./app.log"),
		Level:      setting.Config.MustString("zapkit.level", "info"),
		MaxSize:    setting.Config.MustInt("zapkit.maxsize", 512),
		MaxBackups: setting.Config.MustInt("zapkit.maxbackups", 10),
		MaxAge:     setting.Config.MustInt("zapkit.age", 7),
		Compress:   setting.Config.MustBool("zapkit.compress", false),
	}
	err := log.Init(cfg)
	if err != nil {
		panic(err)
	}
}

func initGinEngine(r *router.ServerRouter) *gin.Engine {
	app := gin.Default()

	app.NoRoute(r.NoRouterHandler)
	app.NoMethod(r.NoMethodHandler)

	r.RegisterApi(app)
	r.RegisterPage(app)
	return app
}
