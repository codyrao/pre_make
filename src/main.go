package main

import (
	cli "cli"
	"os"
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
