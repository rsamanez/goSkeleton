package route

import (
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
	"template/api"
	"template/handler"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Debug()

	// Set Bundle MiddleWare
	e.Use(echoMw.Logger())
	e.Use(echoMw.Gzip())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
	e.SetHTTPErrorHandler(handler.JSONHTTPErrorHandler)

	// Routes
	v1 := e.Group("/api/v1")
	{
		//v1.POST("/members", api.PostMember())
		//v1.GET("/members", api.GetMembers())
		v1.GET("/members/:id", api.GetMember())
		//v1.GET("/members/:id", func(c echo.Context) error {
		//	return c.String(200, "test")
		//})
	}
	return e
}
