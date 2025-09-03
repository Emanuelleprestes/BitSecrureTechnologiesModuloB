package tarefa

import (
	"time"

	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/models/colaborador"
)

// Tarefa representa uma tarefa vinculada a um projeto
type Tarefa struct {
	ID          int                      `json:"id_tarefa"`             // id da tarefa
	Titulo      string                   `json:"titulo"`                // título da tarefa
	ProjetoID   int                      `json:"projeto_id,omitempty"`  // id do projeto
	Prazo       time.Time                `json:"prazo,omitempty"`       // prazo da tarefa
	Prioridade  string                   `json:"prioridade,omitempty"`  // "Baixa","Média","Alta","Urgente"
	Status      string                   `json:"status,omitempty"`      // "To-Do","Doing","Done", etc.
	Responsavel *colaborador.Colaborador `json:"responsavel,omitempty"` // colaborador responsável
}

// NewTarefa cria uma nova Tarefa vazia
func NewTarefa() *Tarefa {
	return &Tarefa{}
}

