package repository

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/configor"
	"fmt"
)

type DataBaseDateFormat struct {
	DateFormat string
}

var debug =  struct {
	Debug struct {
		Logger bool `yaml:"logger"`
	}
}{}

/**
 * return map[string]string
 */
func GetDataBaseDateFormat() map[string]string {

	var data map[string]string
	data = make(map[string]string)
	configor.Load(&debug, "config/parameters.yml")
	if GetManager() != nil {
		if debug.Debug.Logger {
			fmt.Printf("DataBase Connected\n")
		}
		GetManager().Db.LogMode(debug.Debug.Logger)
		var reg DataBaseDateFormat
		GetManager().Db.Raw("SELECT DATE_FORMAT(\"2017-06-15\", \"%Y\") as date_format").First(&reg)
		if debug.Debug.Logger {
			fmt.Printf("%+v\n", reg)
		}
		data["status"] = "200"
		data["value"] = reg.DateFormat
	}else{
		data["sys_error"] = "400"
		data["value"] = ""
	}
	return data
}