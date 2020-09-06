package cli

import (
	"github.com/urfave/cli"
)

var GitPushCmd = &cli.Command{
	Name:        "push",
	Aliases:     []string{"p"},
	Usage:       "git push (-f)",
	Description: "git push current branch",
	Action:      GitPush,
	Subcommands: []cli.Command{
		{
			Name:    "normal",
			Aliases: []string{"n"},
			Usage:   "normal push",
			Action:  GitPush,
		},
		{
			Name:    "force",
			Aliases: []string{"f"},
			Usage:   "force push",
			Action:  GitForcePush,
		},
	},
}

var GitMakePushCmd = &cli.Command{
	Name:        "makePush",
	Aliases:     []string{"mp"},
	Usage:       "git push with making(-f)",
	Description: "git push current branch",
	Action:      GitMakePush,
	Subcommands: []cli.Command{
		{
			Name:    "normal",
			Aliases: []string{"n"},
			Usage:   "normal push",
			Action:  GitMakePush,
		},
		{
			Name:    "force",
			Aliases: []string{"f"},
			Usage:   "force push",
			Action:  GitForceMakePush,
		},
	},
}
