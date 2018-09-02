
package main

import (
	"github.com/jinzhu/configor"
	"rommel_samples/goSkeleton/route"
)

var serverConfig =  struct {
	Server struct {
		HttpPort string `yaml:"http_port"`
		HttpsPort string `yaml:"https_port"`
	}
}{}

func main() {

	configor.Load(&serverConfig, "config/parameters.yml")
	router := route.Init()
	//router.Logger.Fatal(router.Start(":"+serverConfig.Server.HttpPort))
	router.Logger.Fatal(router.StartTLS(":"+serverConfig.Server.HttpsPort, "cert.pem", "key.pem"))
}