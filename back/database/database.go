package database

import (
	"fmt"

	"github.com/shivaZeymaran/Twitter-Web.git/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "2282650166"
	DBNAME   = "twitter_database"
)

var db *gorm.DB

func ConnectToDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME)
	db, _ = gorm.Open("postgres", psqlInfo)
	db.DB().SetMaxIdleConns(3)
}

func AutoMigrate() {
	db.AutoMigrate(
		&model.User{},
		// &model.Follow{},
		// &model.Article{},
		// &model.Comment{},
		// &model.Tag{},
	)
}