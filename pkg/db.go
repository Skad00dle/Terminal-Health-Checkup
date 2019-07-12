package pkg

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func initiateDB()  {

	var db *gorm.DB

	var err error
	db, err = gorm.Open("mysql", "root:admin@/golangproject?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}else {
		fmt.Println("connection to db is successful")
	}
	//Migrate the schema
	db.AutoMigrate(&TerminalModel{})

}

