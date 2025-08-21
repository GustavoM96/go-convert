// go-converter/cmd/jwt.go
package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// jwtCmd representa o comando 'jwt' que agora usa uma flag
var jwtCmd = &cobra.Command{
	Use:   "jwt",
	Short: "Decodifica o payload de um token JWT",
	Long:  `Use a flag --decode (-d) para fornecer o token JWT e exibir seu payload.`,
	// A função Run é executada quando o comando 'jwt' é chamado.
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Pega o valor fornecido para a flag "decode".
		tokenString, _ := cmd.Flags().GetString("decode")

		// 2. Se nenhum valor foi passado para a flag, o tokenString estará vazio.
		//    Nesse caso, mostramos a ajuda do comando e saímos.
		if tokenString == "" {
			cmd.Help()
			return
		}

		// 3. O resto da lógica é a mesma de antes: decodificar e exibir o payload.
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

// A função init é onde a "mágica" acontece.
func init() {
	// Adicionamos o comando 'jwt' ao comando raiz.
	rootCmd.AddCommand(jwtCmd)

	// Aqui, definimos a flag que o comando 'jwt' vai aceitar.
	// StringP cria uma flag que aceita um valor de texto (string).
	// O 'P' no final significa que também definimos uma versão curta (Shorthand).
	//
	// Parâmetros:
	// 1. "decode": Nome longo da flag (--decode)
	// 2. "d":      Nome curto da flag (-d)  <-- É ISSO QUE PERMITE O SEU EXEMPLO!
	// 3. "":       Valor padrão (se a flag não for usada)
	// 4. "Token JWT para ter o payload decodificado": Texto de ajuda da flag.
	jwtCmd.Flags().StringP("decode", "d", "", "Token JWT para ter o payload decodificado")
}
