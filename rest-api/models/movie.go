package models

import (
	"gorm.io/gorm"
)

// the beautiful structured data we want
type Movie struct {
	gorm.Model
	Title  string   `json:"title" binding:"required" gorm:"unique"`
	Year   int16    `json:"year" binding:"required"`
	Genres Genres   `json:"genres" gorm:"serializer:json;type:text"`
	Actors Actors   `json:"cast" gorm:"serializer:json;type:text"`
}

type Genres []string

type Actors []string