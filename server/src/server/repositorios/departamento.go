package repositorios

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/models/departamento"
)

type depto = departamento.Departamento

type DepartamentoRepo struct {
	db              *sql.DB
	colaboradorRepo *ColaboradorRepo
}

// NewDepartamentoRepo cria um novo repositório de Departamento
func NewDepartamentoRepo(db *sql.DB, colaboradorRepo *ColaboradorRepo) *DepartamentoRepo {
	return &DepartamentoRepo{
		db:              db,
		colaboradorRepo: colaboradorRepo,
	}
}

// Get retorna um departamento pelo ID do projeto
func (r *DepartamentoRepo) Get(ctx context.Context, projetoID int) (depto, error) {
	d := depto{
		ID: projetoID,
	}

	// Busca os colaboradores do projeto
	query := "SELECT id_colaborador FROM projetoequipe WHERE id_projeto=?"
	rows, err := r.db.QueryContext(ctx, query, projetoID)
	if err != nil {
		return d, fmt.Errorf("erro ao buscar membros: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var colaboradorID int
		if err := rows.Scan(&colaboradorID); err != nil {
			return d, fmt.Errorf("erro ao ler id_colaborador: %w", err)
		}

		colab, err := r.colaboradorRepo.Get(ctx, colaboradorID)
		if err != nil {
			return d, fmt.Errorf("erro ao buscar colaborador: %w", err)
		}

		d.Membros = append(d.Membros, &colab)
	}

	// Opcional: buscar o gestor (se houver uma regra ou campo específico)
	// Exemplo: primeiro colaborador da lista é gestor
	if len(d.Membros) > 0 {
		d.Gestor = d.Membros[0]
	}

	return d, nil
}

// GetAll retorna todos os departamentos (todos os projetos com suas equipes)
func (r *DepartamentoRepo) GetAll(ctx context.Context) ([]depto, error) {
	// Busca todos os projetos que possuem equipe
	query := "SELECT DISTINCT id_projeto FROM projetoequipe"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar projetos: %w", err)
	}
	defer rows.Close()

	var departamentos []depto
	for rows.Next() {
		var projetoID int
		if err := rows.Scan(&projetoID); err != nil {
			return nil, fmt.Errorf("erro ao ler id_projeto: %w", err)
		}

		d, err := r.Get(ctx, projetoID)
		if err != nil {
			return nil, err
		}

		departamentos = append(departamentos, d)
	}

	return departamentos, nil
}

// AddMember adiciona um colaborador a um projeto/departamento
func (r *DepartamentoRepo) AddMember(ctx context.Context, projetoID int, colaboradorID int) error {
	query := "INSERT INTO projetoequipe (id_projeto, id_colaborador) VALUES (?, ?)"
	_, err := r.db.ExecContext(ctx, query, projetoID, colaboradorID)
	return err
}

// RemoveMember remove um colaborador de um projeto/departamento
func (r *DepartamentoRepo) RemoveMember(
	ctx context.Context,
	projetoID int,
	colaboradorID int,
) error {
	query := "DELETE FROM projetoequipe WHERE id_projeto=? AND id_colaborador=?"
	_, err := r.db.ExecContext(ctx, query, projetoID, colaboradorID)
	return err
}
