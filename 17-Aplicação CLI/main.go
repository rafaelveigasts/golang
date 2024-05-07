package main

import (
	"linhaDeComando/app"
	"log"
	"os"
)

func main(){
	aplicacao := app.Gerar()
	// aplicacao.Run([]string{"linhaDeComando"})
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
