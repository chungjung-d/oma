/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package wsl

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open [path]",
	Short: "Open a file or directory in the default Windows application",
	Long: `This command allows you to open a file with Vscode or directory on default file Explorer in the default
Windows application for the specified file type.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}
		open(path)
	},
}

func init() {
	wslCmd.AddCommand(openCmd)
}

func open(path string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file or directory:", err)
		os.Exit(1)
	}

	var cmd *exec.Cmd
	if fileInfo.IsDir() {
		cmd = exec.Command("explorer.exe", path)
	} else {
		cmd = exec.Command("code", path)
	}

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file or directory:", err)
		os.Exit(1)
	}
}
