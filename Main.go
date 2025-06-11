package main

import (
	"database/sql"  // Para interagir com bancos de dados
//	"encoding/json" // Para lidar com JSON
	"fmt"           // Para impressão de mensagens
	"log"           // Para logs de erro
	"net/http"      // Para criar o servidor web
//	"strconv"       // Para converter strings em números

//	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3" // Importa o driver SQLite
)

// A estrutura Product será definida aqui mais tarde
type Product struct {
	ID int `json:"id"`
	nome string `json:"nome"`
	price float64 `json:"preco"`
	estoque int `json:"preco"`

}

var db *sql.DB // Variável para a conexão com o banco de dados

func main() {
	fmt.Println("Configurando o servidor...")

	var err error
	
	db, err = sql.Open("sqlite3", "./product.db")

	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	createTableSQL := `
		CREATE TABLE IF NOT EXISTS produtos(
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			NOME TEXT NOT NULL,
			PRECO FLOAT NOT NULL,
			ESTOQUE INTEGER NOT NULL
		);`

	_, err = db.Exec(createTableSQL)
	if err != nil{
		log.Fatal(err)
	}
	
	fmt.Println("Tabela 'products' verificada/criada com sucesso.")


	fmt.Println("Servidor iniciado na porta :8080 (aguardando rotas)")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Funções de CRUD virão aqui (handlers)