package models

import (
	"time"
)

type Editora struct {
	ID         uint   `gorm:"primaryKey"`
	Nome       string `gorm:"not null;default:null"`
	Localidade string `gorm:"not null;default:null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
