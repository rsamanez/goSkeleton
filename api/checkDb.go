package api

import (
	"net/http"
	"github.com/labstack/echo"
	"goMysqlChecking/repository"
)


func GetCheckDb() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		result := repository.GetDataBaseDateFormat()
		if err, ok := result["sys_error"] ; ok {
			var msg = "System errors"
			if err == "db_error" {
				msg = "DB connection error"
			}
			return echo.NewHTTPError(http.StatusBadRequest, msg)
		}

		return c.JSON(http.StatusOK, result)
	}
}

