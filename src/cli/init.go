package cli

import (
	"fmt"
	"github.com/urfave/cli"
)

const (
	ConstAppName = "git-push"
)

//全局cliApp对象
var globalCliApp *cli.App

//存储已注册的命令
var commands []cli.Command

// Init ...
//
// @Brief  cli程序初始化
//
// @Param
//
// @Return *cli.App
//
// @Author cody.rao
//
// @Email  cody.rao@ucloud.cn
//
// @Date   2020/8/26 18:46:58
//
func AppInit() bool {
	debugHeader := "Init"
	//初始化cli对象
	globalCliApp = cli.NewApp()
	if nil == globalCliApp {
		fmt.Printf("%s: NewApp fail. error: %s\n", debugHeader, "app is nil")
		return false
	}
	//注册命令
	if !commandRegister(GitPushCmd) {
		return false
	}

	if !commandRegister(GitMakePushCmd) {
		return false
	}

	//配置cli
	globalCliApp.Name = ConstAppName
	globalCliApp.Version = "1.0.0"
	globalCliApp.Commands = commands

	return true
}

// commandRegister ...
//
// @Brief  cli命令注册
//
// @Param
//
// @Return bool
//
// @Author cody.rao
//
// @Email  cody.rao@ucloud.cn
//
// @Date   2020/8/26 18:47:23
//
func commandRegister(command *cli.Command) bool {

	if nil == commands {
		commands = make([]cli.Command, 0)
		if nil == commands {
			fmt.Println("commandRegister: error: commands make fail.")
			return false
		}
	}

	commands = append(commands, *command)

	return true
}

// GetGlobalCliApp ...
//
// @Brief  获取cli对象
//
// @Param
//
// @Return *cli.App
//
// @Author cody.rao
//
// @Email  cody.rao@ucloud.cn
//
// @Date   2020/9/2 17:47:28
//
func GetGlobalCliApp() *cli.App {
	return globalCliApp
}
