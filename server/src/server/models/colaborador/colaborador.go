package colaborador

import "github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/models/cargo"

type Colaborador struct {
	id              int    `json:""`
	nome            string `json:""`
	cpf             string `json:""`
	email_coportivo string `json:""`
	telefone        string `json:""`
	cargo           *cargo.Cargo
}

func Newcolaborador() *Colaborador {
	return &Colaborador{}
}
