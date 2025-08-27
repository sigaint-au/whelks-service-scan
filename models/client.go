package models

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ID			uint
	Address		string
	Scan    	string
	Sessions	[]Session
}


