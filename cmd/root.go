// go-converter/cmd/root.go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd representa o comando base sem subcomandos
var rootCmd = &cobra.Command{
	Use:   "go-convert",
	Short: "Uma ferramenta CLI para conversões rápidas.",
	Long: `go-convert é uma aplicação de linha de comando para realizar
conversões bidirecionais, como Base64 <-> Texto e decodificação de JWT.`,
}

// Execute adiciona todos os comandos filhos ao comando raiz e define as flags apropriadamente.
// É chamado por main.main(). Só precisa acontecer uma vez para o rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Ocorreu um erro: '%s'", err)
		os.Exit(1)
	}
}
