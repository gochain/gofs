package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli"

	"github.com/gochain-io/gofs"
)

func main() {
	// Interrupt cancellation.
	ctx, cancelFn := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	defer close(sigCh)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for range sigCh {
			cancelFn()
		}
	}()

	app := cli.NewApp()
	app.Name = "gofs"
	app.Version = "0.0.1"
	app.Usage = "GoChain filesystem cli tool"
	var url string
	var recursive bool
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "url",
			Usage:       "gofs api url",
			Value:       gofs.APIURL,
			Destination: &url,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "pin",
			Usage: "Pin a CID",
			Action: func(c *cli.Context) error {
				cid := c.Args().First()
				if cid == "" {
					return errors.New("missing CID")
				}
				return gofs.Pin(ctx, cid, gofs.MainnetAddress)
			},
		},
		{
			Name:  "cost",
			Usage: "Get the current storage cost in wei per GigaByteHour.",
			Action: func(c *cli.Context) error {
				return gofs.Rate(ctx, gofs.MainnetRPCURL, gofs.MainnetAddress)
			},
		},
		{
			Name:  "add",
			Usage: "Add and pin a file.",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "recursive, r",
					Usage:       "Add directory recursively",
					Destination: &recursive,
				},
			},
			Action: func(c *cli.Context) error {
				path := c.Args().First()
				if path == "" {
					return errors.New("missing file path")
				}
				return gofs.Add(ctx, url, path)
			},
		},
		{
			Name:  "status",
			Usage: "Get the current storage cost in wei per GigaByteHour.",
			Action: func(c *cli.Context) error {
				cid := c.Args().First()
				if cid == "" {
					return errors.New("missing CID")
				}
				return gofs.Status(ctx, url, cid)
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
	}
}
