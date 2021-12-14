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

	fmt.Println("----------------------------------------")
	fmt.Println("git add start!")
	addu_cmd := exec.Command("git", "add", "-u")
	addu_out, err := addu_cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(addu_out))

	for _, info := range fileinfos {
		filename := info.Name()
		if _, ok := excluded_map[filename]; ok {
			//file in the excluded file list, do reset --
			reset_cmd := exec.Command("git", "reset", "--", filename)
			err := reset_cmd.Run()
			if err != nil {
				return err
			}
		}
	}
	fmt.Println("git add finish!")
	fmt.Println("----------------------------------------")

	fmt.Println("----------------------------------------")
	fmt.Println("git commit start!")
	commit_cmd := exec.Command("git", "commit", "-m", commit_msg)
	commit_out, err := commit_cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(commit_out))
	fmt.Println("git commit finish!")
	fmt.Println("----------------------------------------")

	fmt.Println("----------------------------------------")
	fmt.Println("git push start!")
	push_cmd := exec.Command("git", "push")
	push_out, err := push_cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(push_out))
	fmt.Println("git push finish!")
	fmt.Println("----------------------------------------")

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
