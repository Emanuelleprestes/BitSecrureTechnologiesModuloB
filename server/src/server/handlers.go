package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/controlers"
	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/models/colaborador"
)

type UserSessao struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Cargo string `json:"cargo"`
}

// aqui vai ficar as funçoes que vao ir para para as rotas
type (
	writer   = http.ResponseWriter
	resquest = *http.Request
)

type colaHandler struct {
	conn *sql.DB
}

func (h *colaHandler) Me(w writer, r resquest) {
	// pegando o colaborador da sessao
	user, ok := sessionManager.Get(r.Context(), "user").(UserSessao)

	// neste se, a gente esta verificando o !ok
	// o !ok que dizer se a condição for false entre no bloco abaixo
	if !ok {
		http.Error(
			w,
			fmt.Errorf("erro ao confimar o usuario: %v", ok).Error(),
			http.StatusInternalServerError,
		)
	}

	// criando uma instancia do controle colaborador e passando como argumento a conexão com o banco de dados
	userecon := controlers.NewColaboradorcontroller(h.conn)

	// obtendo o usuario usando o metodo do controller para isso
	// ja que é uma operação envovendo o banco de dados, pode voltar duas variaveis
	// o usuario/colaborador ou um erro
	userdb, err := userecon.Getbyemail(user.Email)
	//
	// esse bloco de se, verifica se a condição do erro não é nulo, no caso vazio
	// se o erro for diferente de vazio, ele entra no bloco
	if err != nil {
		http.Error(
			w,
			fmt.Errorf("não atorizado a entrar: %w", err).Error(),
			http.StatusUnauthorized,
		)
		return
	}

	// zerando alguns campos importantes
	userdb.ID = 0
	userdb.Senha = ""

	// colancando no header o tipo de resposta como json
	w.Header().Set("Content-Type", "application/json")

	// enviando os dados para o cliente http via json
	// eu estou criando um novo encoder passando o http.writer para ele escrever direto na resposta,
	// o econder diz para transformar esses dados em binario
	// o map[string]any significa que ele vai criar um map com chaves para os valores, as chaves e represetado pelo [string], string significa um valor aonde
	// pega caracteres dentro do ""
	// o any significa qualquer coisa
	// no codigo abaixo o any e para o "message": por ele ser um objeto
	json.NewEncoder(w).Encode(map[string]any{
		"status": "200",
		"message": map[string]string{
			"cpf":        userdb.CPF,
			"nome":       userdb.Nome,
			"cargo":      userdb.Cargo,
			"setor":      userdb.Setor,
			"status":     userdb.Status,
			"email":      userdb.Email,
			"ramal":      userdb.Ramal,
			"habilidade": userdb.Habilidades,
		},
	})
}

func (h *colaHandler) Login(w writer, r resquest) {
	// verificando o tipo de conteudo
	// o conteudo vindo do site precisa ser application/json
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type deve ser application/json", http.StatusUnsupportedMediaType)
		return
	}

	// criando uma estrutura basica para pegar os dados basicos do login
	// o value pode ser tanto nome tanto email, preferivel email por ser unico
	var creds struct {
		Value string `json:"value"`
		Senha string `json:"senha"`
	}

	// decodificando o json, isso vai pegar o json e passar os valores do json para a estrutura creds
	err := json.NewDecoder(r.Body).Decode(&creds)
	//
	// verificação de erro basica
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// criando uma nova instancia do controler e passando a conexão com o banco de dados
	usercontroller := controlers.NewColaboradorcontroller(h.conn)

	// criando uma variavel para o colaborador, no caso seria um endereço na memoria para ele
	var user *colaborador.Colaborador

	// chamando uma função para pegar os dados originais do colaborador
	// reatibuindo os valores user e err, o user vai pegar o novo endereço que vier da função
	// ja o err vai pegar o novo erro
	user, err = usercontroller.Login(creds.Value, creds.Senha)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	//
	// criando uma instancia da estrutura usersessao para salvar o dados do login na memoria do servidor
	usersessao := UserSessao{
		ID:    user.ID,
		Nome:  user.Nome,
		Email: user.Email,
		Cargo: user.Cargo,
	}

	// colancando a estrutura na memoria do servidor
	sessionManager.Put(r.Context(), "user", usersessao)

	// setando o header para enviar json de volta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"message": "Login efetuado com sucesso",
	})
}

func (h *colaHandler) Getcolaboladores(w writer, r resquest) {
	colabcontroller := controlers.NewColaboradorcontroller(h.conn)
	colaboradores, err := colabcontroller.Getall()
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// loop responsavel por zerar o id
	for i := range colaboradores {
		colaboradores[i].ID = 0
	}
	data, err := json.Marshal(colaboradores)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(data)
}

func (h *colaHandler) Createcolaborador(w writer, r resquest) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type deve ser application/json", http.StatusUnsupportedMediaType)
		return
	}
	var colaborador colaborador.Colaborador
	controlercolaborador := controlers.NewColaboradorcontroller(h.conn)
	err := json.NewDecoder(r.Body).Decode(&colaborador)
	if err != nil {
		http.Error(
			w,
			fmt.Errorf("o json esta errado: %w", err).Error(),
			http.StatusInternalServerError,
		)
		return
	}
	err = controlercolaborador.Create(&colaborador)
	if err != nil {
		http.Error(
			w,
			fmt.Errorf("erro ao criar o usuario: %w", err).Error(),
			http.StatusInternalServerError,
		)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"Status":  "200",
		"message": "colaborador criado",
	})
}

func (h *colaHandler) Deletecolaborador(w writer, r resquest) {
	fmt.Println(w.Write([]byte("")))
}

func (h *colaHandler) GetcolaboladoresByName(w writer, r resquest) {
	fmt.Println(w.Write([]byte("")))
}

func (h *colaHandler) Updatecolaborador(w writer, r resquest) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type deve ser application/json", http.StatusUnsupportedMediaType)
		return
	}
	var colaborador colaborador.Colaborador
	controlercolaborador := controlers.NewColaboradorcontroller(h.conn)
	err := json.NewDecoder(r.Body).Decode(&colaborador)
	if err != nil {
		http.Error(
			w,
			fmt.Errorf("o json esta errado: %w", err).Error(),
			http.StatusInternalServerError,
		)
		return
	}
	val := sessionManager.Get(r.Context(), "user")
	user, ok := val.(UserSessao)
	if !ok {
		http.Error(
			w,
			fmt.Errorf("erro ao confimar o usuario: %v", ok).Error(),
			http.StatusInternalServerError,
		)
	}
	if user.Nome != colaborador.Nome && user.Cargo != "gestor" {
		http.Error(
			w,
			fmt.Errorf("o usuario não pode modificar outros").Error(),
			http.StatusUnauthorized,
		)
	}
	err = controlercolaborador.Update(&colaborador)
	if err != nil {
		http.Error(
			w,
			fmt.Errorf("erro ao modifcar o usuario: %w", err).Error(),
			http.StatusInternalServerError,
		)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"Status":  "200",
		"message": "colaborador atualizar",
	})
}
