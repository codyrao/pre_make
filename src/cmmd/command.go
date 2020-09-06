package cmmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var Project string

const (
	Path = "/usr/local/review/"
)

func GitBranch() string {
	cmd := exec.Command("git", "branch")

	output, err := cmd.CombinedOutput()
	if nil != err {

		fmt.Println(string(output))
		os.Exit(0)
	}
	branches := string(output)

	branchSplit := strings.Split(branches, "\n")
	for _, branch := range branchSplit {
		if strings.HasPrefix(branch, "*") {
			return branch[2:]
		}
	}

	fmt.Println("no branch")
	os.Exit(0)
	return ""
}

func GitPush(branch string, force bool) {
	var cmd *exec.Cmd
	if force {
		cmd = exec.Command("git", "push", "origin", branch, "-f")
	} else {
		cmd = exec.Command("git", "push", "origin", branch)
	}
	output, err := cmd.CombinedOutput()
	if nil != err {

		fmt.Println(string(output))
		os.Exit(0)
	}
	fmt.Println(string(output))

}

func Make() {
	cmd := exec.Command("make")
	output, err := cmd.CombinedOutput()
	if nil != err {

		fmt.Println(string(output))
		os.Exit(0)
	}

	fmt.Println(string(output))
}

func ReviewProjectMake(path, currentDir string) bool {

	rd, _ := ioutil.ReadDir(path)

	for _, fi := range rd {
		if fi.IsDir() && fi.Name() == currentDir {

			ReviewMake(path + fi.Name())
			return true

		}

		if fi.IsDir() {
			flag := ReviewProjectMake(path+fi.Name()+"/", currentDir)
			if flag {
				return true
			}

		}

	}

	return false

}

func ReviewMake(path string) {
	cmd := exec.Command("make")
	cmd.Dir = path
	cmd.Env = append(os.Environ(), "GOPATH="+Path+Project)
	output, err := cmd.CombinedOutput()
	if nil != err {

		fmt.Println(string(output))
		os.Exit(0)
	}

	fmt.Println(string(output))

	fmt.Println("review make success!")
}

func Clone(branch, remoteUrl, project string) {

	cmd := exec.Command("sudo", "mkdir", "-p", Path)

	output, err := cmd.CombinedOutput()
	if nil != err {

		fmt.Println(string(output))
		os.Exit(0)
	}

	cmd = exec.Command("sudo", "chmod", "777", Path)

	output, err = cmd.CombinedOutput()
	if nil != err {

		fmt.Println(string(output))
		os.Exit(0)
	}

	RM(project)
	cmd = exec.Command("git", "clone", remoteUrl)

	cmd.Dir = Path

	output, err = cmd.CombinedOutput()
	if nil != err {

		fmt.Println(string(output))
		os.Exit(0)
	}
	fmt.Println(string(output))
	fmt.Println("clone " + branch + " success!")
}

func RM(project string) {
	if project == "" {
		fmt.Println("project is nil")
		os.Exit(0)
	}

	cmd := exec.Command("sudo", "rm", "-rf", Path+project)
	output, err := cmd.CombinedOutput()
	if nil != err {

		fmt.Println(string(output))
		os.Exit(0)
	}
}

func ReviewProjectInit(branch, project string) string {

	path := Path + project
	cmd := exec.Command("git", "pull", "origin", branch)
	cmd.Dir = path
	output, err := cmd.CombinedOutput()
	if nil != err {

		fmt.Println(string(output))
		os.Exit(0)
	}
	fmt.Println("git pull origin " + branch + " success!")

	cmd = exec.Command("git", "checkout", branch)
	cmd.Dir = path
	output, err = cmd.CombinedOutput()
	if nil != err {

		fmt.Println(string(output))
		os.Exit(0)
	}

	fmt.Println("git checkout " + branch + " success!")

	cmd = exec.Command("git", "submodule", "init")
	cmd.Dir = path
	output, err = cmd.CombinedOutput()
	if nil != err {

		fmt.Println(string(output))
		os.Exit(0)
	}

	fmt.Println("git submodule init success!")

	cmd = exec.Command("git", "submodule", "update")
	cmd.Dir = path
	output, err = cmd.CombinedOutput()
	if nil != err {

		fmt.Println(string(output))
		os.Exit(0)
	}

	fmt.Println("git submodule update success!")

	fmt.Println("review project:" + project + "[" + branch + "]" + " init success!")
	return path
}

func GitRemoteShow() (string, string) {
	cmd := exec.Command("git", "remote", "-v")

	output, err := cmd.CombinedOutput()
	if nil != err {

		fmt.Println(string(output))
		os.Exit(0)
	}
	remote := strings.Split(string(output), "\n")[0]
	split := strings.Split(strings.Replace(remote, "\t", " ", -1), " ")
	name := split[0]
	if name != "origin" {
		fmt.Printf("warning: remote name:%s\n", name)
		os.Exit(0)
	}

	url := split[1]
	split = strings.Split(url, "/")
	projectGit := split[len(split)-1]
	project := projectGit[:len(projectGit)-4]
	return url, project
}

func GetCurrentDirName() string {
	cmd := exec.Command("pwd")
	output, err := cmd.CombinedOutput()
	if nil != err {

		fmt.Println(string(output))
		os.Exit(0)
	}
	split := strings.Split(string(output), "/")
	return strings.TrimRight(split[len(split)-1], "\n")
}
