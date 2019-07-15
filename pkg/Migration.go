package pkg

import (
	"../models"
)

func MigrateDB()  {
	DB.AutoMigrate(&models.Terminal{})
	DB.AutoMigrate(&models.TerminalHealth{})
	DB.AutoMigrate(&models.TerminalHealthHit{})
}

