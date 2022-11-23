package models

import (
	"time"
)

type Livro struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey"`
	Titulo    string    `json:",omitempty" gorm:"not null;default:null"`
	EditoraID int       `json:",omitempty"`
	Editora   Editora   `json:",omitempty"`
	Autores   []Autor   `json:",omitempty" gorm:"many2many:livro_autores;"`
	Generos   []Genero  `json:",omitempty" gorm:"many2many:livro_generos;"`
	Edicao    int       `json:",omitempty"`
	Ano       int       `json:",omitempty"`
	Paginas   int       `json:",omitempty"`
	ISBN      string    `json:",omitempty"`
	CreatedAt time.Time `json:",omitempty"`
	UpdatedAt time.Time `json:",omitempty"`
}
