package controlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strings"
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
	// aqui eu estou iniciando um novo repositorio passado como referencia a a conexão com o db
	userrepo := repositorios.NewColaboradorRepo(c.conn)
	// criando um contexto para as operações dos repositorios que aqui vai ser utilizado
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// aqui eu estou dizendo para concelar depois que o contexto passar de 4 segundos
	defer cancel()
	// obtendo a senha
	passwd := colab.Senha
	// criando um array de bytes criptografados, funciona assim
	// eu pego a string que o passwd e converto elaem byte pelo []byte(string)
	// assim o GenerateFromPassword pega o array de byte recem criado so para isso e
	// me devolve um erro ou um array de byte de novo, assim eu posso fazer o string(hash) e volta
	// a senha como string
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	// assim em muitos lugares aqui neste if eu posso lidar com o erro como eu quero
	// seria melhor criar uma estrutura de log e etc para isso
	// mais isso so deixaria mais complexo depois para os meus colegas
	// então não planejo fazer isso
	if err != nil {
		return err
	}
	// aqui assim em outras linguagens eu volto a senha modificada para string e mudo a string original servida apenas para gerar a senha nova
	// e poder criar novas instacias direto do json sem problemas com o hash, que aqui vai ser usado
	colab.Senha = string(hash)
	fmt.Println(colab)
	// aqui agora mando para o metodo do repositorio do colaborador para criar o usuario
	_, err = userrepo.Save(ctx, colab)
	if err != nil {
		return err
	}
	return nil
}

func (c *Colaboradorcontroller) Login(creds, pass string) (*colaborador.Colaborador, error) {
	if strings.Contains(creds, "@") {
		return c.loginbyemail(creds, pass)
	} else {
		return c.loginbynome(creds, pass)
	}
}

func (c *Colaboradorcontroller) loginbynome(nome, pass string) (*colaborador.Colaborador, error) {
	userrepo := repositorios.NewColaboradorRepo(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	user, err := userrepo.Getbyname(ctx, nome)
	if err != nil {
		return nil, fmt.Errorf("usuario não existe: %w", err)
	}
	// senha correta: primeiro hash do banco, depois senha digitada
	err = bcrypt.CompareHashAndPassword([]byte(user.Senha), []byte(pass))
	if err != nil {
		return nil, errors.New("senha invalida")
	}

	return &user, nil
}

func (c *Colaboradorcontroller) loginbyemail(email, pass string) (*colaborador.Colaborador, error) {
	userrepo := repositorios.NewColaboradorRepo(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	user, err := userrepo.Getbyemail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("usuario não existe: %w", err)
	}
	// senha correta: primeiro hash do banco, depois senha digitada
	err = bcrypt.CompareHashAndPassword([]byte(user.Senha), []byte(pass))
	if err != nil {
		return nil, errors.New("senha invalida")
	}

	return &user, nil
}

func (c *Colaboradorcontroller) Update(colab *colaborador.Colaborador) error {
	userrepo := repositorios.NewColaboradorRepo(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()
	user, err := userrepo.Getbyemail(ctx, colab.Email)
	if err != nil {
		return err
	}
	_, err = userrepo.Update(ctx, user)
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
