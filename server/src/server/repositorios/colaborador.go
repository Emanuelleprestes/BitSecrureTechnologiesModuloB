package repositorios

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/models/colaborador"
)

// Alias para facilitar
type Colab = colaborador.Colaborador

// ColaboradorRepo implementa o CRUD para Colaborador
type ColaboradorRepo struct {
	conn *sql.DB
}

// NewColaboradorRepo cria um novo repositório de Colaborador
func NewColaboradorRepo(conn *sql.DB) *ColaboradorRepo {
	return &ColaboradorRepo{conn: conn}
}

// Get retorna um colaborador pelo ID
func (r *ColaboradorRepo) Get(ctx context.Context, id int) (Colab, error) {
	var col Colab
	query := `
	SELECT id_colaborador, nome, cargo, setor, status, email, ramal, habilidades
	FROM colaborador
	WHERE id_colaborador = ?
	`
	err := r.conn.QueryRowContext(ctx, query, id).Scan(
		&col.ID, &col.Nome, &col.Cargo, &col.Setor,
		&col.Status, &col.Email, &col.Ramal, &col.Habilidades,
	)
	if err != nil {
		return Colab{}, fmt.Errorf("erro ao buscar colaborador: %w", err)
	}
	return col, nil
}

// GetAll retorna todos os colaboradores
func (r *ColaboradorRepo) GetAll(ctx context.Context) ([]Colab, error) {
	query := `
	SELECT id_colaborador, nome, cargo, setor, status, email, ramal, habilidades
	FROM colaborador
	`
	rows, err := r.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var colaboradores []Colab
	for rows.Next() {
		var col Colab
		if err := rows.Scan(
			&col.ID, &col.Nome, &col.Cargo, &col.Setor,
			&col.Status, &col.Email, &col.Ramal, &col.Habilidades,
		); err != nil {
			return nil, err
		}
		colaboradores = append(colaboradores, col)
	}

	return colaboradores, rows.Err()
}

// Save insere um novo colaborador
func (r *ColaboradorRepo) Save(ctx context.Context, col Colab) (Colab, error) {
	query := `
	INSERT INTO colaborador (nome, cargo, setor, status, email, ramal, habilidades)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	result, err := r.conn.ExecContext(ctx, query,
		col.Nome, col.Cargo, col.Setor, col.Status, col.Email, col.Ramal, col.Habilidades,
	)
	if err != nil {
		return Colab{}, fmt.Errorf("erro ao salvar colaborador: %w", err)
	}

	id, _ := result.LastInsertId()
	col.ID = int(id)
	return col, nil
}

// Update atualiza um colaborador existente
func (r *ColaboradorRepo) Update(ctx context.Context, col Colab) (Colab, error) {
	query := `
	UPDATE colaborador SET nome=?, cargo=?, setor=?, status=?, email=?, ramal=?, habilidades=?
	WHERE id_colaborador=?
	`
	_, err := r.conn.ExecContext(ctx, query,
		col.Nome, col.Cargo, col.Setor, col.Status, col.Email, col.Ramal, col.Habilidades, col.ID,
	)
	if err != nil {
		return col, fmt.Errorf("erro ao atualizar colaborador: %w", err)
	}
	return col, nil
}

// Delete remove um colaborador pelo ID
func (r *ColaboradorRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM colaborador WHERE id_colaborador=?`
	_, err := r.conn.ExecContext(ctx, query, id)
	return err
}

// GetByStatus retorna colaboradores com determinado status
func (r *ColaboradorRepo) GetByStatus(ctx context.Context, status string) ([]Colab, error) {
	query := `
	SELECT id_colaborador, nome, cargo, setor, status, email, ramal, habilidades
	FROM colaborador
	WHERE status = ?
	`
	rows, err := r.conn.QueryContext(ctx, query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var colaboradores []Colab
	for rows.Next() {
		var col Colab
		if err := rows.Scan(
			&col.ID, &col.Nome, &col.Cargo, &col.Setor,
			&col.Status, &col.Email, &col.Ramal, &col.Habilidades,
		); err != nil {
			return nil, err
		}
		colaboradores = append(colaboradores, col)
	}

	return colaboradores, rows.Err()
}

// GetByProjeto retorna colaboradores de um projeto específico
func (r *ColaboradorRepo) GetByProjeto(ctx context.Context, projetoID int) ([]Colab, error) {
	query := `
	SELECT c.id_colaborador, c.nome, c.cargo, c.setor, c.status, c.email, c.ramal, c.habilidades
	FROM colaborador c
	JOIN projetoequipe pe ON c.id_colaborador = pe.id_colaborador
	WHERE pe.id_projeto = ?
	`
	rows, err := r.conn.QueryContext(ctx, query, projetoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var colaboradores []Colab
	for rows.Next() {
		var col Colab
		if err := rows.Scan(
			&col.ID, &col.Nome, &col.Cargo, &col.Setor,
			&col.Status, &col.Email, &col.Ramal, &col.Habilidades,
		); err != nil {
			return nil, err
		}
		colaboradores = append(colaboradores, col)
	}

	return colaboradores, rows.Err()
}

// GetByTarefa retorna o colaborador responsável por uma tarefa
func (r *ColaboradorRepo) GetByTarefa(ctx context.Context, tarefaID int) (*Colab, error) {
	query := `
	SELECT c.id_colaborador, c.nome, c.cargo, c.setor, c.status, c.email, c.ramal, c.habilidades
	FROM colaborador c
	JOIN tarefa t ON c.id_colaborador = t.responsavel
	WHERE t.id_tarefa = ?
	`
	var col Colab
	err := r.conn.QueryRowContext(ctx, query, tarefaID).Scan(
		&col.ID, &col.Nome, &col.Cargo, &col.Setor,
		&col.Status, &col.Email, &col.Ramal, &col.Habilidades,
	)
	if err != nil {
		return nil, err
	}
	return &col, nil
}

