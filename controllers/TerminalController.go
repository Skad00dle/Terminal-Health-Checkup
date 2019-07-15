package controllers

import (
	models "../models"
	DbUtility "../pkg"
	Services "../services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddTerminal(ctx *gin.Context)  {     // add a new terminal into db


	var terms []models.Terminal

	ctx.Bind(&terms)
	go Services.CreateTerminals(terms)
	fmt.Println("terminal handled")
	// TODO return something here
}


func FetchTerminals(ctx *gin.Context)  {
	fmt.Println("fetching all terminals")
	var terminals []models.Terminal
	DbUtility.DB.Find(&terminals)
	if len(terminals) <= 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No terminals found!"})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": terminals})
	}

}