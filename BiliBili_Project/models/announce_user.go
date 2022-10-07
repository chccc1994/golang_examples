package models

type AnnounceUser struct {
	Aid    uint `gorm:"not null"`
	Uid    uint `gorm:"not null"`
	Statue int  `gorm:"default:0"`
}
