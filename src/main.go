package main

import (
	cli "cli"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {

	//初始化cli
	if !cli.AppInit() {
		panic("AppInit  fail")
	}

	//执行cli命令
	cli.GetGlobalCliApp().Run(os.Args)

	//
	//ReviewProject("/home/cody/tmp/a/")
	//
	//
	//if needReview {
	//	Review()
	//}
}

func Review() {

}

func Make() {
	cmd := exec.Command("make")
	output, err := cmd.Output()
	if nil != err {
		panic(err)
	}

	fmt.Println(string(output))
}

func Clone(branch string, remoteUrl, project string) {

	cmd := exec.Command("git", "clone", "origin", remoteUrl+project, "/usr/local/goproject/read/")

	output, err := cmd.Output()
	if nil != err {
		panic(err)
	}

	fmt.Println(string(output))
}

func ShowBaseInfo() {
	cmd := exec.Command("git ", "show", "remote", "origin")

	output, err := cmd.Output()
	if nil != err {
		panic(err)
	}

	fmt.Println(string(output))
}

func ReviewGet(project string) string {
	path := "/usr/local/goproject/read/" + project
	cmd := exec.Command("cd", path)

	output, err := cmd.Output()
	if nil != err {
		panic(err)
	}

	fmt.Println(string(output))

	cmd = exec.Command("git ", "submodule", "init")

	output, err = cmd.Output()
	if nil != err {
		panic(err)
	}

	fmt.Println(string(output))

	cmd = exec.Command("git ", "submodule", "update")

	output, err = cmd.Output()
	if nil != err {
		panic(err)
	}

	fmt.Println(string(output))

	return path
}

func ReviewMake(path string) {

	cmd := exec.Command("cd", "path")

	output, err := cmd.Output()
	if nil != err {
		panic(err)
	}

	fmt.Println(string(output))

	Make()
}

func ReviewProject(pathname string) error {

	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		dir := ""
		if fi.IsDir() {
			dir = pathname + fi.Name()
			fmt.Printf("[%s]\n", pathname+fi.Name())

			err := ReviewProject(pathname + fi.Name() + "/")
			if nil != err {
				return err
			}
			continue
		}

		if strings.ToLower(fi.Name()) == "makefile" {
			ReviewMake(dir)
		}

	}

	return err

}
