package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	

	"github.com/labstack/echo/v4"
	"github.com/shivaZeymaran/Twitter-Web.git/model"
	"github.com/shivaZeymaran/Twitter-Web.git/database"
	"gopkg.in/validator.v2"
)


// For binding handler methods
type User struct {

}


func (user User) CreateUser(c echo.Context) error{
	// make new model from User
	u := &model.User{}
	
	// Bind given model to User struct
	if err := c.Bind(&u); err != nil { // Not successful
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request!")
	}

	// Validate the request values
	if errs := validator.Validate(u); errs != nil {
		return c.JSON(http.StatusBadRequest, errs)
	}

	// Search in DB
	var findUser model.User
	database.DB.Find(&findUser, model.User{Username: u.Username})
	if findUser.Username == "" {  // Successfully created
		database.DB.Create(&u)
		newU, _ := json.Marshal(u)
		fmt.Println(string(newU))
		return c.JSON(http.StatusCreated, "Dear "+ u.Username +", you have signed up successfully!")
	}
	// Username currently exists
	return c.JSON(http.StatusBadRequest, "Username " + u.Username + " currently exists!")
}