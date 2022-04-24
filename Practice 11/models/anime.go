package models

import (
	_ "database/sql"
	"time"
)

type Anime struct {
	ID             uint       `json:"id" gorm:"PRIMARY_KEY"`
	Name           string     `json:"name" gorm:"type:varchar(200);UNIQUE;NOT NULL"`
	YearRelease    uint16     `json:"yearRelease" gorm:"NOT NULL"`
	Rating         float32    `json:"ratingRelease" gorm:"NOT NULL;DEFAULT:0.0"`
	ViewsNumber    uint32     `json:"viewsNumber" gorm:"NOT NULL;DEFAULT:0"`
	CreatedAt      time.Time  `json:"createdAt"`
	DeletedAt      *time.Time `json:"deletedAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
	Director       string     `json:"director" gorm:"type:varchar(150);NOT NULL"`
	ReleaseCountry string     `json:"releaseCountry" gorm:"type:varchar(150);NOT NULL"`
	Genre          string     `json:"genre" gorm:"type:varchar(50);NOT NULL"`
}
