package middleware

import (
	"encoding/gob"
	"fmt"
	"log"
	"time"

	"github.com/dzhcool/sven/setting"
	"github.com/dzhcool/sven/zapkit"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type mysqlConfig struct {
	Type, Host, Name, User, Passwd, Path, SSLMode, Port, Charset, TablePrefix string
}

func InitMysql() *gorm.DB {
	config := loadConfigs("api")

	// new db engine
	db, err := newEngine(config)
	if err != nil {
		log.Fatalf("[orm] error: %v\n", err)
	}

	// sync table structure
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET="+config.Charset).AutoMigrate()

	gob.Register(time.Time{})

	return db
}

func loadConfigs(name string) *mysqlConfig {
	config := new(mysqlConfig)
	config.Host = setting.Config.MustString("db."+name+".host", "")
	config.Name = setting.Config.MustString("db."+name+".name", "")
	config.User = setting.Config.MustString("db."+name+".user", "")
	config.Passwd = setting.Config.MustString("db."+name+".passwd", "")
	config.Port = setting.Config.MustString("db."+name+".port", "")
	config.Charset = setting.Config.MustString("db."+name+".charset", "utf8")
	config.TablePrefix = setting.Config.MustString("db."+name+".prefix", "")
	return config
}

func getEngine(config *mysqlConfig) (*gorm.DB, error) {
	dsn := ""
	if config.Host[0] == '/' { // looks like a unix socket
		dsn = fmt.Sprintf("%s:%s@unix(%s:%s)/%s?charset=%s&timeout=3s&parseTime=true&loc=Local",
			config.User, config.Passwd, config.Host, config.Port, config.Name, config.Charset)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&timeout=3s&parseTime=true&loc=Local",
			config.User, config.Passwd, config.Host, config.Port, config.Name, config.Charset)
	}
	zapkit.Debugf("%s", dsn)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func newEngine(config *mysqlConfig) (*gorm.DB, error) {
	db, err := getEngine(config)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// connection setting
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(500)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// log setting
	logger.Default.LogMode(logger.Silent)
	debug, _ := setting.Config.GetBool("app.debug")
	if !debug {
		logger.Default.LogMode(logger.Info)
	}

	return db, nil
}
