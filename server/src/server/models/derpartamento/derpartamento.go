package derpartamento

import "github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/models/colaborador"

type Derpartamento struct {
	// não preciso colocar o nome correto por casua da
	// da json alguma que me permite colocar o nome da tabela
	id     int
	nome   string
	gestor *colaborador.Colaborador
}

func NewDerpartamento() *Derpartamento {
	return &Derpartamento{}
}
