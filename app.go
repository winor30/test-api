package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/winor30/test-api/db"
	"github.com/winor30/test-api/entity"
	"github.com/winor30/test-api/service"
)

func main() {
	e := echo.New()
	// Routing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!!!!")
	})
	e.GET("/users/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.JSON(http.StatusOK, service.GetUser(name))
	})
	e.POST("/users", func(c echo.Context) error {
		u := new(entity.User)
		if err := c.Bind(u); err != nil {
			return err
		}
		service.SaveUser(u)
		return c.JSON(http.StatusOK, u)
	})

	db.CreateClient()

	e.Logger.Fatal(e.Start(":3000"))
}
