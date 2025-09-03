package repositorios

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/models/cargo"
)

type carg = cargo.Cargo

// CargoRepo implementa Crudinterface[Cargo, int]
type CargoRepo struct {
	db    *sql.DB
	cargo *carg
}

// NewCargoRepo cria um novo repositório de Cargo
func NewCargoRepo(db *sql.DB) (*CargoRepo, error) {
	repo := &CargoRepo{db: db}

	// Compile-time check: garante que CargoRepo implementa a interface
	// var _ repositorios.Crudinterface[Cargo, int] = repo

	return repo, nil
}

// Get retorna um cargo pelo ID
func (r *CargoRepo) Get(id int) carg {
	c := carg{}
	query := "SELECT id_cargo, nome, setor, nivel FROM cargo WHERE id_cargo = ?"
	err := r.db.QueryRow(query, id).Scan(&c.ID, &c.Nome, &c.Setor, &c.Nivel)
	if err != nil {
		fmt.Println("Erro ao buscar cargo:", err)
	}
	return c
}

// GetAll retorna todos os cargos
func (r *CargoRepo) GetAll() ([]carg, error) {
	query := "SELECT id_cargo, nome, setor, nivel FROM cargo"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cargos []carg
	for rows.Next() {
		var c carg
		err := rows.Scan(&c.ID, &c.Nome, &c.Setor, &c.Nivel)
		if err != nil {
			return nil, err
		}
		cargos = append(cargos, c)
	}
	return cargos, nil
}

// Save insere um novo cargo
func (r *CargoRepo) Save() carg {
	query := "INSERT INTO cargo (nome, setor, nivel) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, r.cargo.Nome, r.cargo.Setor, r.cargo.Nivel)
	if err != nil {
		fmt.Println("Erro ao salvar cargo:", err)
		return carg{}
	}
	id, _ := result.LastInsertId()
	r.cargo.ID = int(id)
	return *r.cargo
}

// Update atualiza um cargo existente
func (r *CargoRepo) Update(c carg) carg {
	query := "UPDATE cargo SET nome=?, setor=?, nivel=? WHERE id_cargo=?"
	_, err := r.db.Exec(query, c.Nome, c.Setor, c.Nivel, c.ID)
	if err != nil {
		fmt.Println("Erro ao atualizar cargo:", err)
	}
	return c
}

// Delete remove um cargo pelo ID
func (r *CargoRepo) Delete(id int) error {
	query := "DELETE FROM cargo WHERE id_cargo=?"
	_, err := r.db.Exec(query, id)
	return err
}

// GetBySetor retorna todos os cargos de um setor específico
func (r *CargoRepo) GetBySetor(setor string) ([]carg, error) {
	query := "SELECT id_cargo, nome, setor, nivel FROM cargo WHERE setor=?"
	rows, err := r.db.Query(query, setor)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cargos []carg
	for rows.Next() {
		var c carg
		err := rows.Scan(&c.ID, &c.Nome, &c.Setor, &c.Nivel)
		if err != nil {
			return nil, err
		}
		cargos = append(cargos, c)
	}
	return cargos, nil
}

