package models

import (
	"time"
)

type Livro struct {
	ID        uint   `gorm:"primaryKey"`
	Titulo    string `gorm:"not null;default:null"`
	EditoraID int
	Editora   Editora
	Autores   []Autor  `gorm:"many2many:livro_autores;"`
	Generos   []Genero `gorm:"many2many:livro_generos;"`
	Edicao    int
	Ano       int
	Paginas   int
	ISBN      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
