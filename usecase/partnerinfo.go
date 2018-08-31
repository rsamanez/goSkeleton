package usecase

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
	"template/repository"
)

/**
 * return map[string]string
 */
func GetPartnerInfo(paramId int64) map[string]string {

	var data map[string]string
	data = make(map[string]string)
	var repo repository.PartnerTable
	partner, err := repo.GetItemById(paramId)
	if err != nil {
		data["partner_code"] = ""
		data["partner_id"] = "0"
	} else {
		data["partner_code"] = partner.GetCode()
		data["partner_id"] = strconv.Itoa( partner.GetId() )
	}

	return data
}

