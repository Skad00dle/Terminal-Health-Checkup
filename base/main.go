package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/robfig/cron"
	"net/http"
	"sync"
	"time"
)


/*
|****************************************************************************************************************
|						Init Start
|****************************************************************************************************************
*/


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
	db.AutoMigrate(&TerminalHealth{})
	db.AutoMigrate(&TerminalHealthHit{})





	c := cron.New()

	c.AddFunc(" */1 * * * *", startHealthCheckup)

	c.Start()





}

/*
|****************************************************************************************************************
|						Init End
|****************************************************************************************************************
*/


func main() {

/*
|****************************************************************************************************************
|						Routes Start
|****************************************************************************************************************
 */

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

/*
|****************************************************************************************************************
|						Routes End
|****************************************************************************************************************
*/


}



func ownServerStatus()  {
	fmt.Println("server ready and running")
}


/*
|****************************************************************************************************************
|						Models Start
|****************************************************************************************************************
*/
type (
	Terminal struct {      // terminal model
		gorm.Model   // this will automatically give us id,created at , updated at , deleted at
		Url			 string			`json:"url"`
		Timeout		 int			`json:"timeout"`
		Frequency	 int			`json:"frequency"`
		Threshold	 int			`json:"threshold"`
	}
)

type (
	TerminalHealth struct {
		gorm.Model
		TerminalId	uint
		Result 		int						// 0 down, 1 up, 2 wrong url
	}
)

type (
	TerminalHealthHit struct {
		gorm.Model
		TerminalHealthId	uint
		Result 				int						// -1 for error else status code
	}
)

/*
|****************************************************************************************************************
|						Models End
|****************************************************************************************************************
*/




/*
|****************************************************************************************************************
|						Model Functions Start
|****************************************************************************************************************
*/


func addTerminal(ctx *gin.Context)  {     // add a new terminal into db


	var terms []Terminal

	ctx.Bind(&terms)
	go createTerminals(terms)
	fmt.Println("bana diya")
}



func createTerminals(terms []Terminal)  {
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
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No terminals found!"})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": terminals})
	}

}

/*
|****************************************************************************************************************
|						Model Functions End
|****************************************************************************************************************
*/





/*
|****************************************************************************************************************
|						Cron Functions Start
|****************************************************************************************************************
*/


var wgCron sync.WaitGroup

func startHealthCheckup()  {
	// TODO add lock file here
	fmt.Println("checking health status")

	var terminals []Terminal
	db.Find(&terminals)

	for _,terminal := range terminals {
		if terminal.Url !="" {
			wgCron.Add(1)
			go checkHealth(terminal)
		} else {
			continue
		}
	}
	wgCron.Wait()
	fmt.Println("Health Check Completed")
}

func checkHealth(term Terminal)  {


	tr := &http.Transport{
		IdleConnTimeout:    time.Duration(term.Timeout) * time.Millisecond,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(term.Url)

	termHealth := TerminalHealth{}
	termHealth.TerminalId = term.ID

	db.Create(&termHealth)

	termHealthHit := TerminalHealthHit{}
	termHealthHit.TerminalHealthId =termHealth.ID

	if(err != nil){
		fmt.Println("got unexpected error")
		fmt.Println(err.Error())
		termHealth.Result = 2  // wrong url
		termHealthHit.Result = -1  //wrong hit
		db.Save(&termHealthHit)
		db.Save(&termHealth)
	}else {
		if(resp.StatusCode == 200){
			fmt.Println("Terminal",term.Url , "is working ")
			termHealth.Result = 1  // working
			termHealthHit.Result = 200  //successful hit
			db.Save(&termHealthHit)
			db.Save(&termHealth)
		}else {
			fmt.Println("Terminal",term.Url , "is not working ")
			termHealthHit.Result = resp.StatusCode  //wrong hit
			db.Save(&termHealthHit)
			time.Sleep(time.Duration(term.Frequency) * time.Millisecond)
			retryHealthHit(term,termHealth,2)
		}
	}

	wgCron.Done()
}

func retryHealthHit(term Terminal, termHealth TerminalHealth, retryCount int)  {
	if(retryCount > term.Threshold){
		termHealth.Result = 0 // not working
		db.Save(&termHealth)
		return
	}

	termHealthHit := TerminalHealthHit{}
	termHealthHit.TerminalHealthId =termHealth.ID

	tr := &http.Transport{
		IdleConnTimeout:    time.Duration(term.Timeout) * time.Millisecond,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(term.Url)

	if(err != nil){
		fmt.Println("got unexpected error")
		fmt.Println(err.Error())
		termHealth.Result = 2  // wrong url
		termHealthHit.Result = -1  //wrong hit
		db.Save(&termHealthHit)
		db.Save(&termHealth)
		return
	}else {
		if(resp.StatusCode == 200){
			fmt.Println("Terminal",term.Url , "is working ")
			termHealth.Result = 1  //  working
			termHealthHit.Result = 200  //successful hit
			db.Save(&termHealthHit)
			db.Save(&termHealth)
			return
		}else {
			fmt.Println("Terminal",term.Url , "is not working ")
			termHealthHit.Result = resp.StatusCode  //wrong hit
			db.Save(&termHealthHit)
			time.Sleep(time.Duration(term.Frequency) * time.Millisecond)
			retryHealthHit(term,termHealth,retryCount+1)
			return
		}
	}
}






/*
|****************************************************************************************************************
|						Cron Functions End
|****************************************************************************************************************
*/
