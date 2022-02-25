package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	stringConexao := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True", user, password, db_name)

	db, err := sql.Open("mysql", stringConexao)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão aberta!")

	linhas, err := db.Query("select * from usuario")
	if err != nil {
		log.Fatal(err)
	}

	defer linhas.Close()

	fmt.Println(linhas)

}
