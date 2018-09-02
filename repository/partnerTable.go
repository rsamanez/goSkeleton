package repository

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"rommel_samples/goSkeleton/model"
	"github.com/jinzhu/configor"
	"fmt"
)

type PartnerTable struct {
	Id int `gorm:"column:idpartner; primary_key:yes"`
	Code string `gorm:"column:code;"`
}

var debug =  struct {
	Debug struct {
		Logger bool `yaml:"logger"`
	}
}{}

func (t PartnerTable) TableName() string {
	return "partners"
}

/**
 * return map[string]string
 */
func (t PartnerTable) GetItemById(paramId int64) (model.Partner, error) {

	var item model.Partner
	configor.Load(&debug, "config/parameters.yml")
	GetManager().Db.LogMode(debug.Debug.Logger)
	var reg PartnerTable
	GetManager().Db.Select("idpartner, code").Where("idpartner = ?", paramId).First(&reg)
	//db.Raw("select idpartner, code from partners where idpartner = 3358 limit 1").Scan(&reg)

	if debug.Debug.Logger {
		fmt.Printf("%+v\n", reg)
	}
	item.SetId(reg.Id)
	item.SetCode(reg.Code)
	if debug.Debug.Logger {
		fmt.Printf("%+v\n", item)
	}
	return item, nil
}