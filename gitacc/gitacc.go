package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
)

func GetExcludedFileMap(files string) map[string]int {
	file_list := strings.Split(files, ",")
	filemap := make(map[string]int)

	for _, file := range file_list {
		trimed_file := strings.TrimSpace(file)
		filemap[trimed_file] = 1
	}
	return filemap
}

func FastPushAction(c *cli.Context) error {
	excluded_files := c.String("exclude")
	excluded_map := GetExcludedFileMap(excluded_files)

	commit_msg := c.String("message")

	fileinfos, err := ioutil.ReadDir(".")
	if err != nil {
		return err
	}

	for _, info := range fileinfos {
		filename := info.Name()
		if _, ok := excluded_map[filename]; !ok {
			//file not in the excluded file list, do add commit and push
			add_cmd := exec.Command("git", "add", filename)
			err := add_cmd.Run()
			if err != nil {
				return err
			}
		}
	}

	commit_cmd := exec.Command("git", "commit", "-m", commit_msg)
	commit_out, err := commit_cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(commit_out))

	push_cmd := exec.Command("git", "push")
	push_out, err := push_cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(push_out))

	return nil
}

func main() {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:     "exclude",
			Value:    "",
			Usage:    "excluded files when do git add *, [file1, file2, file3]",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "message",
			Value:    "just a simple commit",
			Usage:    "message for this commit",
			Required: false,
		},
	}

	app := cli.NewApp()
	app.Commands = []*cli.Command{
		{
			Name:    "fastpush",
			Aliases: []string{"fp"},
			Usage:   "fast push",
			Flags:   flags,
			Action:  FastPushAction,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
