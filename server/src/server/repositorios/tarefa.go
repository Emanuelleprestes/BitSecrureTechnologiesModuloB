package repositorios

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/models/tarefa"
)

type tarr = tarefa.Tarefa

type TarefaRepo struct {
	db              *sql.DB
	tarefa          *tarr
	colaboradorRepo *ColaboradorRepo
}

// NewTarefaRepo cria um novo repositório de Tarefa
func NewTarefaRepo(db *sql.DB, colaboradorRepo *ColaboradorRepo) (*TarefaRepo, error) {
	return &TarefaRepo{
		db:              db,
		colaboradorRepo: colaboradorRepo,
	}, nil
}

// Get retorna uma tarefa pelo ID
func (r *TarefaRepo) Get(ctx context.Context, id int) (tarr, error) {
	t := tarr{}
	var responsavelID sql.NullInt64

	query := "SELECT id_tarefa, titulo, id_projeto, prazo, prioridade, status, responsavel FROM tarefa WHERE id_tarefa=?"
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&t.ID,
		&t.Titulo,
		&t.ProjetoID,
		&t.Prazo,
		&t.Prioridade,
		&t.Status,
		&responsavelID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return t, nil
		}
		return t, fmt.Errorf("erro ao buscar tarefa: %w", err)
	}

	if responsavelID.Valid {
		colaborador, err := r.colaboradorRepo.Get(ctx, int(responsavelID.Int64))
		if err != nil {
			return t, err
		}
		t.Responsavel = &colaborador
	}

	return t, nil
}

// GetAll retorna todas as tarefas
func (r *TarefaRepo) GetAll(ctx context.Context) ([]tarr, error) {
	query := "SELECT id_tarefa, titulo, id_projeto, prazo, prioridade, status, responsavel FROM tarefa"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tarefas []tarr
	for rows.Next() {
		var t tarr
		var responsavelID sql.NullInt64
		err := rows.Scan(
			&t.ID,
			&t.Titulo,
			&t.ProjetoID,
			&t.Prazo,
			&t.Prioridade,
			&t.Status,
			&responsavelID,
		)
		if err != nil {
			return nil, err
		}
		if responsavelID.Valid {
			colaborador, err := r.colaboradorRepo.Get(ctx, int(responsavelID.Int64))
			if err != nil {
				return nil, err
			}
			t.Responsavel = &colaborador
		}
		tarefas = append(tarefas, t)
	}

	return tarefas, nil
}

// Save insere uma nova tarefa
func (r *TarefaRepo) Save(ctx context.Context) (tarr, error) {
	responsavelID := 0
	if r.tarefa.Responsavel != nil {
		responsavelID = r.tarefa.Responsavel.ID
	}
	query := "INSERT INTO tarefa (titulo, id_projeto, prazo, prioridade, status, responsavel) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := r.db.ExecContext(
		ctx,
		query,
		r.tarefa.Titulo,
		r.tarefa.ProjetoID,
		r.tarefa.Prazo,
		r.tarefa.Prioridade,
		r.tarefa.Status,
		responsavelID,
	)
	if err != nil {
		return tarr{}, fmt.Errorf("erro ao salvar tarefa: %w", err)
	}
	id, _ := result.LastInsertId()
	r.tarefa.ID = int(id)
	return *r.tarefa, nil
}

// Update atualiza uma tarefa existente
func (r *TarefaRepo) Update(ctx context.Context, t tarr) (tarr, error) {
	responsavelID := 0
	if t.Responsavel != nil {
		responsavelID = t.Responsavel.ID
	}
	query := "UPDATE tarefa SET titulo=?, id_projeto=?, prazo=?, prioridade=?, status=?, responsavel=? WHERE id_tarefa=?"
	_, err := r.db.ExecContext(
		ctx,
		query,
		t.Titulo,
		t.ProjetoID,
		t.Prazo,
		t.Prioridade,
		t.Status,
		responsavelID,
		t.ID,
	)
	if err != nil {
		return t, fmt.Errorf("erro ao atualizar tarefa: %w", err)
	}
	return t, nil
}

// Delete remove uma tarefa pelo ID
func (r *TarefaRepo) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM tarefa WHERE id_tarefa=?"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// GetByProjeto retorna todas as tarefas de um projeto específico
func (r *TarefaRepo) GetByProjeto(ctx context.Context, projetoID int) ([]tarr, error) {
	query := "SELECT id_tarefa, titulo, id_projeto, prazo, prioridade, status, responsavel FROM tarefa WHERE id_projeto=?"
	rows, err := r.db.QueryContext(ctx, query, projetoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tarefas []tarr
	for rows.Next() {
		var t tarr
		var responsavelID sql.NullInt64
		err := rows.Scan(
			&t.ID,
			&t.Titulo,
			&t.ProjetoID,
			&t.Prazo,
			&t.Prioridade,
			&t.Status,
			&responsavelID,
		)
		if err != nil {
			return nil, err
		}
		if responsavelID.Valid {
			colaborador, err := r.colaboradorRepo.Get(ctx, int(responsavelID.Int64))
			if err != nil {
				return nil, err
			}
			t.Responsavel = &colaborador
		}
		tarefas = append(tarefas, t)
	}

	return tarefas, nil
}

// GetByResponsavel retorna todas as tarefas de um colaborador específico
func (r *TarefaRepo) GetByResponsavel(ctx context.Context, responsavelID int) ([]tarr, error) {
	query := "SELECT id_tarefa, titulo, id_projeto, prazo, prioridade, status, responsavel FROM tarefa WHERE responsavel=?"
	rows, err := r.db.QueryContext(ctx, query, responsavelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tarefas []tarr
	for rows.Next() {
		var t tarr
		var rID sql.NullInt64
		err := rows.Scan(&t.ID, &t.Titulo, &t.ProjetoID, &t.Prazo, &t.Prioridade, &t.Status, &rID)
		if err != nil {
			return nil, err
		}
		if rID.Valid {
			colaborador, err := r.colaboradorRepo.Get(ctx, int(rID.Int64))
			if err != nil {
				return nil, err
			}
			t.Responsavel = &colaborador
		}
		tarefas = append(tarefas, t)
	}

	return tarefas, nil
}

