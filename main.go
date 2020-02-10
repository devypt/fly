package main

import (
	"github.com/devypt/fly/core"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
)

var r core.Fly

// Fly cli commands
func main() {
	r.Sync = make(chan string)
	app := &cli.App{
		Name:        strings.Title(core.RPrefix),
		Version:     core.RVersion,
		Description: "Fly live reload tool",
		Commands: []*cli.Command{
			{
				Name:        "start",
				Aliases:     []string{"s"},
				Description: "Start " + strings.Title(core.RPrefix) + " on a given path. If not exist a config file it creates a new one.",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "path", Aliases: []string{"p"}, Value: ".", Usage: "Project base path"},
					&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Value: "", Usage: "Run a project by its name"},
					&cli.BoolFlag{Name: "fmt", Aliases: []string{"f"}, Value: false, Usage: "Enable go fmt"},
					&cli.BoolFlag{Name: "vet", Aliases: []string{"v"}, Value: false, Usage: "Enable go vet"},
					&cli.BoolFlag{Name: "test", Aliases: []string{"t"}, Value: false, Usage: "Enable go test"},
					&cli.BoolFlag{Name: "generate", Aliases: []string{"g"}, Value: false, Usage: "Enable go generate"},
					&cli.BoolFlag{Name: "server", Aliases: []string{"srv"}, Value: false, Usage: "Start server"},
					&cli.BoolFlag{Name: "open", Aliases: []string{"op"}, Value: false, Usage: "Open into the default browser"},
					&cli.BoolFlag{Name: "install", Aliases: []string{"i"}, Value: false, Usage: "Enable go install"},
					&cli.BoolFlag{Name: "build", Aliases: []string{"b"}, Value: false, Usage: "Enable go build"},
					&cli.BoolFlag{Name: "run", Aliases: []string{"nr"}, Value: false, Usage: "Enable go run"},
					&cli.BoolFlag{Name: "no-config", Aliases: []string{"nc"}, Value: false, Usage: "Ignore existing config and doesn't create a new one"},
				},
				Action: start,
			},
			{
				Name:        "version",
				Aliases:     []string{"v"},
				Description: "Print " + strings.Title(core.RPrefix) + " version.",
				Action: func(p *cli.Context) error {
					version()
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// Version print current version
func version() {
	log.Println(r.Prefix(core.Green.Bold(core.RVersion)))
}

// Start fly workflow
func start(c *cli.Context) (err error) {
	// check no-config and read
	if !c.Bool("no-config") {
		// read a config if exist
		r.Settings.Read(&r)
		if c.String("name") != "" {
			// filter by name flag if exist
			r.Schema.Projects = r.Schema.Filter("Name", c.String("name"))
		}
		// increase file limit
		if r.Settings.FileLimit != 0 {
			if err = r.Settings.Flimit(); err != nil {
				return err
			}
		}

	}
	// check project list length
	if len(r.Schema.Projects) == 0 {
		// create a new project based on given params
		project := r.Schema.New(c)
		// Add to projects list
		r.Schema.Add(project)
		// save config
		if !c.Bool("no-config") {
			err = r.Settings.Write(r)
			if err != nil {
				return err
			}
		}
	}

	// start workflow
	return r.Start()
}
