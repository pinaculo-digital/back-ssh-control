package ssh_service

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

func (s *SSHService) sendCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	output := strings.TrimSpace(out.String())
	errorOutput := strings.TrimSpace(stderr.String())

	if err != nil {
		if errorOutput != "" {
			return output + "\n" + errorOutput, fmt.Errorf("command failed: %w | stderr: %s", err, errorOutput)
		}
		return output, fmt.Errorf("command failed: %w", err)
	}

	if errorOutput != "" {

		return output + "\n" + errorOutput, nil
	}

	return output, nil
}

// sendCommandAndPrintResult executa um comando (local ou remoto via ssh)
// imprime informações de execução e devolve a saída completa como string.
func (s *SSHService) sendCommandAndPrintResult(command string, args ...string) (string, error) {
	var stdoutBuf, stderrBuf bytes.Buffer

	// Se quiser que a saída também vá para o terminal em tempo real,
	// use um TeeReader ou um MultiWriter.
	cmd := exec.Command("bash", "-c", command)

	// Captura a saída para retornar como string
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	cmd.Stdin = os.Stdin

	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	fullCommand := command + " " + strings.Join(args, " ")
	fmt.Printf("\nExecutando: %s\n", fullCommand)
	fmt.Println(strings.Repeat("=", 80))

	start := time.Now()
	err := cmd.Run()
	duration := time.Since(start)

	if err != nil {
		fmt.Printf("Erro ao executar comando (durou %s): %v\n", duration.Round(time.Millisecond), err)
		// Inclui a stderr no erro para ajudar no debug
		if stderrBuf.Len() > 0 {
			fmt.Println("=== STDERR ===")
			fmt.Print(stderrBuf.String())
		}
		return "", fmt.Errorf("%w\nSTDERR:\n%s", err, stderrBuf.String())
	}

	fmt.Printf("Comando executado com sucesso em %s\n", duration.Round(time.Millisecond))
	fmt.Println(strings.Repeat("=", 80))

	// Retorna somente o stdout (padrão mais comum)
	return stdoutBuf.String(), nil
}

func (s *SSHService) formatTmux(command string) string {
	return fmt.Sprintf("const cmd = `setsid nohup bash -c '%s' </dev/null >/dev/null 2>&1 &`;", command)
}
