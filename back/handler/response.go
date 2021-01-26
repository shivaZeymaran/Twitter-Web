package handler

import (
	"time"
	
	"github.com/labstack/echo/v4"
	"github.com/shivaZeymaran/Twitter-Web.git/model"
	"github.com/shivaZeymaran/Twitter-Web.git/utils"
	"github.com/shivaZeymaran/Twitter-Web.git/database"
	"github.com/jinzhu/gorm"
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

type SingleTweetResp struct {
	Tweet *TweetResp `json:"tweet"`
}

func newTweetResponse(c echo.Context, t *model.Tweet) *SingleTweetResp {
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
	return &SingleTweetResp{r}
}

/************************************ Timeline ************************************/
type TimelineResp struct {  // list of tweets
	Timeline []*TweetResp `json:"timeline"`
}

func TimelineResponse(c echo.Context, tl []model.Tweet) *TimelineResp {
	r := new(TimelineResp)
	r.Timeline = make([]*TweetResp, 0)
	for _, t := range tl {  // for each tweet in timeline
		tr := new(TweetResp)
		tr.Time = t.Time
		tr.Text = t.Text
		for _, u := range t.Likes {
			if u.ID == userIDFromToken(c) {
				tr.Liked = true
			}
		}
		tr.LikesCount = len(t.Likes)
	
		for _, u := range t.Retweets {
			if u.ID == userIDFromToken(c) {
				tr.Retweeted = true
			}
		}
		tr.RetweetsCount = len(t.Retweets)
	
		tr.Owner.Username = t.Owner.Username
		tr.Owner.Image = t.Owner.Image
		tr.Owner.Following = t.Owner.FollowedBy(userIDFromToken(c))

		r.Timeline = append(r.Timeline, tr)
	}
	return r
}


/******************************* Edit Profile ********************************/
type EditResp struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Image    *string `json:"image"`
}

func EditResponse(u *model.User, token string) *EditResp {
	r := new(EditResp)
	r.Username = u.Username
	r.Email = u.Email
	r.Image = u.Image
	user_token_map[token] = r.Username
	return r
}

/********************************* Follow ***********************************/
type FollowResp struct {
	Username  string  `json:"username"`
	Image     *string `json:"image"`
	Following  bool   `json:"following"`
}

func newFollowResponse(userID uint, u *model.User) *FollowResp {
	r := new(FollowResp)
	r.Username = u.Username
	r.Image = u.Image
	r.Following, _ = IsFollower(u.ID, userID)
	return r
}


func IsFollower(userID, followerID uint) (bool, error) {
	var f model.Follow
	if err := database.DB.Where("following_id = ? AND follower_id = ?", userID, followerID).Find(&f).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}