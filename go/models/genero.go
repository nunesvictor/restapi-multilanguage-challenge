package models

import (
	"time"
)

type Genero struct {
	ID        uint   `gorm:"primaryKey"`
	Descricao string `gorm:"not null;default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
