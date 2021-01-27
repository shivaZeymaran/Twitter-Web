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
	"github.com/labstack/echo/v4/middleware"
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

	// CORS restricted
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.POST("/signup", handler.User{}.Signup)
	e.POST("/login", handler.User{}.Login)
	// todo: go to home page for user: output should be like login
	// todo: search when not loged in
	e.POST("/tweet", handler.User{}.Tweet)
	e.DELETE("/deletetweet", handler.User{}.DeleteTweet)
	e.PUT("/editprofile", handler.User{}.EditProfile)
	e.POST("/follow/:username", handler.User{}.Follow)
	e.DELETE("/unfollow/:username", handler.User{}.UnFollow)
	e.POST("/timeline", handler.User{}.Timeline)
	
	if err := e.Start("0.0.0.0:" + PORT); err != nil {
		fmt.Println("Server not connected!")
	}

	handlers := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":" + PORT, handlers))
}