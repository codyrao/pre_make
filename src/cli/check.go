package cli

import (
	"cmmd"
	"fmt"
)

func Check(project string) {

	cmmd.CheckFiles(project, cmmd.GetCommittedFiles())

	fmt.Println("代码规范检查完成!")
}
