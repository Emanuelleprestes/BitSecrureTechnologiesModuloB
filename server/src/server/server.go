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
	w.WriteHeader(200)
	fmt.Println(w.Write([]byte("helloword")))
}

func (s *Server) config() *chi.Mux {
	r := chi.NewMux()
	r.Get("/", helloword)
	return r
}

func Newserver(ip string, conn *sql.Conn) *Server {
	return &Server{
		ip,
		conn,
	}
}

func (s *Server) Run() error {
	r := s.config()
	fmt.Printf("ip: http://127.0.0.1%s/\n", s.ip)
	return http.ListenAndServe(s.ip, r)
}
