package route

import (
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
	"rommel_samples/goSkeleton/api"
	"github.com/jinzhu/configor"
)

var debug =  struct {
	Debug struct {
		Logger bool `yaml:"logger"`
	}
}{}
func Init() *echo.Echo {
	configor.Load(&debug, "config/parameters.yml")
	e := echo.New()
	if debug.Debug.Logger {
		//e.Debug()
		e.Use(echoMw.Logger())
	}
	// Set Bundle MiddleWare
	e.Use(echoMw.Gzip())

	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	//e.SetHTTPErrorHandler(handler.JSONHTTPErrorHandler)

	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.GET("/partners/:id", api.GetPartner())
	}
	return e
}
