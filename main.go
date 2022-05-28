/*
Service start entrance
*/

package main

import (
	"github.com/labstack/echo/v4"

	// change to yourself service-name
	"github.com/lrtxpra/serv-fromework/controllers"
)

// setup your api router here
func registerRouter(e *echo.Echo) {
	// group router
	group := e.Group("/test/v1")
	group.POST("/api1", controllers.Api1)

	e.GET("/api2", controllers.Api2)
}

func main() {
	e := echo.New()

	e.Logger.Fatal(e.Start(":1323"))
}
