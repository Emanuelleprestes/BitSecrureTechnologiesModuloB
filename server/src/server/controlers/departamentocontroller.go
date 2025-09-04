package controlers

import (
	"database/sql"
	"fmt"

	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/models/colaborador"
	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/models/departamento"
	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/models/projeto"
)

type Departamentocontroller struct {
	conn *sql.DB
}

func NewDepartamentocontroller(conn *sql.DB) *Departamentocontroller {
	return &Departamentocontroller{conn: conn}
}

// NewDepartamento cria um novo Departamento/Estrutura de equipe vazio

func (d *Departamentocontroller) Newdepartamento(
	g *colaborador.Colaborador,
	m []*colaborador.Colaborador,
	p *projeto.Projeto,
) (*departamento.Departamento, error) {
	departamento := departamento.Departamento{Gestor: g, Membros: m, Projeto: p}
	err := d.validacao()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &departamento, nil
}

func (d *Departamentocontroller) validacao() error {
	return nil
}
