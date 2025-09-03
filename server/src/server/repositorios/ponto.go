package repositorios

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/models/ponto"
)

type pont = ponto.Ponto

// PontoRepo implementa CRUD básico
type PontoRepo struct {
	db    *sql.DB
	ponto *pont
}

// NewPontoRepo cria um novo repositório de Ponto
func NewPontoRepo(db *sql.DB) (*PontoRepo, error) {
	repo := &PontoRepo{db: db}

	// Compile-time check: garante que PontoRepo implementa a interface
	// var _ repositorios.Crudinterface[pont, int] = repo

	return repo, nil
}

// Get retorna um registro de ponto pelo ID
func (r *PontoRepo) Get(id int) pont {
	p := pont{}
	query := "SELECT id_ponto, colaborador_id, entrada, saida FROM registroponto WHERE id_ponto=?"
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.ColaboradorID, &p.Entrada, &p.Saida)
	if err != nil {
		fmt.Println("Erro ao buscar ponto:", err)
	}
	return p
}

// GetAll retorna todos os registros de ponto
func (r *PontoRepo) GetAll() ([]pont, error) {
	query := "SELECT id_ponto, colaborador_id, entrada, saida FROM registroponto"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pontos []pont
	for rows.Next() {
		var p pont
		err := rows.Scan(&p.ID, &p.ColaboradorID, &p.Entrada, &p.Saida)
		if err != nil {
			return nil, err
		}
		pontos = append(pontos, p)
	}
	return pontos, nil
}

// Save insere um novo registro de ponto
func (r *PontoRepo) Save() pont {
	query := "INSERT INTO registroponto (colaborador_id, entrada, saida) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, r.ponto.ColaboradorID, r.ponto.Entrada, r.ponto.Saida)
	if err != nil {
		fmt.Println("Erro ao salvar ponto:", err)
		return pont{}
	}
	id, _ := result.LastInsertId()
	r.ponto.ID = int(id)
	return *r.ponto
}

// Update atualiza um registro de ponto existente
func (r *PontoRepo) Update(p pont) pont {
	query := "UPDATE registroponto SET colaborador_id=?, entrada=?, saida=? WHERE id_ponto=?"
	_, err := r.db.Exec(query, p.ColaboradorID, p.Entrada, p.Saida, p.ID)
	if err != nil {
		fmt.Println("Erro ao atualizar ponto:", err)
	}
	return p
}

// Delete remove um registro de ponto pelo ID
func (r *PontoRepo) Delete(id int) error {
	query := "DELETE FROM registroponto WHERE id_ponto=?"
	_, err := r.db.Exec(query, id)
	return err
}

// GetByColaborador retorna todos os registros de ponto de um colaborador
func (r *PontoRepo) GetByColaborador(colaboradorID int) ([]pont, error) {
	query := "SELECT id_ponto, colaborador_id, entrada, saida FROM registroponto WHERE colaborador_id=?"
	rows, err := r.db.Query(query, colaboradorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pontos []pont
	for rows.Next() {
		var p pont
		err := rows.Scan(&p.ID, &p.ColaboradorID, &p.Entrada, &p.Saida)
		if err != nil {
			return nil, err
		}
		pontos = append(pontos, p)
	}
	return pontos, nil
}
