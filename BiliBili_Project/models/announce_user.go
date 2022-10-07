package models

import "gorm.io/gorm"

type AnnounceUser struct {
	gorm.Model
	Aid    uint `gorm:"not null"`
	Uid    uint `gorm:"not null"`
	Statue int  `gorm:"default:0"`
}
