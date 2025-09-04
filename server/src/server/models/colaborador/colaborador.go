package colaborador

// Colaborador representa a tabela colaborador no banco de dados
type Colaborador struct {
	ID          int    `json:"id_colaborador,omitempty"`
	CPF         string `json:"cpf"`
	Nome        string `json:"nome"`
	Senha       string `json:"senha"`
	Cargo       string `json:"cargo"`
	Setor       string `json:"setor"`
	Status      string `json:"status"`
	Email       string `json:"email"`
	Ramal       string `json:"ramal"`
	Habilidades string `json:"habilidades"`
}

// NewColaborador cria um novo Colaborador vazio
func NewColaborador() *Colaborador {
	return &Colaborador{}
}
