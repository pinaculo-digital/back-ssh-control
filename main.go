package main

import (
	"fmt"
	"go_service/core/server"
	"go_service/core/server/shared"
	"os/exec"
)

// @title						Escritorio PINATALK
// @version					@0.0.1
// @description				Documentação do melhor escritorio do mundo, feito em GO
// @BasePath					/
// @SecurityDefinitions.apikey	Bearer Auth
// @in							header
// @name						Authorization
func main() {

	exec.Command("swag", "init")
	database, err := server.InitConnection()
	if err != nil {
		fmt.Println("Erro na conexão do banco: ", err.Error())
		return
	}
	shared.SetDB(database)
	defer database.Close()

	server.MainServer()

}
