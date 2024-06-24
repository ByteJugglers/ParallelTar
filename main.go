package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:           "ParallelTar",
		Usage:          "Tar files parallel.",
		Version:        "1.0.0",
		DefaultCommand: "untar",
		Commands: []*cli.Command{
			{
				Name:  "tar",
				Usage: "tar files or directories",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "dir",
						Aliases: []string{"d"},
						Value:   "",
						Usage:   "The parent directory that needs to be packaged",
					},
					&cli.IntFlag{
						Name:    "jobs",
						Aliases: []string{"j"},
						Value:   2,
						Usage:   "Concurrent jobs",
					},
					&cli.StringFlag{
						Name:    "type",
						Aliases: []string{"t"},
						Value:   "gzip",
						Usage:   "Compressing type gzip/xz/bzip2",
					},
				},
				Action: func(c *cli.Context) error {
					svc := Service{
						Dir:  c.String("dir"),
						Job:  c.Int("jobs"),
						Type: c.String("type"),
					}
					svc.Tar()

					return nil
				},
			},
			{
				Name:  "untar",
				Usage: "Unzip files in the directory",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "dir",
						Aliases: []string{"d"},
						Value:   "test",
						Usage:   "The parent directory that needs to be un-packaged",
					},
					&cli.IntFlag{
						Name:    "jobs",
						Aliases: []string{"j"},
						Value:   1,
						Usage:   "Concurrent jobs",
					},
				},
				Action: func(c *cli.Context) error {
					svc := Service{
						Dir: c.String("dir"),
						Job: c.Int("jobs"),
					}
					svc.UnTar()

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
