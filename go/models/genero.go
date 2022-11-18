package models

import (
	"time"

	"gorm.io/gorm"
)

type Genero struct {
	ID        uint `gorm:"primaryKey"`
	Descricao string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
