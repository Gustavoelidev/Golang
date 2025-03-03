package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local" // conectar ao banco."user:senha@/nomedobanco"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexao esta aberta!")

	linhas, erro := db.Query("select * from usuario")

	if erro != nil {
		log.Fatal(erro)
	}

	defer linhas.Close()

	println(linhas)

}
