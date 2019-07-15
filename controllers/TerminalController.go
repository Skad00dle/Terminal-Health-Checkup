package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addTerminal(ctx *gin.Context)  {     // add a new terminal into db


	var terms []Terminal

	ctx.Bind(&terms)
	go createTerminals(terms)
	fmt.Println("bana diya")
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