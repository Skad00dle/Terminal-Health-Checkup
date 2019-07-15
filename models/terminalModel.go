package models

import "github.com/jinzhu/gorm"

type (
	Terminal struct {      // terminal model
		gorm.Model   // this will automatically give us id,created at , updated at , deleted at
		Url			 string			`json:"url"`
		Timeout		 int			`json:"timeout"`
		Frequency	 int			`json:"frequency"`
		Threshold	 int			`json:"threshold"`
	}
)