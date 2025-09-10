package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-convert",
	Short: "Uma ferramenta CLI para conversões rápidas.",
	Long:  `go-convert é uma aplicação de linha de comando para realizar conversões bidirecionais, como Base64 <-> Texto e decodificação de JWT e geração de senhas.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Ocorreu um erro: '%s'", err)
		os.Exit(1)
	}
}
