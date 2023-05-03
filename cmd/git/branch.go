package git

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func init() {
	gitCmd.AddCommand(branchCmd)
}

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Create a new branch with a given prefix",
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "Select a prefix",
			Items: []string{"feature", "bug", "hotfix"},
		}

		_, prefix, err := prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed: %v\n", err)
		}

		if len(args) < 1 {
			log.Fatalln("You must provide a branch name")
		}

		branchName := fmt.Sprintf("%s/%s", prefix, args[0])

		repo, err := git.PlainOpen(".")
		if err != nil {
			log.Fatalf("Cannot open repository: %v\n", err)
		}

		headRef, err := repo.Head()
		if err != nil {
			log.Fatalf("Cannot get HEAD reference: %v\n", err)
		}

		branch := git

		err = repo.CreateBranch(&git.config.Branch{
			Name:  branchName,
			Force: false,
			Head:  headRef,
		})
		if err != nil {
			log.Fatalf("Cannot create branch: %v\n", err)
		}

		fmt.Printf("Created branch: %s\n", branchName)
	},
}
