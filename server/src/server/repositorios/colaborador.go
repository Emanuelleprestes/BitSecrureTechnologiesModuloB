package repositorios

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/models/colaborador"
)

type Colab = colaborador.Colaborador

type ColaboradorRepo struct {
	db *sql.DB
}

func NewColaboradorRepo(db *sql.DB) *ColaboradorRepo {
	return &ColaboradorRepo{db: db}
}

// Get retorna um colaborador pelo ID
func (r *ColaboradorRepo) Get(ctx context.Context, id int) (Colab, error) {
	c := Colab{}
	query := `
		SELECT id_colaborador, cpf, nome, cargo, setor, status, email, ramal, habilidades
		FROM colaborador
		WHERE id = ?
	`
	err := r.db.QueryRowContext(ctx, query, id).
		Scan(&c.ID, &c.CPF, &c.Nome, &c.Cargo, &c.Setor, &c.Status, &c.Email, &c.Ramal, &c.Habilidades)
	if err != nil {
		if err == sql.ErrNoRows {
			return c, nil
		}
		return c, fmt.Errorf("erro ao buscar colaborador: %w", err)
	}
	return c, nil
}

func (r *ColaboradorRepo) Getbyemail(ctx context.Context, email string) (Colab, error) {
	c := Colab{}
	query := `
		SELECT id_colaborador, cpf, nome, cargo, setor, status, email, ramal, habilidades, senha
		FROM colaborador
		WHERE email = ?
	`
	err := r.db.QueryRowContext(ctx, query, email).
		Scan(&c.ID, &c.CPF, &c.Nome, &c.Cargo, &c.Setor, &c.Status, &c.Email, &c.Ramal, &c.Habilidades, &c.Senha)
	if err != nil {
		if err == sql.ErrNoRows {
			return c, nil
		}
		return c, fmt.Errorf("erro ao buscar colaborador: %w", err)
	}
	return c, nil
}

func (r *ColaboradorRepo) Getbyname(ctx context.Context, name string) (Colab, error) {
	c := Colab{}
	query := `
		SELECT id_colaborador, cpf, nome, cargo, setor, status, email, ramal, habilidades, senha
		FROM colaborador
		WHERE nome = ?
	`
	err := r.db.QueryRowContext(ctx, query, name).
		Scan(&c.ID, &c.CPF, &c.Nome, &c.Cargo, &c.Setor, &c.Status, &c.Email, &c.Ramal, &c.Habilidades, &c.Senha)
	if err != nil {
		if err == sql.ErrNoRows {
			return c, nil
		}
		return c, fmt.Errorf("erro ao buscar colaborador: %w", err)
	}
	return c, nil
}

func (r *ColaboradorRepo) Getall(ctx context.Context) ([]Colab, error) {
	query := `
        SELECT id_colaborador, cpf, nome, cargo, setor, status, email, ramal, habilidades
        FROM colaborador
    `
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar colaboradores: %w", err)
	}
	defer rows.Close() // importante: fecha o cursor
	var colaboradores []Colab
	for rows.Next() {
		var c Colab
		err := rows.Scan(
			&c.ID, &c.CPF, &c.Nome, &c.Cargo, &c.Setor,
			&c.Status, &c.Email, &c.Ramal, &c.Habilidades,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear colaborador: %w", err)
		}
		colaboradores = append(colaboradores, c)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao iterar colaboradores: %w", err)
	}

	return colaboradores, nil
}

// Save insere um novo colaborador
func (r *ColaboradorRepo) Save(ctx context.Context, c *colaborador.Colaborador) (Colab, error) {
	query := `
		INSERT INTO colaborador (cpf, nome, cargo, setor, status, email, ramal, habilidades,senha)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?,?)
	`
	result, err := r.db.ExecContext(ctx, query,
		c.CPF, c.Nome, c.Cargo, c.Setor,
		c.Status, c.Email, c.Ramal, c.Habilidades, c.Senha,
	)
	if err != nil {
		return Colab{}, fmt.Errorf("erro ao salvacaborador: %w", err)
	}
	id, _ := result.LastInsertId()
	c.ID = int(id)
	return *c, nil
}

// Update atualiza um colaborador
func (r *ColaboradorRepo) Update(ctx context.Context, c Colab) (Colab, error) {
	query := `
		UPDATE colaborador
		SET nome=?, cargo=?, setor=?, status=?, email=?, ramal=?, habilidades=?
		WHERE id_colaborador=?
	`
	_, err := r.db.ExecContext(ctx, query,
		c.Nome, c.Cargo, c.Setor, c.Status,
		c.Email, c.Ramal, c.Habilidades, c.ID,
	)
	if err != nil {
		return c, fmt.Errorf("erro ao atualizacaborador: %w", err)
	}
	return c, nil
}
