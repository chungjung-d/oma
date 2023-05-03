package wsl

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var pbcopyCmd = &cobra.Command{
	Use:   "pbcopy",
	Short: "Copy text from stdin to clipboard",
	Long: `Command pbcopy copies the standard input to the clipboard.
	example: echo "hello" | oma wsl pbcopy`,
	Run: func(cmd *cobra.Command, args []string) {

		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
			os.Exit(1)
		}

		copyToClipboard(text)
	},
}

func init() {
	wslCmd.AddCommand(pbcopyCmd)
}

func copyToClipboard(text string) {
	// Windows의 clip.exe를 사용하여 클립보드로 텍스트를 전송합니다.
	cmd := exec.Command("clip.exe")
	cmd.Stdin = bytes.NewBufferString(text)

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error copying to clipboard:", err)
		os.Exit(1)
	}
}
