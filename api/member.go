package api

import (
	"net/http"
	"strconv"
	"rommel_samples/goSkeleton/usecase"
	"github.com/labstack/echo"
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

		return c.JSON(http.StatusOK, result)
	}
}

