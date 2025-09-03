package cargo

// Cargo representa um cargo de um colaborador
type Cargo struct {
	ID    int    `json:"id_cargo"`
	Nome  string `json:"nome"`            // nome do cargo, ex: "Analista de Sistemas"
	Setor string `json:"setor,omitempty"` // setor do cargo, se aplicável
	Nivel string `json:"nivel,omitempty"` // nível do cargo, ex: "Junior", "Pleno", "Senior"
}

// NewCargo cria um novo Cargo vazio
func NewCargo() *Cargo {
	return &Cargo{}
}
