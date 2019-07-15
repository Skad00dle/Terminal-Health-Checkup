package pkg

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

var DB *gorm.DB

func InitiateDB()  {
	var err error
	DB, err = gorm.Open("mysql", "root:admin@/golangproject?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}else {
		fmt.Println("connection to db is successful")
	}
}
