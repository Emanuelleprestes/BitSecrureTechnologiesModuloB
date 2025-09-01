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

// acho que vou mapear as rotas e depois so voltar o mux mesmo para ficar mais organizado

func (s *Server) routesforcolabolador(r *chi.Mux) {
	r.Route("/colaboradores/", func(r chi.Router) {
		r.Get("/", Getcolaboladores)
		r.Get("/{name}", GetcolaboladoresByName)
		r.Post("/", Save)
	})
}
func (s *Server) routerforcargos(r *chi.Mux)         {}
func (s *Server) routerforderpartamento(r *chi.Mux)  {}
func (s *Server) routerfordocumento(r *chi.Mux)      {}
func (s *Server) routerfortarefa(r *chi.Mux)         {}
func (s *Server) routerforponto(r *chi.Mux)          {}
func (s *Server) routerforbackup(r *chi.Mux)         {}
func (s *Server) routerforautentificacao(r *chi.Mux) {}

func (s *Server) config() *chi.Mux {
	r := chi.NewMux()
	r.Get("/", helloword)
	s.routesforcolabolador(r)
	s.routerforautentificacao(r)
	s.routerforbackup(r)
	s.routerforcargos(r)
	s.routerforderpartamento(r)
	s.routerfordocumento(r)
	s.routerforponto(r)
	s.routerfortarefa(r)
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
