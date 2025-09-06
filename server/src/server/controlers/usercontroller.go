package controlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/models/colaborador"
	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/repositorios"
)

type Colaboradorcontroller struct {
	conn *sql.DB
}

func NewColaboradorcontroller(c *sql.DB) *Colaboradorcontroller {
	return &Colaboradorcontroller{conn: c}
}

// 0 seria o null para inteiro, e nil seria o null para referencia
func (c *Colaboradorcontroller) Newuser(
	cpf, nome, cargo, setor, status, email, ramal, habilidades string,
) (*colaborador.Colaborador, error) {
	colaborador := colaborador.Colaborador{
		Nome:        nome,
		CPF:         cpf,
		Setor:       setor,
		Status:      status,
		Email:       email,
		Ramal:       ramal,
		Habilidades: habilidades,
	}
	err := c.validacao(&colaborador)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &colaborador, nil
}

func iscpf(cpf string) bool {
	match, _ := regexp.MatchString(`^\d{3}\.\d{3}\.\d{3}-\d{2}$`, cpf)
	if !match {
		return false
	}

	// Extrai apenas os dígitos (11 no total)
	var d [11]int
	idx := 0
	for _, r := range cpf {
		if r >= '0' && r <= '9' {
			if idx >= 11 {
				return false
			}
			d[idx] = int(r - '0')
			idx++
		}
	}
	if idx != 11 {
		return false
	}

	// Rejeita sequências com todos os dígitos iguais
	allEqual := true
	for i := 1; i < 11; i++ {
		if d[i] != d[0] {
			allEqual = false
			break
		}
	}
	if allEqual {
		return false
	}

	// Calcula 1º dígito verificador
	sum := 0
	for i := range 9 {
		sum += d[i] * (10 - i)
	}
	dv1 := 11 - (sum % 11)
	if dv1 >= 10 {
		dv1 = 0
	}

	// Calcula 2º dígito verificador
	sum = 0
	for i := range 10 {
		sum += d[i] * (11 - i)
	}
	dv2 := 11 - (sum % 11)
	if dv2 >= 10 {
		dv2 = 0
	}

	return dv1 == d[9] && dv2 == d[10]
}

func (c *Colaboradorcontroller) Getall() (*[]colaborador.Colaborador, error) {
	userrepo := repositorios.NewColaboradorRepo(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return userrepo.Getall(ctx)
}

// por vai so isso para validação das tabelas
func (c *Colaboradorcontroller) Create(colab *colaborador.Colaborador) error {
	userrepo := repositorios.NewColaboradorRepo(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	passwd := colab.Senha
	hash, err := bcrypt.GenerateFromPassword(passwd, bcrypt.DefaultCost)
	colab.Senha = hash
	_, err = userrepo.Save(ctx, colab)
	if err != nil {
		return err
	}
	return nil
}

func (c *Colaboradorcontroller) Login(email, pass string) bool {
	return false
}

func (c *Colaboradorcontroller) Update(colab *colaborador.Colaborador) error {
	userrepo := repositorios.NewColaboradorRepo(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := userrepo.Update(ctx, *colab)
	if err != nil {
		return err
	}
	return nil
}

func (c *Colaboradorcontroller) validacao(colab *colaborador.Colaborador) error {
	if iscpf(colab.CPF) {
		return nil
	}
	return errors.New("cpf invalido")
}
