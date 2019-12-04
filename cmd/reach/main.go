package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kmwenja/reach"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "reach",
		Usage: "reach out and affect a host",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "data",
				Aliases: []string{"d"},
				Value:   "",
				Usage:   "inject data into the script",
			},
		},
		Action: func(c *cli.Context) error {
			sp := c.Args().First()
			if sp == "" {
				return fmt.Errorf("first argument must be a script path")
			}

			sb, err := ioutil.ReadFile(sp)
			if err != nil {
				return fmt.Errorf("could not read script path `%s`: %v", sp, err)
			}

			args := c.Args().Slice()[1:]
			data := make(map[string]interface{})

			if err := reach.Run(sb, args, data); err != nil {
				return fmt.Errorf("could not run reach script `%s`: %v", sp, err)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
