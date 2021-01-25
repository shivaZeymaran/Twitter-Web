package handler

import (
	// "encoding/json"
	// "fmt"
	"net/http"
	// "time"
	
	"github.com/labstack/echo/v4"
	"github.com/shivaZeymaran/Twitter-Web.git/model"
	"github.com/shivaZeymaran/Twitter-Web.git/database"
)

var user_token_map map[string]string = initMap()

func initMap() map[string]string{
	return make(map[string]string)
}

// For binding handler methods
type User struct {

}

// SignUp godoc
// @Summary Register a new user
// @Description Register a new user
// @ID signup
// @Accept  json
// @Produce  json
// @Router /signup [post]
func (user User) Signup(c echo.Context) error{
	
	// make new model from User
	u := &model.User {}
	
	// Bind given model to Sign up request struct and check validations
	req := &SignupReq {}
	if err := req.bind(c, u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	// Search in DB
	var findUser model.User
	database.DB.Find(&findUser, model.User{Username: u.Username})
	if findUser.Username == "" {  // Successfully created
		database.DB.Create(&u)
		return c.JSON(http.StatusCreated, newUserResponse(u))
	}
	// Username currently exists
	return c.JSON(http.StatusBadRequest, "Username " + u.Username + " already exists!")
}


// Login godoc
// @Summary Login for existing user
// @Description Login for existing user
// @ID login
// @Accept  json
// @Produce  json
// @Router /login [post]
func (user User) Login(c echo.Context) error{
	// make new model from LoginRequest
	req := &LoginReq {}

	// Bind given model to request struct
	if err := req.bind(c); err != nil { // Not successful
		return  c.JSON(http.StatusUnprocessableEntity, err)
	}

	// Search in DB
	var u model.User
	database.DB.Find(&u, model.User{Username: req.Username})
	if u.Username != "" {  // Successfully find user
		if u.CheckPassword(req.Password) {  // Password is true, user authorized
			return c.JSON(http.StatusOK, newUserResponse(&u))
		}
		// password was wrong
		return c.JSON(http.StatusForbidden, "Password is Wrong!")
	}
	// User does not exist
	return c.JSON(http.StatusNotFound, "Username " + req.Username + " does not exist!")

}


// Tweet godoc
// @Summary Create a tweet
// @Description Create a tweet. Owner is require
// @ID tweet
// @Accept  json
// @Produce  json
// @Router /tweet [post]
func (user User) Tweet(c echo.Context) error {
	t := &model.Tweet {}

	req := &TweetReq {}
	if err := req.bind(c, t); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	t.OwnerID = userIDFromToken(c)

	// add to DB
	database.DB.Create(&t)

	return c.JSON(http.StatusCreated, newTweetResponse(c, t))
}


// EditProfile godoc
// @Summary Edit users profile
// @Description Edit the profile's image and username
// @ID Edit
// @Accept  json
// @Produce  json
// @Router /editprofile [put]
func (user User) EditProfile(c echo.Context) error {
	// make new model from EditRequest
	req := &EditReq {}

	// Bind given model to request struct
	if err := req.bind(c); err != nil { // Not successful
		return  c.JSON(http.StatusUnprocessableEntity, err)
	}

	username := user_token_map[req.Token]
	var u model.User
	database.DB.Find(&u, model.User{Username:username}) // user that wanna edit profile

	database.DB.Model(&u).Update("Image", req.Image) // edit image anyway

	if req.Username != username {  // wanna change username
		// Search in DB for new username
		var findUser model.User
		database.DB.Find(&findUser, model.User{Username: req.Username})
		if findUser.Username == "" {  // Successfully edited username
			database.DB.Model(&u).Update("Username", req.Username)
			return c.JSON(http.StatusCreated, EditResponse(&u, req.Token))
		}
		// Username currently exists
		return c.JSON(http.StatusBadRequest, "Username " + u.Username + " already exists!")
	}
	
	return c.JSON(http.StatusCreated, EditResponse(&u, req.Token))
}

func userIDFromToken(c echo.Context) uint {
	id, ok := c.Get("user").(uint)
	if !ok {
		return 0
	}
	return id
}