package model


type User struct{
	ID int				`json:"cID"`
	Name string  		`json:"cName" validate:"nonzero"`
	Tel int      		`json:"cTel" validate:"nonzero"`
	Address string  	`json:"cAddress" validate:"nonzero"`
	RegisterDate string `json:"cRegisterDate"`
}