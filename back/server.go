package main

import (
	"net/http"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/shivaZeymaran/Twitter-Web.git/handler"
	"github.com/shivaZeymaran/Twitter-Web.git/database"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	PORT = "8090"
)

func main() {
	
	// Database configuration and connection
	router := mux.NewRouter()
	database.ConnectToDB()
	database.AutoMigrate()
	fmt.Println("Successfully connected to database!")
	
	e := echo.New()

	e.POST("/signup", handler.User{}.CreateUser)

	// e.PUT("/customers/:cID", handler.Customer{}.Edit)
	// e.DELETE("/customers/:cID", handler.Customer{}.Delete)
	// e.GET("/report/:month", handler.Customer{}.Report)
	
	if err := e.Start("0.0.0.0:" + PORT); err != nil {
		fmt.Println("Server not connected!")
	}

	handlers := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":" + PORT, handlers))
}
