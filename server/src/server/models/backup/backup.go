package backup

import "time"

// Backup representa a tabela backup no banco de dados
type Backup struct {
	ID          int       `json:"id_backup"`
	Data        time.Time `json:"data"`                  // timestamp do backup
	MantidoAte  time.Time `json:"mantido_ate,omitempty"` // data até quando será mantido
	Responsavel int       `json:"responsavel"`           // id do colaborador responsável
}

// NewBackup cria um novo Backup vazio
func NewBackup() *Backup {
	return &Backup{}
}

