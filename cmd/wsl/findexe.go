package wsl

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/karrick/godirwalk"
	"github.com/manifoldco/promptui"
	lnk "github.com/parsiya/golnk"
	"github.com/spf13/cobra"
)

var (
	exeName string
)

func init() {
	findexeCmd.Flags().StringVarP(&exeName, "name", "n", "", "Name of the .exe file")
	wslCmd.AddCommand(findexeCmd)
}

var findexeCmd = &cobra.Command{
	Use:   "findexe",
	Short: "Find and execute .exe files ",
	Run: func(cmd *cobra.Command, args []string) {
		if exeName == "" {
			fmt.Println("Please provide a .exe file name using --name flag")
			return
		}

		fileLocations := []string{
			`/mnt/c/ProgramData/Microsoft/Windows/Start Menu`,
			`/mnt/c/Users/chung/AppData/Roaming/Microsoft/Windows/Start Menu`,
		}

		foundFiles, _ := findFileAsync(fileLocations, exeName)

		if len(foundFiles) == 0 {
			fmt.Println("No .exe files found with the provided name")
			return
		}

		selectedFile, err := promptUser(foundFiles)
		if err != nil {
			fmt.Printf("Error during selection: %v\n", err)
			return
		}

		if strings.HasSuffix(selectedFile, ".lnk") {

			err := runWindowsLnk(selectedFile)
			if err != nil {
				fmt.Printf("Error while executing the file: %v\n", err)
				return
			}
		}

		if strings.HasSuffix(selectedFile, ".exe") {
			exeCommand := exec.Command(selectedFile)
			err = exeCommand.Run()
			if err != nil {
				fmt.Printf("Error while executing the file: %v\n", err)
				return
			}
		}

	},
}

func runWindowsLnk(lnkPath string) error {
	// 1. WSL 경로를 Windows 경로로 변환

	// 2. PowerShell을 사용하여 .lnk 파일의 대상을 얻음
	target := getLnkTarget(lnkPath)

	// 3. 대상 경로를 WSL 경로로 변환
	wslTarget := strings.TrimSpace(string(target))
	wslTarget = "/mnt/" + strings.ToLower(string(wslTarget[0])) + wslTarget[2:]
	wslTarget = strings.ReplaceAll(wslTarget, "\\", "/")

	// 4. 명령 실행
	execCmd := exec.Command(wslTarget)
	if err := execCmd.Run(); err != nil {
		return fmt.Errorf("failed to execute command: %w", err)
	}
	return nil
}

func getLnkTarget(lnkPath string) string {

	Lnk, err := lnk.File(lnkPath)
	if err != nil {
		panic(err)
	}
	return Lnk.LinkInfo.LocalBasePath

}

func findFileAsync(fileLocations []string, exeName string) ([]string, error) {
	var foundFiles []string
	var wg sync.WaitGroup

	workerPool := make(chan struct{}, 20)

	for i := 0; i < cap(workerPool); i++ {
		workerPool <- struct{}{}
	}

	for _, fileLocation := range fileLocations {

		wg.Add(1)
		<-workerPool

		go func(location string) {
			defer func() {
				workerPool <- struct{}{}
				wg.Done()
			}()
			_ = godirwalk.Walk(location, &godirwalk.Options{
				Callback: func(path string, de *godirwalk.Dirent) error {
					if !de.IsDir() && strings.Contains(strings.ToLower(de.Name()), strings.ToLower(exeName)) && (strings.HasSuffix(de.Name(), ".exe") || strings.HasSuffix(de.Name(), ".lnk")) {
						foundFiles = append(foundFiles, path)
					}
					return nil
				},

				ErrorCallback: func(path string, err error) godirwalk.ErrorAction {
					return godirwalk.SkipNode
				},
				Unsorted: true,
			})
		}(fileLocation)
	}

	wg.Wait()
	return foundFiles, nil
}

func promptUser(foundFiles []string) (string, error) {
	// 파일 이름만 있는 슬라이스 생성
	fileNames := make([]string, len(foundFiles))
	for i, fullPath := range foundFiles {
		fileNames[i] = extractFileName(fullPath)
	}

	// 사용자 정의 항목 템플릿 생성
	itemTemplate := &promptui.SelectTemplates{
		Active:   `{{ "▸" | cyan }} {{ . | cyan }}`,
		Inactive: "  {{ . }}",
	}

	// 사용자에게 파일 이름만 보여주는 프롬프트 생성
	prompt := promptui.Select{
		Label:     "Select a file to execute",
		Items:     fileNames,
		Templates: itemTemplate,
	}

	// 사용자가 선택한 파일 이름의 인덱스를 가져옴
	selectedIndex, _, err := prompt.Run()
	if err != nil {
		return "", err
	}

	// 원래 경로를 사용하여 결과 반환
	return foundFiles[selectedIndex], nil
}

func extractFileName(fullPath string) string {
	_, fileName := filepath.Split(fullPath)
	return fileName
}
