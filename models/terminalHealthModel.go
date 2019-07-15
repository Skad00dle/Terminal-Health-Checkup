package models

import "github.com/jinzhu/gorm"

type (
	TerminalHealth struct {
		gorm.Model
		TerminalId	uint
		Result 		int						// 0 down, 1 up, 2 wrong url
	}
)
