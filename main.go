
package main

import (
	//"OLD/echo-sample/_vendor-20180829132321/github.com/labstack/echo/engine/fasthttp"
	//"OLD/echo-sample/_vendor-20180829132321/github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/engine/fasthttp"
	"template/route"
	//"github.com/labstack/echo"
	//"net/http"
)

func main() {

	router := route.Init()
	router.Run(fasthttp.New(":8888"))
	//router.Start(":1323")

	//e := echo.New()
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	//e.Logger.Fatal(e.Start(":1323"))
}
