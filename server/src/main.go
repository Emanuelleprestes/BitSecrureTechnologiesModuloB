package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server" // importando a estrutura servidor pelo imports
)

// essa função vai ser responsavel por configuras algumas coisas basica
// tipo conexão com o banco de dados que vai ser passado como referencia para uma estrutura
// essa função roda antes da função main
func init() {
	fmt.Println("teste")
}

func ConnectMariaDB() (*sql.DB, error) {
	dsn := "tiago:123@tcp(127.0.0.1:3306)/sys?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão: %w", err)
	}

	// Testa conexão
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("erro ao pingar banco: %w", err)
	}

	// (Opcional) configurar pool
	db.SetMaxOpenConns(10) // máximo de conexões abertas
	db.SetMaxIdleConns(5)  // máximo de conexões inativas
	db.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func main() {
	// instaciando a estrtura do servidor
	// essa estrutura precisa do ip|porta e uma conexão com um banco de dados
	// ele esta aceitando nil porque eu recebo uma referencia para a conexão com banco de dados
	// por causa disso em toda a interação com o bancos de dados e melhor verificar se nulo ou não
	db, err := ConnectMariaDB()
	if err != nil {
		fmt.Println(err)
	}
	server := server.Newserver(":8080", db)
	// logo mais isso vai virar uma gorutines, por agora não precisa
	// para poder parar o servidor via input e etc
	if err := server.Run(); err != nil {
		fmt.Println(err)
	}
}
