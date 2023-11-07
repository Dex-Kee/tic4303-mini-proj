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

type MysqlConfig struct {
	Type, Host, Name, User, Passwd, Path, SSLMode, Port, Charset, TablePrefix string
}

func BuildMysqlDB(cfg *MysqlConfig) *gorm.DB {
	// new db engine
	db, err := newEngine(cfg)
	if err != nil {
		log.Fatalf("[orm] error: %v\n", err)
	}

	// sync table structure
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET="+cfg.Charset).AutoMigrate()

	gob.Register(time.Time{})

	return db
}

func getEngine(config *MysqlConfig) (*gorm.DB, error) {
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

func newEngine(config *MysqlConfig) (*gorm.DB, error) {
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
