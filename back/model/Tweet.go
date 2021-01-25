package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Tweet struct {
	gorm.Model
	Time          time.Time  `json:"time"`
	Text          string     `json:"text" validate:"min=1,max=250"`
	Owner         User       `json:"owner"`
	OwnerID       uint       `json:"ownerID"`
	// Media        
	Likes         []User     `json:"likes"`
	Retweets      []User     `json:"retweets"`
}