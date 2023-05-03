/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"oma/cmd"
	_ "oma/cmd/git"
	_ "oma/cmd/wsl"
)

func main() {
	cmd.Execute()
}
