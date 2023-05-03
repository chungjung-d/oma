/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package wsl

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Help for oma wsl",
	Long:  `This command allows you to get help on how to use oma wsl.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("oma wsl is a command for Windows Subsystem for Linux.")

		fmt.Println("===============================================")
		fmt.Println("")
		fmt.Println("open \t\t\t - Open a file or directory in the default Windows application.")
		fmt.Println("")
		fmt.Println("Usage: \toma wsl open [path]\tOpen the file or directory at the given path in the default Windows application.")
		fmt.Println("Example:\toma wsl open /mnt/c/Users/John/Documents/\tOpen the Documents directory in the default Windows file explorer.")

		fmt.Println("")
		fmt.Println("")

		fmt.Println("===============================================")
		fmt.Println("")
		fmt.Println("pbcopy \t\t\t - Copy the content to the clipboard.")
		fmt.Println("")
		fmt.Println("Usage: \t echo [content] | oma wsl pbcopy \tCopy the content to the clipboard.")
		fmt.Println("Example:\t cat ~/.zshrc | oma wsl pbcopy \tCopy the content of the .zshrc file to the clipboard.")
		fmt.Println("")
		fmt.Println("")

		fmt.Println("===============================================")
		fmt.Println("")
		fmt.Println("pbpaste \t\t\t - Paste the content from the clipboard.")
		fmt.Println("")
		fmt.Println("Usage: \t oma wsl pbpaste \tPaste the content from the clipboard.")
		fmt.Println("Example:\t oma wsl pbpaste > a.txt \tPaste the content from the clipboard to the a.txt file.")
		fmt.Println("")
		fmt.Println("")
	},
}

func init() {
	wslCmd.AddCommand(helpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
