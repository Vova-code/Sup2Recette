package database

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Title       string   `gorm:"column:title"`
	Steps       []string `gorm:"serializer:json"`
	Evaluations int      `gorm:"column:evaluations;default:0"`
}
