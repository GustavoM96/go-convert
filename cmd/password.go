package cmd

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/spf13/cobra"
)

var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Gera uma senha aleatória com letras e números",
	Long:  `Gera uma senha aleatória com letras maiúsculas, minúsculas e números. Permite definir o tamanho da senha e a quantidade de senhas geradas.`,
	Run: func(cmd *cobra.Command, args []string) {
		length, _ := cmd.Flags().GetInt("length")
		times, _ := cmd.Flags().GetInt("times")

		if length <= 0 {
			fmt.Println("O tamanho da senha deve ser maior que zero.")
			return
		}
		if times <= 0 {
			times = 1
		}
		for i := 0; i < times; i++ {
			fmt.Println(generatePassword(length))
		}
	},
}

func generatePassword(length int) string {
	const charset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	for range length {
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		sb.WriteByte(charset[idx.Int64()])
	}
	return sb.String()
}

func init() {
	passwordCmd.Flags().IntP("length", "l", 12, "Tamanho da senha gerada")
	passwordCmd.Flags().IntP("times", "t", 1, "Quantidade de senhas a serem geradas")
	rootCmd.AddCommand(passwordCmd)
}
