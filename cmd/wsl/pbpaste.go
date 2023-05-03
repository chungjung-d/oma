/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package wsl

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// pbpasteCmd represents the pbpaste command

var pbpasteCmd = &cobra.Command{
	Use:   "pbpaste",
	Short: "Paste clipboard contents from WSL",
	Long: `This command allows you to paste the contents of the clipboard
from within the Windows Subsystem for Linux.`,
	Run: func(cmd *cobra.Command, args []string) {
		pbpaste()
	},
}

func init() {
	wslCmd.AddCommand(pbpasteCmd)

}

func pbpaste() {
	cmd := exec.Command("powershell.exe", "-c", "Get-Clipboard")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print(string(out))
}
