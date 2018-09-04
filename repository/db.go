package repository

import (
	"github.com/labstack/gommon/log"
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
		DbUser string `yaml:"db_user"`
		DbPassword string `yaml:"db_password"`
		DbHost string `yaml:"db_host"`
		DbPort string `yaml:"db_port"`
		DbName string `yaml:"db_name"`
	}
}{}

func GetManager() *manager {
	once.Do(func() {
		configor.Load(&parameters, "config/parameters.yml")
		dbParameter := parameters.Gorm.DbUser + ":" +
			parameters.Gorm.DbPassword + "@tcp(" +
			parameters.Gorm.DbHost + ":" +
			parameters.Gorm.DbPort + ")/" +
			parameters.Gorm.DbName + "?charset=utf8&parseTime=True"
		db, err := gorm.Open(parameters.Gorm.Driver, dbParameter)
		if err != nil {
			log.Print("Fail connect to db Error:"+ err.Error())
			instance = nil
		}else {
			instance = &manager{Db: db}
		}
	})
	return instance
}