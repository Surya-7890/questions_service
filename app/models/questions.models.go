package models

import "time"

type Question struct {
	ID        uint    	`gorm:"primaryKey"`
	// CompanyID uint    	`gorm:"not null"`
	Question  string
	Year      time.Time
}
