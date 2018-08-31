package repository

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"template/model"
)

type PartnerTable struct {
	Id int `gorm:"column:idpartner; primary_key:yes"`
	Code string `gorm:"column:code;"`
}

func (t PartnerTable) TableName() string {
	return "partners"
}

/**
 * return map[string]string
 */
func (t PartnerTable) GetItemById(paramId int64) (model.Partner, error) {

	var item model.Partner
	GetManager().Db.LogMode(true)
	var reg PartnerTable
	GetManager().Db.Select("idpartner, code").Where("idpartner = ?", paramId).First(&reg)
	//db.Raw("select idpartner, code from partners where idpartner = 3358 limit 1").Scan(&reg)

	fmt.Printf("%+v\n", reg)

	item.SetId(reg.Id)
	item.SetCode(reg.Code)

	fmt.Printf("%+v\n", item)

	return item, nil
}