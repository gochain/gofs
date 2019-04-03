package main

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gochain-io/gochain/v3/accounts/abi/bind"
	"github.com/gochain-io/gochain/v3/common"
	"github.com/gochain-io/gochain/v3/goclient"
	"github.com/gochain-io/web3"
	cid "github.com/ipfs/go-cid"
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
				return Pin(ctx, cid, gofs.MainnetAddress)
			},
		},
		{
			Name:  "rate",
			Usage: "Get the current storage rate in wei per GigaByteHour.",
			Action: func(c *cli.Context) error {
				return Rate(ctx, gofs.MainnetRPCURL, gofs.MainnetAddress)
			},
		},
		{
			Name:  "cost",
			Usage: "Get the current storage cost in wei for the given GigaByteHour.",
			Flags: []cli.Flag{
				cli.Uint64Flag{
					Name:  "duration, d",
					Usage: "Storage duration in hours.",
					Value: 1,
				},
				cli.Uint64Flag{
					Name:  "size, s",
					Usage: "Storage Size in GigaBytes.",
					Value: 1,
				},
			},
			Action: func(c *cli.Context) error {
				dur := c.Uint64("duration")
				if dur == 0 {
					return fmt.Errorf("duration missing or invalid")
				}
				size := c.Uint64("size")
				if size == 0 {
					return fmt.Errorf("size missing or invalid")
				}
				return Cost(ctx, gofs.MainnetRPCURL, gofs.MainnetAddress, size, dur)
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
				return Add(ctx, url, path)
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
				return Status(ctx, url, cid)
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
	}
}

func Pin(ctx context.Context, ci string, contract common.Address) error {
	cid, err := cid.Parse(ci)
	if err != nil {
		return err
	}
	if cid.Version() == 0 {
		return errors.New("version 0 CID not supported")
	}
	gc, err := goclient.Dial(gofs.MainnetRPCURL)
	if err != nil {
		return err
	}
	p, err := gofs.NewPinner(contract, gc)
	if err != nil {
		return err
	}
	opts := &bind.TransactOpts{
		Context: ctx,
		//TODO set value
	}
	_, err = p.Pin(opts, cid.Bytes())
	if err != nil {
		return fmt.Errorf("failed to pin %q: %v", cid, err)
	}
	//TODO wait for receipt, parse log
	var gbh int64
	fmt.Printf("Purchased %d GigaByte Hours of storage for %s.\n", gbh, cid)
	return nil
}

func Add(ctx context.Context, url, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file %q: %v", path, err)
	}
	defer f.Close()
	ar, err := gofs.NewClient(url).Add(ctx, f)
	if err != nil {
		return fmt.Errorf("failed to add file %q: %v", path, err)
	}
	fmt.Println("File uploaded and pinned.")
	fmt.Println("Pinned until:", time.Unix(ar.Expiration, 0))
	return nil
}

// Cost calculates the cost of storage size Giga Bytes for dur hours at the current rate.
func Cost(ctx context.Context, rpcurl string, contract common.Address, size, dur uint64) error {
	gc, err := goclient.Dial(rpcurl)
	if err != nil {
		return err
	}
	p, err := gofs.NewPinner(contract, gc)
	if err != nil {
		return err
	}
	rate, err := p.Rate(&bind.CallOpts{Context: ctx})
	if err != nil {
		return err
	}
	gbhs := size * dur
	cost := new(big.Int).Mul(rate, big.NewInt(int64(gbhs)))

	fmt.Printf(costStr(size, dur, cost))

	return nil
}

func costStr(size uint64, dur uint64, cost *big.Int) string {
	return fmt.Sprintf("%d GBs for %s: %s", size, time.Duration(dur)*time.Hour, web3.WeiAsBase(cost))
}

func Rate(ctx context.Context, rpcurl string, contract common.Address) error {
	gc, err := goclient.Dial(rpcurl)
	if err != nil {
		return err
	}
	p, err := gofs.NewPinner(contract, gc)
	if err != nil {
		return err
	}
	rate, err := p.Rate(&bind.CallOpts{Context: ctx})
	if err != nil {
		return err
	}
	fmt.Printf("Current storage rate: %d wei per GigaByteHour.", rate)
	//TODO optionally accept params for gbs and hrs instead
	for _, vals := range []struct {
		gbs uint64
		hrs uint64
	}{
		{gbs: 1, hrs: 1},
		{gbs: 10, hrs: 1},
		{gbs: 1000, hrs: 1},
		{gbs: 1, hrs: 24},
		{gbs: 1, hrs: 24 * 7},
		{gbs: 1, hrs: 24 * 7 * 52},
		{gbs: 10, hrs: 24 * 7 * 52},
		{gbs: 1000, hrs: 24 * 7 * 52},
	} {
		gbh := big.NewInt(int64(vals.gbs * vals.hrs))
		cost := new(big.Int).Mul(gbh, rate)

		fmt.Print(fmt.Printf(costStr(vals.gbs, vals.hrs, cost)))
	}
	return nil
}

func Status(ctx context.Context, url, ci string) error {
	cid, err := cid.Decode(ci)
	if err != nil {
		return err
	}
	if cid.Version() == 0 {
		return errors.New("version 0 CID not supported")
	}
	st, err := gofs.NewClient(url).Status(ctx, cid)
	if err != nil {
		return err
	}
	if st.Expiration == 0 {
		fmt.Println("Never been pinned.")
		return nil
	}
	exp := time.Unix(st.Expiration, 0)
	if until := time.Until(exp); until > 0 {
		fmt.Printf("Expires in %s at %s.\n", until, exp)
	} else {
		fmt.Printf("Expired %s ago at %s.\n", -until, exp)
	}

	return nil
}
