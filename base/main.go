package main

import (
	CronUtility "../crons"
	DbUtility "../pkg"
	Routes "../pkg"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


/*
|****************************************************************************************************************
|						Init Start
|****************************************************************************************************************
*/

func init() {
	//open a db connection
	DbUtility.InitiateDB()

	//initiate model migrations
	DbUtility.MigrateDB()

	//initiate routes
	Routes.InitiateRoutes()


	//starting crons
	CronUtility.InititateCrons()



}

/*
|****************************************************************************************************************
|						Init End
|****************************************************************************************************************
*/


func main() {

	ownServerStatus()


}



func ownServerStatus()  {
	fmt.Println("server ready and running")
}


