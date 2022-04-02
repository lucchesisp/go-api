package main

import (
	"fmt"
	"github.com/lucchesisp/go-api/database"
	"github.com/lucchesisp/go-api/routes"
)

func main() {

	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor Rest com GO")
	routes.HandleRequest()
}
