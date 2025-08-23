package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	ip   string
	conn *sql.Conn
}

func helloword(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w.Write([]byte("helloword")))
}

func Newserver(ip string, conn *sql.Conn) Server {
	return Server{
		ip,
		conn,
	}
}

func (s *Server) Run() error {
	r := chi.NewRouter()
	r.Get("/", helloword)
	http.HandleFunc("/", helloword)
	return http.ListenAndServe(s.ip, r)
}
