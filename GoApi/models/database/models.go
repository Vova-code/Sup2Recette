package database

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Title      string   `gorm:"column:title"`
	Steps      []string `gorm:"serializer:json"`
	ThumbsUp   int      `gorm:"column:thumbs_up;default:0"`
	ThumbsDown int      `gorm:"column:thumbs_down;default:0"`
}
