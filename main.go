
package main

import (
	"github.com/labstack/echo/engine/fasthttp"
	"template/route"
)

func main() {
	router := route.Init()
	router.Run(fasthttp.New(":8888"))
}
