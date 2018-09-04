package gorutine

import (
	"github.com/labstack/gommon/log"
	"github.com/jinzhu/gorm"
		"fmt"
	"time"
	"github.com/jinzhu/configor"
)

type DataBaseDateFormat struct {
	DateFormat string
}

var dbConnection *gorm.DB

var parameters =  struct {
	Gorm struct {
		Driver string `yaml:"driver"`
		DbUser string `yaml:"db_user"`
		DbPassword string `yaml:"db_password"`
		DbHost string `yaml:"db_host"`
		DbPort string `yaml:"db_port"`
		DbName string `yaml:"db_name"`
	}
	Debug struct {
		Logger bool `yaml:"logger"`
	}
	Checkdb struct {
		LoopTimer time.Duration `yaml:"loop_timer"`
	}
}{}

func OpenConnection(){
	dbParameter := parameters.Gorm.DbUser + ":" +
		parameters.Gorm.DbPassword + "@tcp(" +
		parameters.Gorm.DbHost + ":" +
		parameters.Gorm.DbPort + ")/" +
		parameters.Gorm.DbName + "?charset=utf8&parseTime=True"
	db, err := gorm.Open(parameters.Gorm.Driver, dbParameter)
	if err != nil {
		log.Print("Fail connect to db Error:"+ err.Error())
	}else {
		fmt.Printf("DataBase Connected\n")
		dbConnection = db
		dbConnection.LogMode(parameters.Debug.Logger)
	}
}

func Start(){
	configor.Load(&parameters, "config/parameters.yml")
	OpenConnection()
}

func DbCheckRoutine(){
	dbConnection = nil
	Start()
	for {
		fmt.Println("Infinite Loop 1")
		time.Sleep(time.Second * parameters.Checkdb.LoopTimer)
		if dbConnection != nil{
			var reg DataBaseDateFormat
			if result := dbConnection.Raw("SELECT DATE_FORMAT(\"2017-06-15\", \"%Y\") as date_format").First(&reg); result.Error != nil {
				log.Print("Fail connect to db Error: ")
			}else{
				if reg.DateFormat != "2017"{
					log.Print("Fail Db response: ")
				}
			}
		}else{
			Start()
		}
	}
}


