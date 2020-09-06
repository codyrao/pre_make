package cli

import (
	"fmt"
	"github.com/urfave/cli"
)

func GitPush(c *cli.Context) {

	fmt.Println("GitPush")
}

func GitForceMakePush(c *cli.Context) {

	fmt.Println("GitForceMakePush")
}

func GitMakePush(c *cli.Context) {

	fmt.Println("GitMakePush")
}

func GitForcePush(c *cli.Context) {
	fmt.Println("GitForcePush")
}
