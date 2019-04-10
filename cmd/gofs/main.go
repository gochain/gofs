package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"text/tabwriter"
	"time"

	"github.com/gochain-io/gochain/v3/common/hexutil"

	"github.com/gochain-io/gofs"

	"github.com/gochain-io/gochain/v3/common"
	"github.com/gochain-io/gochain/v3/core/types"
	"github.com/gochain-io/gochain/v3/crypto"
	"github.com/gochain-io/web3"
	cid "github.com/ipfs/go-cid"
	"github.com/urfave/cli"
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
				},
				cli.StringFlag{
					Name:   "private-key, pk",
					Usage:  "Private key.",
					EnvVar: "WEB3_PRIVATE_KEY",
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
				contract, err := parseAddress(contract)
				if err != nil {
					return fmt.Errorf("invalid contract: %v", err)
				}
				pkStr := c.String("private-key")
				pk, err := crypto.HexToECDSA(strings.TrimPrefix(pkStr, "0x"))
				if err != nil {
					return fmt.Errorf("invalid private key %q: %v", pkStr, err)
				}
				return Pin(ctx, rpc, contract, pk, cid, dur)
			},
		},
		{
			Name:  "rate",
			Usage: "Get the current storage rate in wei per GigaByteHour.",
			Action: func(c *cli.Context) error {
				contract, err := parseAddress(contract)
				if err != nil {
					return fmt.Errorf("invalid contract: %v", err)
				}
				return Rate(ctx, rpc, contract)
			},
		},
		{
			Name:  "cost",
			Usage: "Get the current storage cost in wei for the given size and duration.",
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
				contract, err := parseAddress(contract)
				if err != nil {
					return fmt.Errorf("invalid contract: %v", err)
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
			Usage: "Get the current storage status for a CID.",
			Action: func(c *cli.Context) error {
				cid := c.Args().First()
				if cid == "" {
					return errors.New("missing CID")
				}
				return Status(ctx, api, cid)
			},
		},
		{
			Name:  "events",
			Usage: "Get Pinned events from filtered logs.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "hash",
					Usage: "Specific block hash",
				},
				cli.Int64Flag{
					Name:  "from",
					Usage: "Starting block number.",
				},
				cli.Int64Flag{
					Name:  "to",
					Usage: "Ending block number.",
				},
				cli.StringFlag{
					Name:  "cid",
					Usage: "CID to filter on.",
				},
				cli.StringFlag{
					Name:  "user",
					Usage: "User to filter on.",
				},
			},
			Action: func(c *cli.Context) error {
				contract, err := parseAddress(contract)
				if err != nil {
					return fmt.Errorf("invalid contract: %v", err)
				}
				var f gofs.EventFilter
				if c.IsSet("hash") {
					hash := c.String("hash")
					b, err := hex.DecodeString(strings.TrimPrefix(hash, "0x"))
					if err != nil {
						return fmt.Errorf("invalid hex for hash %q: %v", hash, err)
					} else if len(b) != common.HashLength {
						return fmt.Errorf("invalid hash len %d bytes: must be %d", len(b), common.HashLength)
					}
					h := common.BytesToHash(b)
					f.Hash = &h
				}
				if c.IsSet("from") {
					f.From = big.NewInt(c.Int64("from"))
				}
				if c.IsSet("to") {
					f.To = big.NewInt(c.Int64("to"))
				}
				f.CID = c.String("cid")
				if c.IsSet("user") {
					if a, err := parseAddress(c.String("user")); err != nil {
						return fmt.Errorf("invalid user: %v", err)
					} else {
						f.User = &a
					}
				}

				return Events(ctx, rpc, contract, f)
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
	}
}

func parseAddress(addr string) (common.Address, error) {
	if !common.IsHexAddress(addr) {
		return common.Address{}, fmt.Errorf("invalid hex address: %s", addr)
	}
	return common.HexToAddress(addr), nil
}

func Pin(ctx context.Context, rpcURL string, contract common.Address, pk *ecdsa.PrivateKey, ci string, dur uint64) error {
	h, r, err := gofs.Pin(ctx, rpcURL, contract, pk, ci, dur)
	if err != nil {
		return fmt.Errorf("failed to pin: %v", err)
	}
	switch r.Status {
	case types.ReceiptStatusFailed:
		return fmt.Errorf("tx %s failed", h.Hex())
	case types.ReceiptStatusSuccessful:
		fmt.Printf("Purchased %d GigaByte Hours of storage for %s.\n", dur, ci)
		fmt.Printf(`https://testnet-explorer.gochain.io/tx/%s`, h.Hex())
		return nil
	default:
		return fmt.Errorf("tx %s unrecognized receipt status: %d", h.Hex(), r.Status)
	}
}

func Add(ctx context.Context, apiURL, path string) error {
	ar, err := gofs.AddFile(ctx, apiURL, path)
	if err != nil {
		return fmt.Errorf("failed to add file %q: %v", path, err)
	}
	fmt.Println("File uploaded and pinned.")
	fmt.Println("Pinned until:", ar.Expiration)
	return nil
}

func Cost(ctx context.Context, rpcURL string, contract common.Address, size, dur uint64) error {
	_, cost, err := gofs.Cost(ctx, rpcURL, contract, size, dur)
	if err != nil {
		return err
	}

	fmt.Println(costStr(size, dur, cost))

	return nil
}

func costStr(size uint64, dur uint64, cost *big.Int) string {
	return fmt.Sprintf("%d GBs for %s: %s GO", size, time.Duration(dur)*time.Hour, web3.WeiAsBase(cost))
}

func Rate(ctx context.Context, rpcURL string, contract common.Address) error {
	rate, err := gofs.Rate(ctx, rpcURL, contract)
	if err != nil {
		return err
	}
	//TODO friendlier units?
	fmt.Printf("Current storage rate: %d wei per GigaByteHour.\n\n", rate)

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

func Status(ctx context.Context, apiURL, ci string) error {
	st, err := gofs.Status(ctx, apiURL, ci)
	if err != nil {
		return err
	}
	if st.Expiration == (time.Time{}) {
		fmt.Println("Never been pinned.")
		return nil
	}
	if until := time.Until(st.Expiration); until > 0 {
		fmt.Printf("Expires in %s at %s.\n", until, st.Expiration)
	} else {
		fmt.Printf("Expired %s ago at %s.\n", -until, st.Expiration)
	}

	return nil
}

func Events(ctx context.Context, rpcURL string, contract common.Address, f gofs.EventFilter) error {
	events, err := gofs.Events(ctx, rpcURL, contract, f)
	if err != nil {
		return err
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "Block\tTx\tLog\tRemoved\tCID\tGBH\t")
	for _, e := range events {
		ci, err := cid.Parse(e.Cid)
		if err != nil {
			return fmt.Errorf("invalid cid %s: %v", hexutil.Encode(e.Cid), err)
		}
		fmt.Fprintf(w,
			"%d\t%d\t%d\t%t\t%s\t%d\t\n",
			e.Raw.BlockNumber, //TODO why are all these idxs 0?
			e.Raw.TxIndex,
			e.Raw.Index,
			e.Raw.Removed,
			ci.String(),
			e.Gbh,
		)
	}
	w.Flush()

	return nil
}
