package cmd

import (
	"fmt"
	"os/exec"
)

func GitBranch() string {
	cmd := exec.Command("git", "branch")

	output, err := cmd.Output()
	if nil != err {
		panic(err)
	}
	branch := string(output)
	fmt.Println(branch)
	return branch
}

func GitPush(branch string, args string) {

	cmd := exec.Command("git", "push", branch, "-"+args)

	output, err := cmd.Output()
	if nil != err {
		panic(err)
	}

	fmt.Println(string(output))

}
