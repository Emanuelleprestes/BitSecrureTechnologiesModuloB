package departamento

import (
	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/models/colaborador"
	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/models/projeto"
)

// Departamento representa um departamento ou equipe de colaboradores
type Departamento struct {
	Gestor  *colaborador.Colaborador   `json:"gestor,omitempty"`  // gestor do departamento/equipe
	Membros []*colaborador.Colaborador `json:"membros,omitempty"` // lista de colaboradores da equipe
	Projeto *projeto.Projeto           `json:"projeto"`
}

// NewDepartamento cria um novo Departamento/Estrutura de equipe vazio
func NewDepartamento() *Departamento {
	return &Departamento{}
}
