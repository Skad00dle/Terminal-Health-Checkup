package pkg

import (
	"fmt"
	"miniProject/models"
)

func MigrateDB()  {
	fmt.Println("starting migrations")
	DB.AutoMigrate(&models.Terminal{})
	DB.AutoMigrate(&models.TerminalHealth{})
	DB.AutoMigrate(&models.TerminalHealthHit{})
	fmt.Println("migrations complete")
}

