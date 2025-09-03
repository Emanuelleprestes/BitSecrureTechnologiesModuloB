package departamento

import "github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/models/colaborador"

// Departamento representa um departamento ou equipe de colaboradores
type Departamento struct {
	ID      int                        `json:"id"`                // id do departamento/equipe
	Nome    string                     `json:"nome"`              // nome do departamento/equipe
	Gestor  *colaborador.Colaborador   `json:"gestor,omitempty"`  // gestor do departamento/equipe
	Membros []*colaborador.Colaborador `json:"membros,omitempty"` // lista de colaboradores da equipe
}

// NewDepartamento cria um novo Departamento/Estrutura de equipe vazio
func NewDepartamento() *Departamento {
	return &Departamento{}
}
