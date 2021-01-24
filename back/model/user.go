package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model          
	Username string 	`json:"username" validate:"min=3,max=40,regexp=^[a-zA-Z]*$"`
	Email string		`json:"email" validate:"regexp=^[^\.][^@\s]+@[^@\s]+\.[^@\s\.]+$"`  // todo: doesn't work!
	Password string 	`json:"password" validate:"min=8"`
}



// cd .\back
// go run server.go