package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var paths []string

func init() {
	path := os.Getenv("path")
	paths = strings.Split(path, ";")
}

func joinPahts(p []string) string {
	return strings.Join(p, ";")
}

func main() {
	app := cli.NewApp()
	app.Name = "cpath"
	app.Version = "1.0.0"
	app.Usage = "Change current user's environment variable PATH."

	app.Commands = []cli.Command{
		{
			Name:      "list",
			ShortName: "l",
			Usage:     "List current PATH",
			Action:    List,
		},
		{
			Name:      "append",
			ShortName: "a",
			Usage:     "Append new path in PATH",
			Action:    AppendPath,
		},
		{
			Name:      "search",
			ShortName: "s",
			Usage:     "Search key word in PATH",
			Action:    Search,
		},
		{
			Name:      "delete",
			ShortName: "d",
			Usage:     "Delete count number",
			Action:    Delete,
		},
		{
			Name:      "create",
			ShortName: "c",
			Usage:     "Create name value",
			Action:    Create,
		},
	}

	app.Run(os.Args)

}

func List(c *cli.Context) {
	for i, v := range paths {
		printPath(i, v)
	}
}

func AppendPath(c *cli.Context) {
	paths = append(paths, c.Args()[0])
	updatePath(paths)
}

func Delete(c *cli.Context) {
	count, err := strconv.Atoi(c.Args()[0])

	if err != nil {
		return
	}

	if count < 0 && count >= len(paths) {
		return
	}

	dest := make([]string, len(paths)-1)
	source := paths

	j := 0
	for i := 0; i < len(source); i++ {
		if i == count {
			continue
		}

		dest[j] = source[i]
		j++
	}

	paths = dest

	updatePath(paths)
}

func updatePath(p []string) {
	np := joinPahts(p)
	cmd := exec.Command("setx", "path", np, "/M")
	cmd.Run()
}

func Search(c *cli.Context) {
	searchWord := c.Args()[0]
	for k, v := range paths {
		lower := strings.ToLower(v)
		if strings.Contains(lower, searchWord) {
			printPath(k, v)
		}
	}
}

func Create(c *cli.Context) {
	name := c.Args()[0]
	value := c.Args()[1]

	cmd := exec.Command("setx", name, value, "/M")
	cmd.Run()
}

func printPath(i int, p string) {
	fmt.Printf("%0d.\t%s\n", i, p)
}
