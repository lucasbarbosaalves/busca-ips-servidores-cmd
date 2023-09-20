package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ips públicos do host")

	applicacao := app.Gerar()

	// Executa a aplicação, OS.Args é o argumento que vem da linha de comando
	// Metodo Run retorna um erro
	// Tratando erro com if init
	if erro := applicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	} 
}