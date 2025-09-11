package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var textCmd = &cobra.Command{
	Use:   "text",
	Short: "Realiza conversão de palavras",
	Long:  `Realiza conversão de palavras, como lower, upper, adição de caracter etc...`,
	Run: func(cmd *cobra.Command, args []string) {
		text, _ := cmd.Flags().GetString("text")
		result := text

		if len(text) <= 0 {
			fmt.Println("O tamanho do texto deve ser maior que zero.")
			return
		}

		if cmd.Flags().Changed("lower") {
			position, _ := cmd.Flags().GetInt("lower")
			result = transformByPosition(result, position, strings.ToLower)
		}

		if cmd.Flags().Changed("upper") {
			position, _ := cmd.Flags().GetInt("upper")
			result = transformByPosition(result, position, strings.ToUpper)
		}

		if start, _ := cmd.Flags().GetString("add-start"); start != "" {
			result = addAtWord(result, start, false)
		}

		if end, _ := cmd.Flags().GetString("add-end"); end != "" {
			result = addAtWord(result, end, true)
		}

		fmt.Println(result)

	},
}

func addAtWord(text string, insert string, toEnd bool) string {
	words := strings.Fields(text)
	for i, w := range words {
		if toEnd {
			words[i] = w + insert
			continue
		}
		words[i] = insert + w
	}
	return strings.Join(words, " ")
}

func transformByPosition(text string, position int, transformFunc func(string) string) string {
	words := strings.Fields(text)
	for i, w := range words {
		runes := []rune(w)
		switch position {
		case 0:
			words[i] = transformFunc(w)
		case 1:
			if len(runes) > 0 {
				runes[0] = []rune(transformFunc(string(runes[0])))[0]
				words[i] = string(runes)
			}
		case -1:
			if len(runes) > 0 {
				runes[len(runes)-1] = []rune(transformFunc(string(runes[len(runes)-1])))[0]
				words[i] = string(runes)
			}
		default:
			if position > 0 && position <= len(runes) {
				runes[position-1] = []rune(transformFunc(string(runes[position-1])))[0]
				words[i] = string(runes)
			}
		}
	}
	return strings.Join(words, " ")
}

func init() {

	textCmd.Flags().StringP("text", "t", "", "Conteúdo do texto")
	textCmd.Flags().IntP("lower", "l", -1, "deixe o texto em minúsculo pela posição, 0 para tudo, 1 para a primeira letra, -1 para a última letra")
	textCmd.Flags().IntP("upper", "u", -1, "deixe o texto em maiúsculo pela posição, 0 para tudo, 1 para a primeira letra, -1 para a última letra")
	textCmd.Flags().StringP("add-start", "s", "", "adiciona caracteres para cada palavra do texto")
	textCmd.Flags().StringP("add-end", "e", "", "adiciona caracteres para cada palavra do texto")

	flagU := textCmd.Flags().Lookup("upper")
	flagU.NoOptDefVal = "0"
	flagL := textCmd.Flags().Lookup("lower")
	flagL.NoOptDefVal = "0"

	rootCmd.AddCommand(textCmd)
}
