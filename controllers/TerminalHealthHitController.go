package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miniProject/models"
	DbUtility "miniProject/pkg"
	"net/http"
)

func FetchTerminalHealthHit(ctx *gin.Context)  {
	fmt.Println("fetching all terminals")
	var terminalHealthHits []models.TerminalHealthHit
	DbUtility.DB.Where("terminal_health_id = ? ",ctx.Param("id") ).Find(&terminalHealthHits)
	if len(terminalHealthHits) <= 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No terminals found!"})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": terminalHealthHits})
	}

}