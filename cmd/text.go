package cmd

import (
	"fmt"
	"go-convert/utils"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

var textCmd = &cobra.Command{
	Use:   "text",
	Short: "Realiza conversão de palavras",
	Long:  `Realiza conversão de palavras, como lower, upper, adição de caracter etc...`,
	Run: func(cmd *cobra.Command, args []string) {
		text, _ := cmd.Flags().GetString("text")
		useClipboard, _ := cmd.Flags().GetBool("clipboard")

		if useClipboard && text != "" {
			fmt.Println("As flags --text e --clipboard não podem ser usadas juntas.")
			return
		}

		if useClipboard {
			clipboardText, err := utils.ReadFromClipboard()
			if err != nil {
				fmt.Printf("Erro ao ler da área de transferência: %v\n", err)
				return
			}
			text = clipboardText
		}

		if len(text) <= 0 {
			fmt.Println("O tamanho do texto deve ser maior que zero.")
			return
		}

		fmt.Println("Convertendo o texto:", text)
		if cmd.Flags().Changed("lower") {
			position, _ := cmd.Flags().GetInt("lower")
			text = modifyText(text, position, strings.ToLower)
		}

		if cmd.Flags().Changed("upper") {
			position, _ := cmd.Flags().GetInt("upper")
			text = modifyText(text, position, strings.ToUpper)
		}

		if start, _ := cmd.Flags().GetString("add-start"); start != "" {
			text = modifyText(text, 0, func(word string) string {
				return start + word
			})
		}

		if end, _ := cmd.Flags().GetString("add-end"); end != "" {
			text = modifyText(text, 0, func(word string) string {
				return word + end
			})
		}

		fmt.Println(text)
		if copyToClipboard, _ := cmd.Flags().GetBool("copy"); copyToClipboard {
			if err := utils.WriteToClipboard(text); err != nil {
				fmt.Printf("Erro ao copiar para a área de transferência: %v\n", err)
				return
			}
			fmt.Println("Texto copiado para a área de transferência.")
		}
	},
}

func modifierWord(word string, position int, modifier func(string) string) string {
	if position == 0 {
		return modifier(word)
	}
	runes := []rune(word)
	lenRunes := len(runes)

	var index int
	if position < 0 {
		index = lenRunes + position
	} else {
		index = position - 1
	}

	if index < 0 || index >= lenRunes {
		return word
	}

	prefix := string(runes[:index])
	charToModify := string(runes[index])
	suffix := string(runes[index+1:])

	return prefix + modifier(charToModify) + suffix
}

func modifyText(text string, position int, modifier func(string) string) string {
	var builder strings.Builder
	builder.Grow(len(text))
	wordStartIndex := -1

	processPendingWord := func(endIndex int) {
		if wordStartIndex != -1 {
			word := text[wordStartIndex:endIndex]
			modifiedWord := modifierWord(word, position, modifier)
			builder.WriteString(modifiedWord)
			wordStartIndex = -1
		}
	}

	for i, r := range text {
		isWordChar := unicode.IsLetter(r) || unicode.IsNumber(r)

		if isWordChar && wordStartIndex == -1 {
			wordStartIndex = i
		}

		if !isWordChar {
			processPendingWord(i)
			builder.WriteRune(r)
		}
	}

	processPendingWord(len(text))
	return builder.String()
}

func init() {
	textCmd.Flags().BoolP("clipboard", "c", false, "Usar o texto da área de transferência (clipboard)")
	textCmd.Flags().BoolP("copy", "C", false, "Copia o resultado para a área de transferência")
	textCmd.Flags().StringP("text", "t", "", "Conteúdo do texto")
	textCmd.Flags().IntP("lower", "l", 0, "deixe o texto em minúsculo pela posição, 0 para tudo, 1 para a primeira letra, -1 para a última letra")
	textCmd.Flags().IntP("upper", "u", 0, "deixe o texto em maiúsculo pela posição, 0 para tudo, 1 para a primeira letra, -1 para a última letra")
	textCmd.Flags().StringP("add-start", "s", "", "adiciona caracteres para cada palavra do texto")
	textCmd.Flags().StringP("add-end", "e", "", "adiciona caracteres para cada palavra do texto")

	rootCmd.AddCommand(textCmd)
}
