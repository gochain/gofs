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

	"github.com/alecthomas/units"
	"github.com/gochain-io/gochain/v3/common"
	"github.com/gochain-io/gochain/v3/core/types"
	"github.com/gochain-io/gochain/v3/crypto"
	"github.com/gochain-io/gofs"
	"github.com/gochain-io/web3"
	cid "github.com/ipfs/go-cid"
	"github.com/urfave/cli"
)

var version string

func init() {
	if version == "" {
		version = "unknown"
	}
}

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
	app.Version = version
	app.Usage = "GoChain filesystem cli tool"
	var api, rpc, contract string
	//TODO var recursive bool
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
					Name:  "bh",
					Usage: "Storage to purchase in ByteHours.",
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
				bh := c.Uint64("bh")
				if bh == 0 {
					return fmt.Errorf("bh missing or invalid")
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
				return Pin(ctx, rpc, contract, pk, cid, bh)
			},
		},
		{
			Name:  "wallet",
			Usage: "Get the deposit wallet for the CID.",
			Flags: []cli.Flag{
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
				contract, err := parseAddress(contract)
				if err != nil {
					return fmt.Errorf("invalid contract: %v", err)
				}
				pkStr := c.String("private-key")
				pk, err := crypto.HexToECDSA(strings.TrimPrefix(pkStr, "0x"))
				if err != nil {
					return fmt.Errorf("invalid private key %q: %v", pkStr, err)
				}
				return Wallet(ctx, rpc, contract, pk, cid)
			},
		},
		{
			Name:  "rate",
			Usage: "Get the current storage rate in attoGO per ByteHour.",
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
			Usage: "Get the current storage cost in attoGO for the given size and duration.",
			Flags: []cli.Flag{
				cli.Uint64Flag{
					Name:  "duration, d",
					Usage: "Storage duration in hours.",
					Value: 1,
				},
				cli.StringFlag{
					Name:  "size, s",
					Usage: "Storage size.",
					Value: units.GB.String(),
				},
			},
			Action: func(c *cli.Context) error {
				//TODO accept go duration?
				dur := c.Int64("duration")
				if dur == 0 {
					return fmt.Errorf("duration missing or invalid")
				}
				sizeStr := c.String("size")
				bytes, err := units.ParseMetricBytes(sizeStr)
				if err != nil {
					return fmt.Errorf("invalid size %q: %v", sizeStr, err)
				}
				if bytes == 0 {
					return fmt.Errorf("size must be greater than 0")
				}
				contract, err := parseAddress(contract)
				if err != nil {
					return fmt.Errorf("invalid contract: %v", err)
				}
				return Cost(ctx, rpc, contract, int64(bytes), dur)
			},
		},
		{
			Name:  "add",
			Usage: "Add and pin a file.",
			Flags: []cli.Flag{
				//TODO
				//cli.BoolFlag{
				//	Name:        "recursive, r",
				//	Usage:       "Add directory recursively",
				//	Destination: &recursive,
				//},
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
			Name:  "receipts",
			Usage: "Query for receipts.",
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
					Name:  "cids",
					Usage: "Comma separated CIDs to filter on.",
				},
				cli.StringFlag{
					Name:  "users",
					Usage: "Comma separated users to filter on.",
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
				if c.IsSet("cids") {
					for _, s := range c.StringSlice("cids") {
						ci, err := cid.Parse(s)
						if err != nil {
							return fmt.Errorf("invalid cid %s: %v", s, err)
						}
						f.CIDs = append(f.CIDs, ci)
					}

				}
				if c.IsSet("users") {
					for _, s := range c.StringSlice("user") {
						a, err := parseAddress(s)
						if err != nil {
							return fmt.Errorf("invalid user: %v", err)
						}
						f.Users = append(f.Users, a)
					}
				}

				return Receipts(ctx, rpc, contract, f)
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

func Pin(ctx context.Context, rpcURL string, contract common.Address, pk *ecdsa.PrivateKey, ci string, bh uint64) error {
	r, err := gofs.Pin(ctx, rpcURL, contract, pk, ci, bh)
	if err != nil {
		return fmt.Errorf("failed to pin: %v", err)
	}
	switch r.Status {
	case types.ReceiptStatusFailed:
		return fmt.Errorf("tx %s failed", r.TxHash.Hex())
	case types.ReceiptStatusSuccessful:
		fmt.Printf("Purchased %d ByteHours of storage for %s.\n", bh, ci)
		fmt.Println("Tx:", r.TxHash.Hex())
		return nil
	default:
		return fmt.Errorf("tx %s unrecognized receipt status: %d", r.TxHash.Hex(), r.Status)
	}
}

func Wallet(ctx context.Context, rpcURL string, contract common.Address, pk *ecdsa.PrivateKey, ci string) error {
	w, err := gofs.Wallet(ctx, rpcURL, contract, ci)
	if err != nil {
		return fmt.Errorf("failed to get deposit wallet: %v", err)
	}

	if w != (common.Address{}) {
		fmt.Println("Deposit wallet:", w.Hex())
		return nil
	}
	fmt.Println("No deposit wallet exists for this CID. Create one? Y/N")
	var s string
	if _, err := fmt.Scan(&s); err != nil {
		return fmt.Errorf("failed to read input: %v", err)
	}
	s = strings.TrimSpace(strings.ToLower(s))
	switch s {
	case "y", "yes":

	default:
		return nil
	}

	r, err := gofs.NewWallet(ctx, rpcURL, contract, pk, ci)
	if err != nil {
		return fmt.Errorf("failed to deploy deposit wallet: %v", err)
	}
	switch r.Status {
	case types.ReceiptStatusFailed:
		return fmt.Errorf("tx %s failed", r.TxHash.Hex())
	case types.ReceiptStatusSuccessful:
		fmt.Println("Created new wallet - Tx:", r.TxHash.Hex())
		w, err := gofs.Wallet(ctx, rpcURL, contract, ci)
		if err != nil {
			return fmt.Errorf("failed to get deposit wallet: %v", err)
		}
		if w == (common.Address{}) {
			return errors.New("address not found after successful creation")
		}
		fmt.Println("Deposit wallet:", w.Hex())
		return nil
	default:
		return fmt.Errorf("tx %s unrecognized receipt status: %d", r.TxHash.Hex(), r.Status)
	}
}

func Add(ctx context.Context, apiURL, path string) error {
	ar, err := gofs.AddFile(ctx, apiURL, path)
	if err != nil {
		return fmt.Errorf("failed to add file %q: %v", path, err)
	}
	fmt.Println("File uploaded and pinned.")
	fmt.Println("CID:", ar.CID)
	fmt.Println("Pinned until:", time.Unix(ar.Expiration, 0))
	fmt.Println("File size:", units.MetricBytes(ar.Size))
	return nil
}

func Cost(ctx context.Context, rpcURL string, contract common.Address, bytes, hrs int64) error {
	_, cost, err := gofs.Cost(ctx, rpcURL, contract, bytes, hrs)
	if err != nil {
		return err
	}

	fmt.Println(costStr(bytes, hrs, cost))

	return nil
}

func costStr(bytes int64, hrs int64, cost *big.Int) string {
	return fmt.Sprintf("%s for %s: %s GO", units.MetricBytes(bytes), prettyHours(hrs), web3.WeiAsBase(cost))
}

func prettyHours(hours int64) string {
	days := hours / 24
	hours -= days * 24

	weeks := days / 7
	days -= weeks * 7

	years := weeks / 52
	weeks -= years * 52

	var s strings.Builder
	if years > 0 {
		fmt.Fprintf(&s, "%dy", years)
	}
	if weeks > 0 {
		fmt.Fprintf(&s, "%dw", weeks)
	}
	if days > 0 {
		fmt.Fprintf(&s, "%dd", days)
	}
	if hours > 0 {
		fmt.Fprintf(&s, "%dh", hours)
	}
	return s.String()
}

func Rate(ctx context.Context, rpcURL string, contract common.Address) error {
	rate, err := gofs.Rate(ctx, rpcURL, contract)
	if err != nil {
		return err
	}
	fmt.Printf("Current storage rate: %d attoGO per ByteHour.\n\n", rate)

	fmt.Println("Cost:")
	for _, vals := range []struct {
		bytes units.MetricBytes
		hrs   int64
	}{
		{bytes: units.GB, hrs: 1},
		{bytes: 10 * units.GB, hrs: 1},
		{bytes: units.TB, hrs: 1},
		{bytes: units.GB, hrs: 24},
		{bytes: units.GB, hrs: 24 * 7},
		{bytes: units.GB, hrs: 24 * 7 * 52},
		{bytes: 10 * units.GB, hrs: 24 * 7 * 52},
		{bytes: units.Terabyte, hrs: 24 * 7 * 52},
	} {
		bh := new(big.Int).Mul(big.NewInt(int64(vals.bytes)), big.NewInt(vals.hrs))
		cost := bh.Mul(bh, rate)

		fmt.Println("\t", costStr(int64(vals.bytes), vals.hrs, cost))
	}
	return nil
}

func Status(ctx context.Context, apiURL, ci string) error {
	st, err := gofs.Status(ctx, apiURL, ci)
	if err != nil {
		return err
	}
	if st.Expiration == 0 {
		fmt.Println("Never been pinned.")
		return nil
	}
	fmt.Println("File size:", units.MetricBytes(st.Size))
	exp := time.Unix(st.Expiration, 0)
	if until := time.Until(exp).Round(time.Second); until > 0 {
		fmt.Printf("Expires in %s at %s.\n", until, exp)
	} else {
		fmt.Printf("Expired %s ago at %s.\n", -until, exp)
	}
	return nil
}

func Receipts(ctx context.Context, rpcURL string, contract common.Address, f gofs.EventFilter) error {
	receipts, err := gofs.Receipts(ctx, rpcURL, contract, f)
	if err != nil {
		return err
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "Block\tTx\tLog\tRemoved\tCID\tBH\tUser\t")
	for _, r := range receipts {
		fmt.Fprintf(w,
			"%d\t%d\t%d\t%t\t%s\t%s\t%s\t\n",
			r.BlNum,
			r.TxNum,
			r.LogNum,
			r.Removed,
			r.CID.String(),
			r.BH,
			r.User.Hex(),
		)
	}
	w.Flush()

	return nil
}
