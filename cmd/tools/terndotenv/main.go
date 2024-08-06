package main

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	// if err := cmd.Run(); err != nil {
	// 	panic(err)
	// }

	// TODO: Colocar isso numa função??
	// Buffers para capturar a saída padrão e a saída de erro
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// Executar o comando
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao executar o comando:", err)
		fmt.Println("Saída de erro:", stderr.String())
		return
	}

	// Imprimir a saída do comando se não houver erro
	fmt.Println("Saída do comando:", out.String())
}
