package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// a declaração do objeto
type Server struct {
	ip   string
	conn *sql.Conn
}

func helloword(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Println(w.Write([]byte("helloword")))
}

// rota /colaboradores
// acho que vou mapear as rotas e depois so voltar o mux mesmo para ficar mais organizado
// função que vai configurar as rotas, middleware e as handle functions para a mesma
func (s *Server) routesforcolabolador(r *chi.Mux) {
	r.Route("/colaboradores/", func(r chi.Router) {
		r.Get("/", Getcolaboladores)
		r.Get("/{name}", GetcolaboladoresByName)
		r.Post("/", Save)
	})
}

// rotas /cargo
// função que vai configurar as rotass, middleware e as handle functions para a mesma
func (s *Server) routerforcargos(_ *chi.Mux) {
	fmt.Println("nada")
}

// rotas /derpartamento
// função que vai configurar as rotass, middleware e as handle functions para a mesma
func (s *Server) routerforderpartamento(_ *chi.Mux) {
	fmt.Println("nada")
}

// rotas /documento
// função que vai configurar as rotass, middleware e as handle functions para a mesma
func (s *Server) routerfordocumento(_ *chi.Mux) {
	fmt.Println("nada")
}

// rotas /tarefa
// função que vai configurar as rotass, middleware e as handle functions para a mesma
func (s *Server) routerfortarefa(_ *chi.Mux) {
	fmt.Println("nada")
}

// rotas /ponto
// função que vai configurar as rotass, middleware e as handle functions para a mesma
func (s *Server) routerforponto(_ *chi.Mux) {
	fmt.Println("nada")
}

// rotas /backup
// função que vai configurar as rotass, middleware e as handle functions para a mesma
func (s *Server) routerforbackup(_ *chi.Mux) {
	fmt.Println("nada")
}

// rotas /autentificação
// função que vai configurar as rotas, middleware e as handle functions para a mesma
func (s *Server) routerforautentificacao(_ *chi.Mux) {
	fmt.Println("nada")
}

// função que chama as funções para configurar as routas para o *chi.Mux
// de modo essa função aqui não ficar muito grande
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

// função que inicia um objeto e volta a referencia dele
func Newserver(ip string, conn *sql.Conn) *Server {
	return &Server{
		ip,
		conn,
	}
}

// função que inicia o ouvinte para a porta (8080,80 e etc) para as coneçoes http/https
// atualmente so configurado para a coneção http
func (s *Server) Run() error {
	r := s.config()
	fmt.Printf("ip: http://127.0.0.1%s/\n", s.ip)
	return http.ListenAndServe(s.ip, r)
}
