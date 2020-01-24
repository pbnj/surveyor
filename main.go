package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "surveyor",
		Usage: "A Command-Line Interface (CLI) utility for interactive terminal prompts",
		Commands: []*cli.Command{
			{
				Name:  "input",
				Usage: "Prompt user for text input",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "message",
						Aliases:  []string{"m"},
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					answer := ""
					prompt := &survey.Input{
						Message: c.String("message"),
					}
					survey.AskOne(prompt, &answer)
					fmt.Println(answer)
					return nil
				},
			},
			{
				Name:  "multiline",
				Usage: "Prompt user for multiline text input",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "message",
						Aliases:  []string{"m"},
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					answer := ""
					prompt := &survey.Multiline{
						Message: c.String("message"),
					}
					survey.AskOne(prompt, &answer)
					fmt.Println(answer)
					return nil
				},
			},
			{
				Name:  "password",
				Usage: "Prompt user for password input",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "message",
						Aliases:  []string{"m"},
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					answer := ""
					prompt := &survey.Password{
						Message: c.String("message"),
					}
					survey.AskOne(prompt, &answer)
					fmt.Println(answer)
					return nil
				},
			},
			{
				Name:  "confirm",
				Usage: "Prompt user for confirmation",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "message",
						Aliases:  []string{"m"},
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					answer := false
					prompt := &survey.Confirm{
						Message: c.String("message"),
					}
					survey.AskOne(prompt, &answer)
					fmt.Println(answer)
					return nil
				},
			},
			{
				Name:  "select",
				Usage: "Prompt user for single-choice selection",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "message",
						Aliases:  []string{"m"},
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					answer := ""
					prompt := &survey.Select{
						Message: c.String("message"),
						Options: c.Args().Slice(),
					}
					survey.AskOne(prompt, &answer)
					fmt.Println(answer)
					return nil
				},
			},
			{
				Name:  "multiselect",
				Usage: "Prompt user for multi-choice selection",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "message",
						Aliases:  []string{"m"},
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					answer := []string{}
					prompt := &survey.MultiSelect{
						Message: c.String("message"),
						Options: c.Args().Slice(),
					}
					survey.AskOne(prompt, &answer)
					fmt.Println(strings.Trim(fmt.Sprint(answer), "[]"))
					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}
