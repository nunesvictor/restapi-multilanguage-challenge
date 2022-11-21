package models

import (
	"time"
)

type Genero struct {
	ID        uint `gorm:"primaryKey"`
	Descricao string
	CreatedAt time.Time
	UpdatedAt time.Time
}
