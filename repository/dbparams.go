package repository

import (
	"sync"
)

// DEPRECATED

type dbparams struct {
	Driver string
	ConnStr string
}

var dbp *dbparams
var once2 sync.Once

func SetParams(driver string, connstr string) {
	once.Do(func() {
		dbp = &dbparams {Driver: driver, ConnStr: connstr}
	})
}

func GetParams() *dbparams {
	return dbp
}
