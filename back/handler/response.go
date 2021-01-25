package handler

import (
	"time"
	
	"github.com/labstack/echo/v4"
	"github.com/shivaZeymaran/Twitter-Web.git/model"
	"github.com/shivaZeymaran/Twitter-Web.git/utils"
)




/******************************* Signup & Login ********************************/
type SignupResp struct {
	User struct {
		Username string  `json:"username"`
		Email    string  `json:"email"`
		Image    *string `json:"image"`
		Token    string  `json:"token"`
	} `json:"user"`
}

func newUserResponse(u *model.User) *SignupResp {
	r := new(SignupResp)
	r.User.Username = u.Username
	r.User.Email = u.Email
	r.User.Image = u.Image
	r.User.Token = utils.GenerateJWT(u.ID)
	user_token_map[r.User.Token] = r.User.Username
	return r
}


/************************************ Tweet ************************************/
type TweetResp struct {
	Time          time.Time  `json:"time"`
	Text          string     `json:"text" validate:"min=1,max=250"`
	// Media        
	Liked         bool      `json:"liked"`
	LikesCount    int       `json:"likesCount"`
	Retweeted     bool      `json:"retweeted"`
	RetweetsCount int       `json:"retweetsCount"`
	Owner         struct {
		Username  string    `json:"username"`
		Image     *string   `json:"image"`
		Following bool      `json:"following"`
	} `json:"owner"`
}

func newTweetResponse(c echo.Context, t *model.Tweet) *TweetResp {
	r := new(TweetResp)
	r.Time = t.Time
	r.Text = t.Text
	for _, u := range t.Likes {
		if u.ID == userIDFromToken(c) {
			r.Liked = true
		}
	}
	r.LikesCount = len(t.Likes)

	for _, u := range t.Retweets {
		if u.ID == userIDFromToken(c) {
			r.Retweeted = true
		}
	}
	r.RetweetsCount = len(t.Retweets)

	r.Owner.Username = t.Owner.Username
	r.Owner.Image = t.Owner.Image
	r.Owner.Following = t.Owner.FollowedBy(userIDFromToken(c))
	return r
}