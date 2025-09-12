package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func ReadFromClipboard() (string, error) {
	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("xclip", "-selection", "clipboard", "-o")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			cmd = exec.Command("xsel", "--clipboard", "--output")
			var outSel bytes.Buffer
			cmd.Stdout = &outSel
			errSel := cmd.Run()
			if errSel != nil {
				return "", fmt.Errorf("falha ao executar xclip e xsel. Certifique-se de que um deles está instalado")
			}
			return outSel.String(), nil
		}
		return out.String(), nil
	case "darwin":
		cmd := exec.Command("pbpaste")
		output, err := cmd.Output()
		if err != nil {
			return "", fmt.Errorf("falha ao executar pbpaste: %w", err)
		}
		return string(output), nil
	case "windows":
		cmd := exec.Command("powershell", "-command", "Get-Clipboard")
		output, err := cmd.Output()
		if err != nil {
			return "", fmt.Errorf("falha ao executar o comando do PowerShell Get-Clipboard: %w", err)
		}
		return strings.TrimSpace(string(output)), nil
	default:
		return "", fmt.Errorf("sistema operacional não suportado: %s", runtime.GOOS)
	}
}

func WriteToClipboard(text string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xclip", "-selection", "clipboard")
		cmd.Stdin = strings.NewReader(text)
		err := cmd.Run()
		if err != nil {

			cmd = exec.Command("xsel", "--clipboard", "--input")
			cmd.Stdin = strings.NewReader(text)
			if errSel := cmd.Run(); errSel != nil {
				return fmt.Errorf("falha ao executar xclip e xsel. Certifique-se de que um deles está instalado")
			}
		}
		return nil
	case "darwin":
		cmd = exec.Command("pbcopy")
		cmd.Stdin = strings.NewReader(text)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("falha ao executar pbcopy: %w", err)
		}
		return nil
	case "windows":
		cmd = exec.Command("powershell", "-command", "$input|Set-Clipboard")
		cmd.Stdin = strings.NewReader(text)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("falha ao executar o comando do PowerShell Set-Clipboard: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("sistema operacional não suportado para cópia: %s", runtime.GOOS)
	}
}
