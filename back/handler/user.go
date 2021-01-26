package handler

import (
	"net/http"
	"fmt"

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
		return c.JSON(http.StatusOK, err)
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


// Delete Tweet godoc
// @Summary Delete a tweet
// @Description Delete a tweet
// @ID delete-tweet
// @Accept  json
// @Produce  json
// @Router /deletetweet [delete]
func (user User) DeleteTweet(c echo.Context) error {
	// make new model from DeleteTweetRequest
	req := &DeleteTweetReq {}

	// Bind given model to request struct
	if err := req.bind(c); err != nil { // Not successful
		return  c.JSON(http.StatusUnprocessableEntity, err)
	}

	// get user ID from token
	username := user_token_map[req.Token]
	var u model.User
	database.DB.Find(&u, model.User{Username:username}) // user that wanna delete his/her tweet

	// identify Tweet from tweet text and ownerID
	var t model.Tweet
	// Delete from DB
	if err := database.DB.Where(&model.Tweet{Text: req.Text, OwnerID: u.ID}).Find(&t).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "Tweet Not found for user")
	}
	
	// if err := database.DB.Model(&u).Association("Tweets").Find(&t).Error; err != nil {
	if err := database.DB.Delete(t).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "Tweet successfully deleted!")
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


// Follow godoc
// @Summary Follow a user
// @Description Follow a user by username
// @ID follow
// @Accept  json
// @Produce  json
// @Router /follow/{username} [post]
func (user User) Follow(c echo.Context) error {
	// make new model from FollowRequest
	req := &FollowReq {}

	// Bind given model to request struct
	if err := req.bind(c); err != nil { // Not successful
		return  c.JSON(http.StatusUnprocessableEntity, err)
	}

	// get user ID from token
	un := user_token_map[req.Token]
	var ru model.User
	database.DB.Find(&ru, model.User{Username:un})
	followerID := ru.ID
	
	username := c.Param("username")
	var u model.User
	database.DB.Find(&u, model.User{Username:username}) // user that wanna be followed
	if u.Username == "" {  // chosen user not exists
		return c.JSON(http.StatusNotFound, "Username " + u.Username + "does not exist!")
	}
	// Add to DB
	database.DB.Model(&u).Association("Followers").Append(&model.Follow{FollowerID: followerID, FollowingID: u.ID})

	return c.JSON(http.StatusOK, newFollowResponse(followerID, &u))
}


// Unfollow godoc
// @Summary Unfollow a user
// @Description Unfollow a user by username
// @ID unfollow
// @Accept  json
// @Produce  json
// @Router /unfollow/{username} [delete]
func (user User) UnFollow(c echo.Context) error {
	// make new model from FollowRequest
	req := &FollowReq {}

	// Bind given model to request struct
	if err := req.bind(c); err != nil { // Not successful
		return  c.JSON(http.StatusUnprocessableEntity, err)
	}

	// get user ID from token
	un := user_token_map[req.Token]
	var ru model.User
	database.DB.Find(&ru, model.User{Username:un})
	followerID := ru.ID
	
	username := c.Param("username")
	var u model.User // user that wanna be unfollowed
	if err := database.DB.Find(&u, model.User{Username:username}).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Username " + username + " does not exist!") // chosen user not exists
	} 

	// Delete from DB
	f := model.Follow{
		FollowerID: followerID,
		FollowingID: u.ID,
	}
	if err := database.DB.Model(&u).Association("Followers").Find(&f).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "You did not follow user " + username)
	}
	if err := database.DB.Delete(f).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, newFollowResponse(followerID, &u))
}



func userIDFromToken(c echo.Context) uint {
	id, ok := c.Get("user").(uint)
	if !ok {
		return 0
	}
	return id
}
