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
	db  *sql.DB
	col *Colab
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
		WHERE id_colaborador = ?
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

// Save insere um novo colaborador
func (r *ColaboradorRepo) Save(ctx context.Context) (Colab, error) {
	query := `
		INSERT INTO colaborador (cpf, nome, cargo, setor, status, email, ramal, habilidades)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := r.db.ExecContext(ctx, query,
		r.col.CPF, r.col.Nome, r.col.Cargo, r.col.Setor,
		r.col.Status, r.col.Email, r.col.Ramal, r.col.Habilidades,
	)
	if err != nil {
		return Colab{}, fmt.Errorf("erro ao salvar colaborador: %w", err)
	}
	id, _ := result.LastInsertId()
	r.col.ID = int(id)
	return *r.col, nil
}

// Update atualiza um colaborador
func (r *ColaboradorRepo) Update(ctx context.Context, c Colab) (Colab, error) {
	query := `
		UPDATE colaborador
		SET cpf=?, nome=?, cargo=?, setor=?, status=?, email=?, ramal=?, habilidades=?
		WHERE id_colaborador=?
	`
	_, err := r.db.ExecContext(ctx, query,
		c.CPF, c.Nome, c.Cargo, c.Setor, c.Status,
		c.Email, c.Ramal, c.Habilidades, c.ID,
	)
	if err != nil {
		return c, fmt.Errorf("erro ao atualizar colaborador: %w", err)
	}
	return c, nil
}

