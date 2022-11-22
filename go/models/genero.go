package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Genero struct {
	ID        uint   `gorm:"primaryKey"`
	Descricao string `gorm:"not null;default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (g *Genero) Update(c *gin.Context, db *gorm.DB, n *Genero) {
	g.Descricao = n.Descricao
}
