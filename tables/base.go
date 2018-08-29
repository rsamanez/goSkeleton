package tables

import (
	"github.com/jinzhu/gorm"
	"template/models"
)

type PartnerTable struct {
	name string
}

func (t PartnerTable) TableName() string {
	return "partners"
}

func connect(p models.Partner) gorm.DB {
	db, err := gorm.Open("mysql", "aldo.diaz:wSHFk5pMrC6k@tcp(172.19.20.212:3306)/bongo")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	return db
}

func getItem(p models.Partner) models.Partner {
	db := connect(p)
	var par models.Partner
	db.First(&par)
	return par
}