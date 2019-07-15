package services

import (
	"../models"
	DbUtility "../pkg"
	"fmt"
)

func CreateTerminals(terms []models.Terminal)  {
	for _,term := range terms{
		test := models.Terminal{}
		DbUtility.DB.Where("url = ? ", fmt.Sprintf(term.Url)).First(&test)
		if (test.Url == term.Url) {
			test.Timeout 	= term.Timeout
			test.Threshold 	= term.Threshold
			test.Frequency 	= term.Frequency
			DbUtility.DB.Save(&test)
			fmt.Println("Updated one terminal")
		}else {
			DbUtility.DB.Save(&term)
			fmt.Println("added a new terminal")
		}
	}
}