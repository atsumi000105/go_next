package models

import (
	"time"

	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/config"
	"gorm.io/gorm"
)

type User struct{
	ID 			uint   		 		`json:"id"`
	Name     	string 		 		`json:"name"`
	Email  		string 		 		`json:"email"`
	Password 	string 		 		`json:"password"`
	Token 	 	string	 		 	`json:"token"`
	CreatedAt 	time.Time      	`json:"created_at"`
	UpdatedAt 	time.Time 			`json:"updated_at"`
  	DeletedAt 	gorm.DeletedAt 	`gorm:"index" json:"deleted_at"`
}

func init(){
	config.Connect()
	DB = config.GetDB()
	DB.AutoMigrate(&User{})
}