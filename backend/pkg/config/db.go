package config

import (
	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Connect(){
	d,err := gorm.Open("mysql",utils.GetEnvVariable("DB_URL"))
	if err != nil {
		panic(err)
	}
	DB = d
}

func GetDB() *gorm.DB{
	return DB
}