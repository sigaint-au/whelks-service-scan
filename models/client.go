package models

import (
	"time"
	"gorm.io/gorm"
)

type IP struct {
	gorm.Model
	ID			uint
	Address		string
	Scan    	string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Sessions	[]Session
}


