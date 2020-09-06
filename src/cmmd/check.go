package cmmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"utils"
)

func GetCommittedFiles() []string {

	cmd := exec.Command("git", "show", "--name-status")
	output, err := cmd.CombinedOutput()
	if nil != err {

		fmt.Println(string(output))
		os.Exit(0)
	}
	files := string(output)

	split := strings.Split(files, "\n")
	length := len(split)
	fileNames := make([]string, 0)
	for i := length - 2; i >= 0; i-- {
		if len(split[i]) == 0 {
			break
		}
		files := strings.Split(split[i], "\t")
		if files[0] == "D" {
			continue
		}
		fileNames = append(fileNames, files[1])
	}

	return fileNames
}

func CheckFiles(project string, filePaths []string) bool {
	pwd, err := os.Getwd()
	if nil != err {
		fmt.Println(err)
		return false
	}

	projectPath := pwd[:strings.Index(pwd, project)] + project + "/"
	for _, path := range filePaths {

		if strings.HasSuffix(path, ".go") {
			GoFmt(projectPath + path)

			utils.ReadFileRow(projectPath + path)

		}

	}

	return true
}

func GoFmt(path string) {
	cmd := exec.Command("go", "fmt", path)
	output, err := cmd.CombinedOutput()
	if nil != err {

		fmt.Println(string(output))
		os.Exit(0)
	}
}
