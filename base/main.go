package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func main() {

	router := gin.Default()

	v1 := router.Group("/terminal")
	{
		v1.POST("/", addTerminal)
		v1.GET("/", fetchTerminals)
		//v1.GET("/:id", fetchSingleTodo)
		//v1.PUT("/:id", updateTodo)
		//v1.DELETE("/:id", deleteTodo)
	}

	router.Run()


}

func addTerminal(ctx *gin.Context)  {     // add a new terminal into db


	var terms []Terminal

	ctx.Bind(&terms)
	go createTerminals(terms)
	fmt.Println("bana diya")
}

var db *gorm.DB

func init() {
	//open a db connection
	// TODO initialise db connection from db package

	var err error
	db, err = gorm.Open("mysql", "root:admin@/golangproject?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}else {
		fmt.Println("connection to db is successful")
	}
	//Migrate the schema
	db.AutoMigrate(&Terminal{})

}

func ownServerStatus()  {
	fmt.Println("server ready and running")
}

type (
	Terminal struct {      // terminal model
		gorm.Model   // this will automatically give us id,created at , updated at , deleted at
		Url			 string			`json:"url"`
		Timeout		 int			`json:"timeout"`
		Frequency	 int			`json:"frequency"`
		Threshold	 int			`json:"threshold"`
	}
)

func createTerminals(terms []Terminal)  {
	fmt.Println("terminals: ",terms)
	for _,term := range terms{
		test := Terminal{}
		db.Where("url = ? ", fmt.Sprintf(term.Url)).First(&test)
		if (test.Url == term.Url) {
			test.Timeout 	= term.Timeout
			test.Threshold 	= term.Threshold
			test.Frequency 	= term.Frequency
			db.Save(&test)
			fmt.Println("Updated one terminal")
		}else {
			db.Save(&term)
			fmt.Println("added a new terminal")
		}
	}
}

func fetchTerminals(ctx *gin.Context)  {
	fmt.Println("fetching all terminals")
	var terminals []Terminal
	db.Find(&terminals)
	if len(terminals) <= 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": terminals})
	}

}

