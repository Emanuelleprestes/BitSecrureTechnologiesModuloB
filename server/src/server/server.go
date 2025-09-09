package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/alexedwards/scs/v2/memstore"
	"github.com/go-chi/chi/v5"
)

// a declaração do objeto
type Server struct {
	ip     string
	conn   *sql.DB
	handle *Handlers
}

func funchandcle(
	f http.HandlerFunc,
	privilegio []string,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val := sessionManager.Get(r.Context(), "user")
		if val == nil {
			http.Error(w, "sem permissão para acessar essa página", http.StatusUnauthorized)
			return
		}

		userusersessao, ok := val.(UserSessao)
		if !ok {
			http.Error(w, "sessão inválida", http.StatusUnauthorized)
			return
		}
		if slices.Contains(privilegio, userusersessao.Cargo) {
			f(w, r)
			return
		} else {
			http.Error(w, "acesso negado", http.StatusForbidden)
			return
		}
	}
}

// rota /colaboradores
// acho que vou mapear as rotas e depois so voltar o mux mesmo para ficar mais organizado
// função que vai configurar as rotas, middleware e as handle functions para a mesma
func (s *Server) routesforcolabolador(r *chi.Mux) {
	r.Route("/colaboradores", func(r chi.Router) {
		r.Get("/", funchandcle(s.handle.Getcolaboladores, []string{"gestor", "adm"}))
		r.Get("/{name}", s.handle.GetcolaboladoresByName)
		r.Post("/", funchandcle(s.handle.Createcolaborador, []string{"gestor", "adm"}))
		r.Put("/", s.handle.Updatecolaborador)
		r.Delete("/{name}", funchandcle(s.handle.Deletecolaborador, []string{"gestor", "adm"}))
	})
}

// json de array ou json so a estrutura em si
// status: 200
// rotas /cargo
// função que vai configurar as rotass, middleware e as handle functions para a mesma
func (s *Server) routerforcargos(r *chi.Mux) {
	r.Route("/cargo", func(r chi.Router) {
		r.Get("/", nil)
		r.Get("/{nome}", nil)
		r.Post("/", nil)
	})
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
	r.Post("/login", s.handle.Login)
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
func Newserver(ip string, conn *sql.DB) *Server {
	Handlers := Handlers{conn: conn}
	return &Server{
		ip,
		conn,
		&Handlers,
	}
}

var sessionManager *scs.SessionManager

// função que inicia o ouvinte para a porta (8080,80 e etc) para as coneçoes http/https
// atualmente so configurado para a coneção http
func (s *Server) Run() error {
	r := s.config()
	sessionManager = scs.New()
	sessionManager.Store = memstore.New()
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Lifetime = 2 * time.Hour
	fmt.Printf("ip: http://127.0.0.1%s/\n", s.ip)
	return http.ListenAndServe(s.ip, sessionManager.LoadAndSave(r))
}
