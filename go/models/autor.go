package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Autor struct {
	ID        uint   `gorm:"primaryKey"`
	Nome      string `gorm:"not null;default:null"`
	Sobrenome string `gorm:"not null;default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Autor) TableName() string {
	return "autores"
}

func (a *Autor) Update(c *gin.Context, db *gorm.DB, n *Autor) {
	a.Nome = n.Nome
	a.Sobrenome = n.Sobrenome

	Update(c, db, a, nil)
}
