package colaborador

import (
	"database/sql"

	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/repositorios"
)

type Colaboradorrepo struct{}

// Get implements repositorios.Crudinterface.
func (c Colaboradorrepo) Get(y) T {
	panic("unimplemented")
}

// Save implements repositorios.Crudinterface.
func (c Colaboradorrepo) Save() T {
	panic("unimplemented")
}

// Update implements repositorios.Crudinterface.
func (c Colaboradorrepo) Update(T) T {
	panic("unimplemented")
}

func Newcolarepo(conn *sql.Conn) (error, Colaboradorrepo) {
	var colaborador repositorios.Crudinterface[any, any] = Colaboradorrepo{}
}
