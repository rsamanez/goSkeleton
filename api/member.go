package api

import (
	//"strconv"

	//"github.com/Sirupsen/logrus"
	//"github.com/eurie-inc/echo-sample/model"
	//"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)


func GetMember() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		//number, _ := strconv.ParseInt(c.Param("id"), 0, 64)

		//tx := c.Get("Tx").(*dbr.Tx)

		//member := new(model.Member)
		//if err := member.Load(tx, number); err != nil {
		//	logrus.Debug(err)
		//	return echo.NewHTTPError(fasthttp.StatusNotFound, "Member does not exists.")
		//}
		//return c.JSON(fasthttp.StatusOK, member)
		return c.JSON(fasthttp.StatusOK, "da")
	}
}

