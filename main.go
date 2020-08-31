package main

import (
	"github.com/labstack/echo"
	"github.com/surari/favoritetagger/handler"
)

func main() {
	e := echo.New()
	e.GET("/favorite", handler.Favorite)
	e.Logger.Fatal(e.Start(":1323"))
}
