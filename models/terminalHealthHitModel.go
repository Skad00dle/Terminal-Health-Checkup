package models

import "github.com/jinzhu/gorm"

type (
	TerminalHealthHit struct {
		gorm.Model
		TerminalHealthId	uint
		Result 				int						// -1 for error else status code
	}
)
