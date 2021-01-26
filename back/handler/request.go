package handler

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shivaZeymaran/Twitter-Web.git/model"
	"github.com/shivaZeymaran/Twitter-Web.git/database"
	"gopkg.in/validator.v2"
)

/************************************ Sign up ************************************/
type SignupReq struct {
	Username string `json:"username" validate:"min=3,max=40,regexp=^[a-zA-Z]*$"`
	Email    string `json:"email" validate:"regexp=^[^\.][^@\s]+@[^@\s]+\.[^@\s\.]+$"`
	Password string `json:"password" validate:"min=8"`
}

func (r *SignupReq) bind(c echo.Context, u *model.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if errs := validator.Validate(r); errs != nil {
		return errs
	}
	u.Username = r.Username
	u.Email = r.Email
	h, err := u.HashPassword(r.Password)
	if err != nil {
		return err
	}
	u.Password = h
	return nil
}

/************************************ Log in ************************************/
type LoginReq struct {
	Username string	`json:"username" validate:"min=3,max=40,regexp=^[a-zA-Z]*$"`
	Password string `json:"password" validate:"min=8"`
}

func (r *LoginReq) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if errs := validator.Validate(r); errs != nil {
		return errs
	}
	return nil
}

/************************************ Tweet ************************************/
type TweetReq struct {
	Text  string `json:"text" validate:"min=1,max=250"`
	Token string `json:"token"`
}

func (r *TweetReq) bind(c echo.Context, t *model.Tweet) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if errs := validator.Validate(r); errs != nil {
		return errs
	}
	username := user_token_map[r.Token]
	var u model.User
	database.DB.Find(&u, model.User{Username:username})
	t.Owner = u
	t.Time = time.Now()
	t.Text = r.Text
	return nil
}

/******************************** Delete Tweet **********************************/
type DeleteTweetReq struct {
	Text  string `json:"text"`
	Token string `json:"token"`
}

func (r *DeleteTweetReq) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if errs := validator.Validate(r); errs != nil {
		return errs
	}
	return nil
}

/********************************* Edit Profile **********************************/
type EditReq struct {
	Username  string  `json:"username" validate:"min=3,max=40,regexp=^[a-zA-Z]*$"`
	Image     *string `json:"image"` 
	Token     string  `json:"token"`
}

func (r *EditReq) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if errs := validator.Validate(r); errs != nil {
		return errs
	}
	return nil
}

/******************************** Follow & UnFollow ********************************/
type FollowReq struct {
	Token  string  `json:"token"`
}

func (r *FollowReq) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if errs := validator.Validate(r); errs != nil {
		return errs
	}
	return nil
}
