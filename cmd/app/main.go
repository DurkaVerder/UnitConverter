package main

import (
	"UnitConverter/interval/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	//Init Server
	e := echo.New()

	//Run Server
	
	handler.Run(e)
}
