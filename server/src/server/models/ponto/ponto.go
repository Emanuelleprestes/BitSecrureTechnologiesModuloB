package ponto

import (
	"time"
)

// Ponto representa o registro de ponto de um colaborador
type Ponto struct {
	ID            int       `json:"id_ponto"`          // id do registro de ponto
	ColaboradorID int       `json:"colaborador_id"`    // id do colaborador
	Entrada       time.Time `json:"entrada,omitempty"` // hora de entrada
	Saida         time.Time `json:"saida,omitempty"`   // hora de saída
}

// NewPonto cria um novo registro de ponto vazio
func NewPonto() *Ponto {
	return &Ponto{}
}

