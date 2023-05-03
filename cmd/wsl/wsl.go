package wsl

import (
	"fmt"
	"oma/cmd"

	"github.com/spf13/cobra"
)

// wslCmd represents the wsl command
var wslCmd = &cobra.Command{
	Use:   "wsl",
	Short: "Comamd for Windows Subsystem for Linux",
	Long:  `Command wsl is a command for Windows Subsystem for Linux.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please use the following command to get help on how to use oma wsl:")
		fmt.Println("oma wsl help")
	},
}

func init() {

	cmd.GetRootCommandInstance().AddCommand(wslCmd)

}
