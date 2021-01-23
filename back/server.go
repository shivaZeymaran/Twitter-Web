package main

import (
	"github.com/labstack/echo/v4"
)


func main() {
	e := echo.New()

	e.GET("/customers", handler.Customer{}.GetorFind)

	// e.POST("/customers", handler.Customer{}.Create)
	// e.PUT("/customers/:cID", handler.Customer{}.Edit)
	// e.DELETE("/customers/:cID", handler.Customer{}.Delete)
	// e.GET("/report/:month", handler.Customer{}.Report)

	if err := e.Start("0.0.0.0:8080"); err != nil {
		fmt.Println("Server not connected!")
	}
}
