
package main

import (
	"github.com/jinzhu/configor"
	"goMysqlChecking/route"
	"goMysqlChecking/gorutine"
)

var serverConfig =  struct {
	Server struct {
		HttpPort string `yaml:"http_port"`
		HttpsPort string `yaml:"https_port"`
	}
}{}

func main() {

	configor.Load(&serverConfig, "config/parameters.yml")
	go gorutine.DbCheckRoutine()
	router := route.Init()
	router.Logger.Fatal(router.Start(":"+serverConfig.Server.HttpPort))

}