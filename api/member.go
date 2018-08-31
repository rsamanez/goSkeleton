package api

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strconv"
	"template/usecase"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)


func GetPartner() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		paramId, _ := strconv.ParseInt(c.Param("id"), 0, 64)

		result := usecase.GetPartnerInfo(paramId)
		if err, ok := result["sys_error"] ; ok {
			var msg string = "System errors"
			if err == "db_error" {
				msg = "DB connection error"
			}
			return echo.NewHTTPError(http.StatusBadRequest, msg)
		}

		return c.JSON(fasthttp.StatusOK, result)
	}
}

