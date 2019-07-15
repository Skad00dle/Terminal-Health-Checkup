package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"miniProject/crons"
	"miniProject/pkg"
	"miniProject/routes"
)


/*
|****************************************************************************************************************
|						Init Start
|****************************************************************************************************************
*/

func init() {
	////open a db connection
	pkg.InitiateDB()
	//
	////initiate model migrations
	pkg.MigrateDB()

	//
	//starting crons
	crons.InititateCrons()

	//initiate routes
	routes.InitiateRoutes()





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


