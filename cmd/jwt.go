package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var jwtCmd = &cobra.Command{
	Use:   "jwt",
	Short: "Decodifica o payload de um token JWT",
	Long:  `Use a flag --decode (-d) para fornecer o token JWT e exibir seu payload.`,
	Run: func(cmd *cobra.Command, args []string) {
		tokenString, _ := cmd.Flags().GetString("decode")

		if tokenString == "" {
			cmd.Help()
			return
		}

		parts := strings.Split(tokenString, ".")
		if len(parts) != 3 {
			fmt.Fprintf(os.Stderr, "Erro: Formato de token JWT inválido. Esperado 3 partes separadas por '.'.\n")
			os.Exit(1)
		}

		payload, err := base64.RawURLEncoding.DecodeString(parts[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao decodificar o payload do JWT: %v\n", err)
			os.Exit(1)
		}

		var prettyPayload map[string]interface{}
		if err := json.Unmarshal(payload, &prettyPayload); err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao formatar o payload JSON: %v\n", err)
			fmt.Println(string(payload))
			os.Exit(1)
		}

		prettyJSON, err := json.MarshalIndent(prettyPayload, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao formatar o payload JSON: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(prettyJSON))
	},
}

func init() {
	jwtCmd.Flags().StringP("decode", "d", "", "Token JWT para ter o payload decodificado")
	rootCmd.AddCommand(jwtCmd)
}
