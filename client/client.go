package main

import (
	"fmt"
	"gcloud/gcorecloud-go/client/flags"
	"gcloud/gcorecloud-go/client/magnum/clusters"
	"gcloud/gcorecloud-go/client/magnum/templates"
	"gcloud/gcorecloud-go/client/networks"
	"gcloud/gcorecloud-go/client/tasks"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {

	app := cli.NewApp()
	app.Version = "0.0.1"
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		{
			Name:  "password",
			Usage: fmt.Sprintf("GCloud API client\n%s", flags.PasswordClientHelpText),
			Flags: flags.PasswordClientFlags,
			Before: func(c *cli.Context) error {
				return c.Set("client-type", "password")
			},
			Subcommands: []*cli.Command{
				&networks.NetworkCommands,
				&tasks.TaskCommands,
				{
					Name:  "magnum",
					Usage: "Magnum commands",
					Subcommands: []*cli.Command{
						&clusters.ClusterCommands,
						&templates.ClusterTemplatesCommands,
					},
				},
			},
		},
		{
			Name:  "token",
			Usage: fmt.Sprintf("GCloud API client\n%s", flags.TokenClientHelpText),
			Flags: flags.TokenClientFlags,
			Subcommands: []*cli.Command{
				&networks.NetworkCommands,
				&tasks.TaskCommands,
				{
					Name:  "magnum",
					Usage: "Magnum commands",
					Subcommands: []*cli.Command{
						&clusters.ClusterCommands,
						&templates.ClusterTemplatesCommands,
					},
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		logrus.Errorf("Cannot initialize application: %+v", err)
		os.Exit(1)
	}
}