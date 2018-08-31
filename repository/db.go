package repository

import (
	"echo-sample/_vendor-20180829165813/github.com/labstack/gommon/log"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	"sync"
)

type manager struct {
	Db *gorm.DB
}

var instance *manager
var once sync.Once

var parameters =  struct {
	Gorm struct {
		Driver string `yaml:"driver"`
		ConnStr string `yaml:"conn_str"`
	}
}{}

func GetManager() *manager {
	once.Do(func() {
		configor.Load(&parameters, "config/parameters.yml")

		db, err := gorm.Open(parameters.Gorm.Driver, parameters.Gorm.ConnStr)
		if err != nil {
			log.Fatal("Fail connect to db")
		}
		instance = &manager {Db: db}
	})
	return instance
}