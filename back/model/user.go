package model

import (
	"errors"
	
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model          
	Username 	string 	 	`json:"username" validate:"min=3,max=40,regexp=^[a-zA-Z]*$"`
	Email 		string	 	`json:"email" validate:"regexp=^[^\.][^@\s]+@[^@\s]+\.[^@\s\.]+$"`  // todo: doesn't work!
	Password 	string 	   	`json:"password" validate:"min=8"`
	Image       *string     `json:"image"`
	Followers   []Follow   	`json:"followers"`
	Followings  []Follow   	`json:"followings"`
	// Timeline    []Tweet     `json:"timeline"`
}

type Follow struct {
	Follower    User  `json:"follower"`
	FollowerID  uint  `json:"followerUsername"`
	Following   User  `json:"following"`
	FollowingID uint  `json:"followingUsername"`	
}

func (u User) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (u User) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}

// FollowedBy Followings should be pre loaded
func (u *User) FollowedBy(id uint) bool {
	if u.Followers == nil {   // no one follows this user so as you :)
		return false
	}
	for _, f := range u.Followers {
		if f.FollowerID == id {
			return true
		}
	}
	return false
}

// cd .\back
// go run server.go

// git add .
// git pull origin master
// git commit -m "
// git push origin master
