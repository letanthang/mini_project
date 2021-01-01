package repo

import (
	"fmt"
	"sync"
	"time"

	"app/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/gommon/log"
)

var pgDB *gorm.DB

func GetDB() *gorm.DB {

	var once sync.Once

	once.Do(func() {
		pgDB = new()
	})
	return pgDB
}

func new() *gorm.DB {
	gorm.NowFunc = func() time.Time {
		return time.Now().UTC().Truncate(1000 * time.Nanosecond)
	}
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		config.Config.Db.Host, config.Config.Db.User, config.Config.Db.Name, config.Config.Db.Password))
	if err != nil {
		log.Infof("Connect not success to postgres database at host:%s with user:%s and db:%s and %s",
			config.Config.Db.Host, config.Config.Db.User, config.Config.Db.Name, config.Config.Db.Password)
		panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().Ping()
	if config.Config.Db.Debug {
		db.LogMode(true)
	}
	return db
}
