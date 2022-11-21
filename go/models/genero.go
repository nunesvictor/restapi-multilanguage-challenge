package models

import (
	"gin-restapi/schemas"
	"time"
)

type Genero struct {
	ID        uint   `gorm:"primaryKey"`
	Descricao string `gorm:"not null;default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (g *Genero) UpdateFromInput(s *schemas.GeneroInput) {
	g.Descricao = s.Descricao
}
