package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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
			Action: func(c *cli.Context) {
				cid := c.Args().First()
				if cid == "" {
					log.Fatalln("Missing CID")
				}
				Pin(ctx, url, cid)
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
			Action: func(c *cli.Context) {
				path := c.Args().First()
				if path == "" {
					log.Fatalln("Missing file path")
				}
				Add(ctx, url, path)
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func Pin(ctx context.Context, url, cid string) {
	c := gofs.NewClient(url)
	pr, err := c.Pin(ctx, cid)
	if err != nil {
		log.Fatalf("Failed to pin %q: %v", cid, err)
	}
	if pr.Pinned {
		fmt.Println("File pinned.")
	} else {
		fmt.Println("Pinning in progress.")
	}
	fmt.Println("Pinned until:", time.Unix(pr.Expiration, 0))
}

func Add(ctx context.Context, url, path string) {
	c := gofs.NewClient(url)
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open file %q: %v", path, err)
	}
	defer f.Close()
	ar, err := c.Add(ctx, f)
	if err != nil {
		log.Fatalf("Failed to add file %q: %v", path, err)
	}
	fmt.Println("File uploaded and pinned.")
	fmt.Println("Pinned until:", time.Unix(ar.Expiration, 0))
}
