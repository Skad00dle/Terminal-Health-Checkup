package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miniProject/models"
	DbUtility "miniProject/pkg"
	"net/http"
)

func FetchTerminalHealth(ctx *gin.Context)  {
	fmt.Println("fetching all terminals")
	var terminalHealths []models.TerminalHealth
	DbUtility.DB.Where("terminal_id = ? ",ctx.Param("id") ).Find(&terminalHealths)
	if len(terminalHealths) <= 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No terminals found!"})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": terminalHealths})
	}

}
