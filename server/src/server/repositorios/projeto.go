package repositorios

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/models/projeto"
)

type proj = projeto.Projeto

// ProjetoRepo implementa CRUD básico
type ProjetoRepo struct {
	db      *sql.DB
	projeto *proj
}

// NewProjetoRepo cria um novo repositório de Projeto
func NewProjetoRepo(db *sql.DB) (*ProjetoRepo, error) {
	repo := &ProjetoRepo{db: db}

	// Compile-time check: garante que ProjetoRepo implementa a interface
	// var _ repositorios.Crudinterface[proj, int] = repo

	return repo, nil
}

// Get retorna um projeto pelo ID
func (r *ProjetoRepo) Get(id int) (*proj, error) {
	p := proj{}
	query := "SELECT id_projeto, nome, tipo, status, progresso, responsavel, descricao FROM projeto WHERE id_projeto=?"
	err := r.db.QueryRow(query, id).
		Scan(&p.ID, &p.Nome, &p.Tipo, &p.Status, &p.Progresso, &p.Responsavel, &p.Descricao)
	if err != nil {
		fmt.Println("Erro ao buscar projeto:", err)
		return nil, err
	}
	return &p, nil
}

// GetAll retorna todos os projetos
func (r *ProjetoRepo) GetAll() ([]proj, error) {
	query := "SELECT id_projeto, nome, tipo, status, progresso, responsavel, descricao FROM projeto"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projetos []proj
	for rows.Next() {
		var p proj
		err := rows.Scan(
			&p.ID,
			&p.Nome,
			&p.Tipo,
			&p.Status,
			&p.Progresso,
			&p.Responsavel,
			&p.Descricao,
		)
		if err != nil {
			return nil, err
		}
		projetos = append(projetos, p)
	}
	return projetos, nil
}

// Save insere um novo projeto
func (r *ProjetoRepo) Save() proj {
	query := "INSERT INTO projeto (nome, tipo, status, progresso, responsavel, descricao) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := r.db.Exec(
		query,
		r.projeto.Nome,
		r.projeto.Tipo,
		r.projeto.Status,
		r.projeto.Progresso,
		r.projeto.Responsavel,
		r.projeto.Descricao,
	)
	if err != nil {
		fmt.Println("Erro ao salvar projeto:", err)
		return proj{}
	}
	id, _ := result.LastInsertId()
	r.projeto.ID = int(id)
	return *r.projeto
}

// Update atualiza um projeto existente
func (r *ProjetoRepo) Update(p proj) proj {
	query := "UPDATE projeto SET nome=?, tipo=?, status=?, progresso=?, responsavel=?, descricao=? WHERE id_projeto=?"
	_, err := r.db.Exec(
		query,
		p.Nome,
		p.Tipo,
		p.Status,
		p.Progresso,
		p.Responsavel,
		p.Descricao,
		p.ID,
	)
	if err != nil {
		fmt.Println("Erro ao atualizar projeto:", err)
	}
	return p
}

// Delete remove um projeto pelo ID
func (r *ProjetoRepo) Delete(id int) error {
	query := "DELETE FROM projeto WHERE id_projeto=?"
	_, err := r.db.Exec(query, id)
	return err
}

// GetByResponsavel retorna todos os projetos de um responsável específico
func (r *ProjetoRepo) GetByResponsavel(responsavelID int) ([]proj, error) {
	query := "SELECT id_projeto, nome, tipo, status, progresso, responsavel, descricao FROM projeto WHERE responsavel=?"
	rows, err := r.db.Query(query, responsavelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projetos []proj
	for rows.Next() {
		var p proj
		err := rows.Scan(
			&p.ID,
			&p.Nome,
			&p.Tipo,
			&p.Status,
			&p.Progresso,
			&p.Responsavel,
			&p.Descricao,
		)
		if err != nil {
			return nil, err
		}
		projetos = append(projetos, p)
	}
	return projetos, nil
}
