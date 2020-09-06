package cli

import (
	"cmmd"
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func GitPush(c *cli.Context) {
	_, project := cmmd.GitRemoteShow()
	Check(project)
	branch := cmmd.GitBranch()
	cmmd.GitPush(branch, false)
	fmt.Println("git push " + branch + " success!")
}

func GitForcePush(c *cli.Context) {
	_, project := cmmd.GitRemoteShow()
	Check(project)
	branch := cmmd.GitBranch()

	cmmd.GitPush(branch, true)
	fmt.Println("git force push " + branch + " success!")
}

func GitForceMakePush(c *cli.Context) {
	url, project := cmmd.GitRemoteShow()
	branch := cmmd.GitBranch()
	PreMake(branch)

	GitForcePush(c)

	PostMake(url, project, branch)
}

func GitMakePush(c *cli.Context) {
	url, project := cmmd.GitRemoteShow()
	branch := cmmd.GitBranch()
	PreMake(branch)

	GitPush(c)

	PostMake(url, project, branch)
}

func PreMake(branch string) {
	cmmd.Make()
	fmt.Println("pre make " + branch + " success!")
}
func PostMake(url, project, branch string) {
	cmmd.Project = project
	currentDir := cmmd.GetCurrentDirName()
	cmmd.Clone(branch, url, project)

	cmmd.ReviewProjectInit(branch, project)
	path := cmmd.Path + project + "/"
	if !cmmd.ReviewProjectMake(path, currentDir) {
		fmt.Println(currentDir + ": no makefile!")
		os.Exit(0)
	}
	fmt.Println("post make " + branch + " success!")
}
