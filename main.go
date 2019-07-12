package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	router := gin.Default()

	v1 := router.Group("/terminal")
	{
		v1.POST("/", addTerminal)
		//v1.GET("/", fetchAllTodo)
		//v1.GET("/:id", fetchSingleTodo)
		//v1.PUT("/:id", updateTodo)
		//v1.DELETE("/:id", deleteTodo)
	}

	router.Run()


}

func addTerminal(ctx *gin.Context)  {     // add a new terminal into db

	var term Terminal
	ctx.BindJSON(&term)      // storing data received into variable "term"

	//fmt.Println("json? ",term)
	//timeout, _ := strconv.Atoi(ctx.PostForm("timeout"))
	//frequency, _ := strconv.Atoi(ctx.PostForm("frequency"))
	//threshold, _ := strconv.Atoi(ctx.PostForm("threshold"))
	//terminal := Terminal{
	//	Url: ctx.PostForm("url"),
	//	Timeout: timeout,
	//	Frequency: frequency,
	//	Threshold: threshold,
	//}
	//fmt.Println(terminal)
	//db.Save(&terminal)

	db.Save(&term)
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

