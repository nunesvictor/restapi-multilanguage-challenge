package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Genero struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey"`
	Descricao string    `json:",omitempty" gorm:"not null;default:null"`
	CreatedAt time.Time `json:",omitempty"`
	UpdatedAt time.Time `json:",omitempty"`
}

func (g *Genero) Update(c *gin.Context, db *gorm.DB, n *Genero) {
	g.Descricao = n.Descricao

	Update(c, db, g, nil)
}
