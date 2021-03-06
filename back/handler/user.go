package handler

import (
	"net/http"
	"sort"
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
		// Add to DB (each user is his/her own follower)
		u.Followers = make([]model.Follow, 0)
		u.Tweets = make([]model.Tweet, 0)
		// u.Likes = make([]model.Tweet, 0)
		database.DB.Model(&u).Association("Followers").Append(&model.Follow{FollowerID: u.ID, FollowingID: u.ID})
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
	// t.Likes = make([]model.User, 0)

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
		return c.JSON(http.StatusNotFound, "Tweet Not found for user")
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
			fmt.Println(req.Username)
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
	// make new model from SimpleRequest
	req := &SimpleReq {}

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
	// make new model from SimpleRequest
	req := &SimpleReq {}

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


// Timeline godoc
// @Summary Timeline of a user
// @Description Display recent tweets sorted by their time in timeline
// @ID timeline
// @Accept  json
// @Produce  json
// @Router /timeline [post]
func (user User) Timeline(c echo.Context) error {
	// make new model from SimpleRequest
	req := &SimpleReq {}

	// Bind given model to request struct
	if err := req.bind(c); err != nil { // Not successful
		return  c.JSON(http.StatusUnprocessableEntity, err)
	}

	// get user ID from token
	username := user_token_map[req.Token]
	var u model.User
	database.DB.Find(&u, model.User{Username:username}) // user that going to display his/her timeline

	var f = make([]model.Follow, 0) // all followers of current user
	if err := database.DB.Where(&model.Follow{FollowerID: u.ID}).Find(&f).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "User does not follow anyone yet!")
	}

	var tl = make([]model.Tweet, 0)  // Timeline tweets
	for _, fo := range f {
		var t = make([]model.Tweet, 0)  // each following user tweets		
		var fu model.User
		database.DB.Find(&fu, fo.FollowingID)
		if err := database.DB.Model(&fu).Association("Tweets").Find(&t).Error; err != nil {
			return c.JSON(http.StatusNotFound, "Can't retreive tweets of User " + fo.Following.Username)
		}
		for _, tw := range t {
			tw.Owner = fu
			tl = append(tl, tw)
		}
	}
	// sort timeline tweets by their time
	sort.SliceStable(tl, func(i, j int) bool {
		return tl[i].Time.After(tl[j].Time)
	})
	return c.JSON(http.StatusOK, TimelineResponse(c, tl))
}


// Search User godoc
// @Summary Search a user
// @Description Find a user by searching it's username (user may logged in or not)
// @ID search-user
// @Accept  json
// @Produce  json
// @Router /search@/{username} [get]
func (user User) SearchUser(c echo.Context) error {
	username := c.Param("username")
	var u model.User
	if err := database.DB.Find(&u, model.User{Username:username}).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Username " + username + " does not exist!")
	} 
	return c.JSON(http.StatusOK, SearchUserResponse(&u))
}


// Search User godoc
// @Summary Search a user
// @Description Find a user by searching it's username (user may logged in or not)
// @ID search-user
// @Accept  json
// @Produce  json
// @Router /search@/{username} [post]
func (user User) SearchUserWithLogin(c echo.Context) error {
	// make new model from SimpleRequest
	req := &SimpleReq {}

	// Bind given model to request struct
	if err := req.bind(c); err != nil { // Not successful
		return  c.JSON(http.StatusUnprocessableEntity, err)
	}

	// get user ID from token
	uname := user_token_map[req.Token]
	var u model.User
	database.DB.Find(&u, model.User{Username:uname})
	uid := u.ID

	username := c.Param("username")
	var su model.User  // user that be searched by its username
	if err := database.DB.Find(&su, model.User{Username:username}).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Username " + username + " does not exist!")
	} 
	return c.JSON(http.StatusOK, newFollowResponse(uid, &su))
}


// Like godoc
// @Summary Like a tweet
// @Description Like a tweet
// @ID like
// @Accept  json
// @Produce  json
// @Router /like [post]
func (user User) LikeTweet(c echo.Context) error {
	// make new model from LikeRequest
	req := &LikeReq {}

	// Bind given model to request struct
	if err := req.bind(c); err != nil { // Not successful
		return  c.JSON(http.StatusUnprocessableEntity, err)
	}

	// get user ID from token
	username := user_token_map[req.Token]
	var u model.User  // user that wanna like a tweet
	database.DB.Find(&u, model.User{Username:username})

	// get user ID from given username (owner of tweet)
	var ou model.User  // owner user
	database.DB.Find(&ou, model.User{Username:req.OwnerUsername})

	// identify Tweet from tweet text and owner id
	var t model.Tweet
	// Add to DB
	if err := database.DB.Where(&model.Tweet{Text: req.Text, OwnerID: ou.ID}).Find(&t).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "Tweet Not found for user")
	}
	t.Likes = make([]model.User, 1)
	database.DB.Model(&t).Association("Likes").Append(&u)

	t.Owner = ou
	return c.JSON(http.StatusOK, newTweetResponse(c, &t))
}


// Logout godoc
// @Summary Logout from current user
// @Description Logout from current user
// @ID logout
// @Accept  json
// @Produce  json
// @Router /logout [post]
func (user User) Logout(c echo.Context) error{
	// make new model from SimpleRequest
	req := &SimpleReq {}

	// Bind given model to request struct
	if err := req.bind(c); err != nil { // Not successful
		return  c.JSON(http.StatusUnprocessableEntity, err)
	}

	delete(user_token_map, req.Token) // remove element user_token_map[req.Token] from map of tokens
	
	return c.JSON(http.StatusOK, "You logged out successfully!")

}


func userIDFromToken(c echo.Context) uint {
	id, ok := c.Get("user").(uint)
	if !ok {
		return 0
	}
	return id
}
