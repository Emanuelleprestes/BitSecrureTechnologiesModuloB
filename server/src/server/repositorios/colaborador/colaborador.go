package colaborador

import (
	"database/sql"

	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/repositorios"
)

type Colaboradorrepo[T any, y any] struct{}

// Get implements repositorios.Crudinterface.
func (c Colaboradorrepo[T, y]) Get(y) T {
	panic("unimplemented")
}

// Save implements repositorios.Crudinterface.
func (c Colaboradorrepo[T, y]) Save() T {
	panic("unimplemented")
}

// Update implements repositorios.Crudinterface.
func (c Colaboradorrepo[T, y]) Update(T) T {
	panic("unimplemented")
}

func NewColaRepo(conn *sql.Conn) (Colaboradorrepo[string, string], error) {
	repo := Colaboradorrepo[string, string]{}
	// compile-time check: garante que Colaboradorrepo implementa a interface
	var _ repositorios.Crudinterface[string, string] = repo

	return repo, nil
}
