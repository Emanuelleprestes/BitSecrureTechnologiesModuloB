package colaborador

// Colaborador representa a tabela colaborador no banco de dados
type Colaborador struct {
	ID          int    `json:"id_colaborador"`
	CPF         string `json:"cpf,omitempty"`
	Nome        string `json:"nome"`
	Senha       string `json:"senha"`
	Cargo       string `json:"cargo,omitempty"`
	Setor       string `json:"setor,omitempty"`
	Status      string `json:"status,omitempty"`
	Email       string `json:"email,omitempty"`
	Ramal       string `json:"ramal,omitempty"`
	Habilidades string `json:"habilidades,omitempty"`
}

// NewColaborador cria um novo Colaborador vazio
func NewColaborador() *Colaborador {
	return &Colaborador{}
}
