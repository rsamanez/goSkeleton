package models

type Partner struct {
	Idpartner int `gorm:"column:IndepYear;"`
	Code string `gorm:"column:Code; primary_key:yes"`
}

func getId(p Partner) int {
	return p.Idpartner
}

func getCode(p Partner) string {
	return p.Code
}
