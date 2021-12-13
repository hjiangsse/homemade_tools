package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func NameChangeAction(c *cli.Context) error {
	fromstr := c.String("from")
	deststr := c.String("to")

	fileinfos, err := ioutil.ReadDir(".")
	if err != nil {
		return err
	}

	for _, info := range fileinfos {
		file, err := os.Stat(info.Name())
		if err != nil {
			return err
		}

		fileMode := file.Mode()
		if fileMode.IsRegular() {
			originName := file.Name()
			newName := strings.Replace(originName, fromstr, deststr, -1)
			fmt.Printf("OriginName: %s, rename to: %s\n", originName, newName)
			os.Rename(originName, newName)
		}
	}
	return nil
}

func ContentChangeAction(c *cli.Context) error {
	fromstr := c.String("from")
	deststr := c.String("to")

	fileinfos, err := ioutil.ReadDir(".")
	if err != nil {
		return err
	}

	for _, info := range fileinfos {
		file, err := os.Stat(info.Name())
		if err != nil {
			return err
		}

		fileMode := file.Mode()
		if fileMode.IsRegular() {
			originContents, err := ioutil.ReadFile(info.Name())
			if err != nil {
				return err
			}

			newContents := strings.Replace(string(originContents), fromstr, deststr, -1)
			err = ioutil.WriteFile(info.Name(), []byte(newContents), 0)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:     "from",
			Value:    "",
			Usage:    "Departure of the change",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "to",
			Value:    "",
			Usage:    "Destination of the change",
			Required: true,
		},
	}

	app := cli.NewApp()
	app.Commands = []*cli.Command{
		{
			Name:    "name",
			Aliases: []string{"n"},
			Usage:   "change file names in `PWD`",
			Flags:   flags,
			Action:  NameChangeAction,
		},
		{
			Name:    "content",
			Aliases: []string{"c"},
			Usage:   "change file contents in `PWD`",
			Flags:   flags,
			Action:  ContentChangeAction,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
