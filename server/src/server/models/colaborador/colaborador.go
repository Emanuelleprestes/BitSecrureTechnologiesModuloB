package colaborador

import "github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/models/cargo"

type Colaborador struct {
	id              int
	nome            string
	cpf             string
	email_coportivo string
	telefone        string
	cargo           *cargo.Cargo
}

func Newcolaborador() *Colaborador {
	return &Colaborador{}
}
