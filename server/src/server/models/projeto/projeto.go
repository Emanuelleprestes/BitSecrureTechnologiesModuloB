package projeto

import (
	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/models/colaborador"
)

// Projeto representa um projeto da empresa
type Projeto struct {
	ID          int                      `json:"id_projeto"`            // id do projeto
	Nome        string                   `json:"nome"`                  // nome do projeto
	Tipo        string                   `json:"tipo,omitempty"`        // "Software" ou "Security"
	Status      string                   `json:"status,omitempty"`      // "Em Andamento","Atrasado","No Prazo"
	Progresso   int                      `json:"progresso,omitempty"`   // progresso em %
	Responsavel *colaborador.Colaborador `json:"responsavel,omitempty"` // colaborador responsável
	Descricao   string                   `json:"descricao,omitempty"`   // descrição do projeto
}

// NewProjeto cria um novo Projeto vazio
func NewProjeto() *Projeto {
	return &Projeto{}
}
