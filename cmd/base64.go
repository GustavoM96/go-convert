package cmd

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "Codifica ou decodifica strings para Base64 usando flags",
	Long:  `Use as flags --encode (-e) ou --decode (-d) para realizar a conversão. As duas flags não podem ser usadas ao mesmo tempo.`,
	Run: func(cmd *cobra.Command, args []string) {

		encodeStr, _ := cmd.Flags().GetString("encode")
		decodeStr, _ := cmd.Flags().GetString("decode")

		if encodeStr != "" && decodeStr != "" {
			fmt.Fprintln(os.Stderr, "Erro: As flags --encode e --decode são mutuamente exclusivas.")
			os.Exit(1)
		}

		if encodeStr != "" {
			encoded := base64.StdEncoding.EncodeToString([]byte(encodeStr))
			fmt.Println(encoded)
			return
		}

		if decodeStr != "" {
			decoded, err := base64.StdEncoding.DecodeString(decodeStr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erro: A entrada fornecida para --decode não é uma string Base64 válida.\n")
				os.Exit(1)
			}
			fmt.Println(string(decoded))
			return
		}

		cmd.Help()
	},
}

func init() {
	base64Cmd.Flags().StringP("encode", "e", "", "Texto a ser codificado para Base64")
	base64Cmd.Flags().StringP("decode", "d", "", "String Base64 a ser decodificada para texto")
	rootCmd.AddCommand(base64Cmd)
}
