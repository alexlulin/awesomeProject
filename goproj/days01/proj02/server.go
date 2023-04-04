package main

import (
	"github.com/labstack/echo"
	"net/http"
)
var name string
var avatar string

func main(){
	e:=echo.New()
	e.GET("/",func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world! adf")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
