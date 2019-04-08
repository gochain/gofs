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

	"github.com/gochain-io/gochain/v3/common"
	"github.com/gochain-io/gochain/v3/core/types"
	"github.com/gochain-io/web3"
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
	var api, rpc, contract string
	var recursive bool
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "url",
			Usage:       "GOFS API URL.",
			Value:       gofs.APIURL,
			EnvVar:      "GOFS_API",
			Destination: &api,
		},
		cli.StringFlag{
			Name:        "contract",
			Usage:       "Contract address.",
			Value:       gofs.MainnetAddress.Hex(),
			EnvVar:      "GOFS_CONTRACT",
			Destination: &contract,
		},
		cli.StringFlag{
			Name:        "rpc",
			Usage:       "RPC URL.",
			Value:       gofs.MainnetRPCURL,
			EnvVar:      "GOFS_RPC",
			Destination: &rpc,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "pin",
			Usage: "Pin a CID",
			Flags: []cli.Flag{
				cli.Uint64Flag{
					Name:  "duration, d",
					Usage: "Storage duration in hours.",
					Value: 1,
				},
			},
			Action: func(c *cli.Context) error {
				cid := c.Args().First()
				if cid == "" {
					return errors.New("missing CID")
				}
				dur := c.Uint64("duration")
				if dur == 0 {
					return fmt.Errorf("duration missing or invalid")
				}
				contract, err := parseContract(c)
				if err != nil {
					return err
				}
				return Pin(ctx, rpc, contract, cid, dur)
			},
		},
		{
			Name:  "rate",
			Usage: "Get the current storage rate in wei per GigaByteHour.",
			Action: func(c *cli.Context) error {
				contract, err := parseContract(c)
				if err != nil {
					return err
				}
				return Rate(ctx, rpc, contract)
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
				contract, err := parseContract(c)
				if err != nil {
					return err
				}
				return Cost(ctx, rpc, contract, size, dur)
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
				return Add(ctx, api, path)
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
				return Status(ctx, api, cid)
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
	}
}

func parseContract(c *cli.Context) (common.Address, error) {
	contract := c.String("contract")
	if !common.IsHexAddress(contract) {
		return common.Address{}, fmt.Errorf("invalid hex contract address: %s", contract)
	}
	return common.HexToAddress(contract), nil
}

func Pin(ctx context.Context, url string, contract common.Address, ci string, dur uint64) error {
	h, r, err := gofs.Pin(ctx, url, contract, ci, dur)
	if err != nil {
		return fmt.Errorf("failed to get receipt for tx %s: %v", h.Hex(), err)
	}
	switch r.Status {
	case types.ReceiptStatusFailed:
		return fmt.Errorf("tx %s failed", h.Hex())
	case types.ReceiptStatusSuccessful:
		fmt.Printf("Purchased %d GigaByte Hours of storage for %s with tx %s.\n", dur, ci, h.Hex())
		return nil
	default:
		return fmt.Errorf("tx %s unrecognized receipt status: %d", h.Hex(), r.Status)
	}
}

func Add(ctx context.Context, url, path string) error {
	ar, err := gofs.AddFile(ctx, url, path)
	if err != nil {
		return fmt.Errorf("failed to add file %q: %v", path, err)
	}
	fmt.Println("File uploaded and pinned.")
	fmt.Println("Pinned until:", time.Unix(ar.Expiration, 0))
	return nil
}

func Cost(ctx context.Context, rpcurl string, contract common.Address, size, dur uint64) error {
	_, cost, err := gofs.Cost(ctx, rpcurl, contract, size, dur)
	if err != nil {
		return err
	}

	fmt.Println(costStr(size, dur, cost))

	return nil
}

func costStr(size uint64, dur uint64, cost *big.Int) string {
	return fmt.Sprintf("%d GBs for %s: %s", size, time.Duration(dur)*time.Hour, web3.WeiAsBase(cost))
}

func Rate(ctx context.Context, rpcurl string, contract common.Address) error {
	rate, err := gofs.Rate(ctx, rpcurl, contract)
	if err != nil {
		return err
	}
	//TODO friendlier units?
	fmt.Printf("Current storage rate: %d wei per GigaByteHour.", rate)

	fmt.Println("Cost:")
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

		fmt.Println("\t", costStr(vals.gbs, vals.hrs, cost))
	}
	return nil
}

func Status(ctx context.Context, url, ci string) error {
	st, err := gofs.Status(ctx, url, ci)
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
