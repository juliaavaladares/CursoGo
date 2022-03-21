package servidor

import (
	"CursoGo/crud-basico/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario usuario

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler corpo da requisição"))
		return
	}

	err = json.Unmarshal(corpoRequisicao, &usuario)
	if err != nil {
		w.Write([]byte("Falha ao converter usuario para struct"))
		return
	}

	db, err := config.ConectaBancoDeDados()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("insert into usuario (nome, email) values (?, ?)")
	if err != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer statement.Close()

	insercao, err := statement.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		w.Write([]byte("Erro ao executar statement"))
		return
	}

	idInserido, err := insercao.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao obter id inserido"))
		return
	}

	mensagem := fmt.Sprintf("Usuario %d inserido com sucesso", idInserido)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(mensagem))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter parametro para inteiro"))
		return
	}

	db, err := config.ConectaBancoDeDados()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	linha, err := db.Query("select * from usuario where id = ?", ID)
	if err != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer linha.Close()

	var usuario usuario
	if linha.Next() {
		err = linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Nome)
		if err != nil {
			w.Write([]byte("Erro ao escanear usuario"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(usuario)
	if err != nil {
		w.Write([]byte("Erro ao converter usuarios para json"))
		return
	}
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	var usuarios []usuario

	db, err := config.ConectaBancoDeDados()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	linhas, err := db.Query("select * from usuario")
	if err != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer linhas.Close()

	for linhas.Next() {
		var usuario usuario
		err = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nome)
		if err != nil {
			w.Write([]byte("Erro ao escanear usuario"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(usuarios)
	if err != nil {
		w.Write([]byte("Erro ao converter usuarios para json"))
		return
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter parametro para inteiro"))
		return
	}

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler corpo da requisição"))
		return
	}
	var usuario usuario
	err = json.Unmarshal(corpoRequisicao, &usuario)
	if err != nil {
		w.Write([]byte("Falha ao converter usuario para struct"))
		return
	}

	db, err := config.ConectaBancoDeDados()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("update usuario set nome = ?, email = ? where id = ?")
	if err != nil {
		fmt.Println("Err: ", err)
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer statement.Close()

	_, err = statement.Exec(usuario.Nome, usuario.Email, ID)
	if err != nil {
		w.Write([]byte("Erro ao executar statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter parametro para inteiro"))
		return
	}

	db, err := config.ConectaBancoDeDados()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("delete from  usuario where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer statement.Close()

	_, err = statement.Exec(ID)
	if err != nil {
		w.Write([]byte("Erro ao executar statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
