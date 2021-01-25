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

var DB *gorm.DB

func ConnectToDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME)
	DB, _ = gorm.Open("postgres", psqlInfo)
	DB.DB().SetMaxIdleConns(10)
}

func AutoMigrate() {
	DB.AutoMigrate(
		&model.User{},
		&model.Tweet{},
	)
}