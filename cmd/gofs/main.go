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
	"strconv"
	"strings"
	"syscall"
	"text/tabwriter"
	"time"

	"github.com/alecthomas/units"
	"github.com/gochain/gochain/v3/common"
	"github.com/gochain/gochain/v3/core/types"
	"github.com/gochain/gochain/v3/crypto"
	"github.com/gochain/gofs"
	"github.com/gochain/web3"
	cid "github.com/ipfs/go-cid"
	shell "github.com/ipfs/go-ipfs-api"
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
			Usage:       "Hex contract address.",
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
			Usage: "Pin a new CID or extend the expiry of an existing file.",
			UsageText: `Pin a new CID in GOFS or extend the expiry of an existing file by purchasing storage at the current rate. Requires one of --bh, --extend, or --until.
	Examples: 
		gofs pin --until 2020-10-01 <cid>
		gofs pin --extend 1y6m --size 1.4MB <cid>
		gofs pin --bh 30000000 <cid>`,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "private-key, pk",
					Usage:  "Private key.",
					EnvVar: "WEB3_PRIVATE_KEY",
				},
				cli.StringFlag{
					Name:  "extend, e", //TODO "duration, d"?
					Usage: "Duration to extend expiration. Format: 1y200d Not compatible with --until or --bh.",
				},
				cli.StringFlag{
					Name:  "until, u",
					Usage: "Date to extend expiration until. Format: yyyy-MM-dd. Not compatible with --extend or --bh.",
				},
				cli.StringFlag{
					Name:  "size, s",
					Usage: "File size in bytes. Example: 1.4KB. Optional. If omitted, file size will be fetched from GOFS or IPFS.",
				},
				cli.Uint64Flag{
					Name:  "bh",
					Usage: "Storage to purchase in byte-hours. Use with caution. Prefer --extend or --until. Not compatible with --extend or --until.",
				},
				// TODO optional ipfs api url?
			},
			Action: func(c *cli.Context) error {
				cid := c.Args().First()
				if cid == "" {
					return errors.New("missing CID")
				}
				bhFn, err := byteHourActionHelper(c, api, cid)
				if err != nil {
					return err
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
				bh, err := bhFn(ctx)
				if err != nil {
					return err
				}
				return Pin(ctx, rpc, contract, pk, cid, bh)
			},
		},
		{
			Name:  "wallet",
			Usage: "Get the deposit wallet for the CID or create one if none exists.",
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
			Usage: "Get the current storage rate in attoGO per byte-hour.",
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
			Usage: "Get the current storage cost for the given size and duration.",
			UsageText: `Requires one of --bh, --extend, or --until. File size is required for --extend and --until, and can be passed via --size, or fetched based on the CID argument.
	Examples: 
		gofs cost --until 2020-10-01 <cid>
		gofs cost --extend 1y6m --size 1.4MB
		gofs cost --bh 30000000`,
			Flags: []cli.Flag{ // keep aligned with pin
				cli.StringFlag{
					Name:  "extend, e",
					Usage: "Duration to extend expiration. Format: 1y200d Not compatible with --until or --bh.",
				},
				cli.StringFlag{
					Name:  "until, u",
					Usage: "Date to extend expiration until. Format: yyyy-MM-dd. Not compatible with --extend or --bh.",
				},
				cli.StringFlag{
					Name:  "size, s",
					Usage: "File size in bytes. Example: 1.4KB. Required for --extend and --until without CID.",
				},
				cli.Uint64Flag{
					Name:  "bh",
					Usage: "Storage to purchase in byte-hours. Use with caution. Prefer --extend or --until. Not compatible with --extend or --until.",
				},
			},
			Action: func(c *cli.Context) error {
				cid := c.Args().First()
				bhFn, err := byteHourActionHelper(c, api, cid)
				if err != nil {
					return err
				}
				contract, err := parseAddress(contract)
				if err != nil {
					return fmt.Errorf("invalid contract: %v", err)
				}
				bh, err := bhFn(ctx)
				if err != nil {
					return err
				}
				return Cost(ctx, rpc, contract, bh)
			},
		},
		{
			Name:      "add",
			Usage:     "Add and pin a file.",
			UsageText: "Add and pin a file by uploading to GOFS.",
			Flags:     []cli.Flag{
				// TODO recursive
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
				cli.StringSliceFlag{
					Name:  "cids",
					Usage: "Comma separated CIDs to filter on.",
				},
				cli.StringSliceFlag{
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

// byteHourFn returns bytes-hours
type byteHourFn func(context.Context) (*big.Int, error)

// byteHourActionHelper resolves the set of flags [until, extend, bh, size] in to a byteHourFn.
// cid is optional.
func byteHourActionHelper(c *cli.Context, api, cid string) (byteHourFn, error) {
	if c.IsSet("until") {
		if c.IsSet("extend") || c.IsSet("bh") {
			return nil, errors.New("--until is not compatible with --extend or --bh")
		}
		u := c.String("until")
		y, m, d, err := parseDate(u)
		if err != nil {
			return nil, fmt.Errorf("failed to parse --until %q: %v", u, err)
		}
		date := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
		var size int64
		if c.IsSet("size") {
			size, err = units.ParseStrictBytes(c.String("size"))
			if err != nil {
				return nil, fmt.Errorf("failed to parse size: %v", err)
			}
			if size <= 0 {
				return nil, fmt.Errorf("invalid size")
			}
		} else if cid == "" {
			return nil, errors.New("missing file size: --size flag or cid required")
		}
		return byteHoursFromDate(api, cid, date, size), nil
	} else if c.IsSet("extend") {
		if c.IsSet("until") || c.IsSet("bh") {
			return nil, errors.New("--extend is not compatible with --until or --bh")
		}
		e := c.String("extend")
		var y, d int
		y, d, err := parseDuration(e)
		if err != nil {
			return nil, fmt.Errorf("invalid duration %q: %v", e, err)
		}
		now := time.Now()
		then := now.AddDate(y, 0, d)
		dur := then.Sub(now)
		if !c.IsSet("size") {
			if cid == "" {
				return nil, errors.New("missing file size: --size flag or cid required")
			}
			return byteHoursFromDuration(api, cid, dur), nil
		}
		size, err := units.ParseStrictBytes(c.String("size"))
		if err != nil {
			return nil, fmt.Errorf("failed to parse size: %v", err)
		}
		if size <= 0 {
			return nil, fmt.Errorf("invalid size")
		}
		hours := int64(dur / time.Hour)
		bh := new(big.Int).Mul(big.NewInt(size), big.NewInt(hours))
		return func(context.Context) (*big.Int, error) { return bh, nil }, nil
	} else if c.IsSet("bh") {
		if c.IsSet("size") {
			return nil, errors.New("--bh and --size are incompatible")
		}
		return func(ctx context.Context) (*big.Int, error) {
			bh, ok := new(big.Int).SetString(c.String("bh"), 10)
			if !ok {
				return nil, fmt.Errorf("bh invalid")
			}
			if bh.Cmp(big.NewInt(1)) == -1 {
				return nil, fmt.Errorf("bh invalid")
			}
			return bh, nil
		}, nil
	}
	return nil, errors.New("one of --extend, --until, or --bh is required")
}

func parseAddress(addr string) (common.Address, error) {
	if !common.IsHexAddress(addr) {
		return common.Address{}, fmt.Errorf("invalid hex address: %s", addr)
	}
	return common.HexToAddress(addr), nil
}

// parseDate returns year, month, and day, from a yyyy-MM-dd formatted date.
// yyyy-MM and yyyy are also accepted.
func parseDate(date string) (int, time.Month, int, error) {
	parts := strings.Split(date, "-")
	var (
		y int
		m = time.January
		d = 1
	)
	l := len(parts)
	switch {
	case l > 3, l == 0:
		return 0, 0, 0, fmt.Errorf("invalid date %q: must match yyyy-MM-dd", date)
	case l == 3:
		d64, err := strconv.ParseInt(parts[2], 10, 0)
		if err != nil {
			return 0, 0, 0, fmt.Errorf("invalid day %q: %v", parts[2], err)
		}
		d = int(d64)
		fallthrough
	case l == 2:
		m64, err := strconv.ParseInt(parts[1], 10, 0)
		if err != nil {
			return 0, 0, 0, fmt.Errorf("invalid month %q: %v", parts[1], err)
		}
		m = time.Month(m64)
		fallthrough
	case l == 1:
		y64, err := strconv.ParseInt(parts[0], 10, 0)
		if err != nil {
			return 0, 0, 0, fmt.Errorf("invalid year %q: %v", parts[0], err)
		}
		y = int(y64)
	default:
		panic("unreachable")
	}
	return y, m, d, nil
}

func Pin(ctx context.Context, rpcURL string, contract common.Address, pk *ecdsa.PrivateKey, ci string, bh *big.Int) error {
	r, err := gofs.Pin(ctx, rpcURL, contract, pk, ci, bh)
	if err != nil {
		return fmt.Errorf("failed to pin: %v", err)
	}
	switch r.Status {
	case types.ReceiptStatusFailed:
		return fmt.Errorf("tx %s failed", r.TxHash.Hex())
	case types.ReceiptStatusSuccessful:
		fmt.Printf("Purchased %d byte-hours of storage for %s.\n", bh, ci)
		fmt.Println("Tx:", r.TxHash.Hex())
		return nil
	default:
		return fmt.Errorf("tx %s unrecognized receipt status: %d", r.TxHash.Hex(), r.Status)
	}
}

// byteHoursFromDate returns a func to compute byte-hours from a date,
// based on any existing expiration and on file size fetched from GOFS or IPFS, when size is 0.
func byteHoursFromDate(api, cid string, date time.Time, size int64) func(ctx context.Context) (*big.Int, error) {
	return func(ctx context.Context) (*big.Int, error) {
		st, err := gofs.Status(ctx, api, cid)
		if err != nil {
			return nil, err
		}
		var dur time.Duration
		if st.Expiration != 0 {
			exp := time.Unix(st.Expiration, 0)
			if exp.After(date) {
				// Already pinned through that date.
				return nil, nil
			}
			dur = date.Sub(exp)
			if size == 0 {
				size = st.Size
			} else if size != st.Size { // sanity check user input
				return nil, fmt.Errorf("size mismatch: GOFS %d, input %d", st.Size, size)
			}
		} else {
			dur = time.Until(date)
			os, err := ipfsObjectStats(ctx, cid)
			if err != nil {
				return nil, err
			}
			if os == nil {
				return nil, fmt.Errorf("cid not found on IPFS: %s", cid)
			}
			if size == 0 {
				size = int64(os.CumulativeSize)
			} else if size != int64(os.CumulativeSize) { // sanity check user input
				return nil, fmt.Errorf("size mismatch: IPFS %d, input %d", os.CumulativeSize, size)
			}
		}
		hours := int64(dur / time.Hour)
		return new(big.Int).Mul(big.NewInt(size), big.NewInt(hours)), nil
	}
}

// byteHoursFromDuration returns a func to compute byte-hours from a duration,
// based on file size fetched from GOFS or IPFS.
func byteHoursFromDuration(api, cid string, dur time.Duration) func(context.Context) (*big.Int, error) {
	return func(ctx context.Context) (*big.Int, error) {
		st, err := gofs.Status(ctx, api, cid)
		if err != nil {
			return nil, err
		}
		var size int64
		if st.Expiration != 0 {
			size = st.Size
		} else {
			os, err := ipfsObjectStats(ctx, cid)
			if err != nil {
				return nil, err
			}
			if os == nil {
				return nil, fmt.Errorf("cid not found on IPFS: %s", cid)
			}
			size = int64(os.CumulativeSize)
		}
		hours := int64(dur / time.Hour)
		return new(big.Int).Mul(big.NewInt(size), big.NewInt(hours)), nil
	}
}

var ipfsURLs = []string{"https://ipfs.infura.io:5001"}

// ipfsObjectStats fetches objects stats concurrently from multiple sources.
func ipfsObjectStats(ctx context.Context, ci string) (*shell.ObjectStats, error) {
	osCh := make(chan *shell.ObjectStats, len(ipfsURLs))
	errCh := make(chan error, len(ipfsURLs))
	for _, url := range ipfsURLs {
		go func(url string) {
			os, err := shell.NewShell(url).ObjectStat(ci)
			if err != nil {
				errCh <- fmt.Errorf("%s: %v", url, err)
				return
			}
			if os == nil {
				errCh <- fmt.Errorf("%s: not found", url)
				return
			}
			osCh <- os
		}(url)
	}
	var errs []error
	for len(errs) < len(ipfsURLs) {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case os := <-osCh:
			return os, nil
		case err := <-errCh:
			errs = append(errs, err)
		}
	}
	return nil, fmt.Errorf("all ipfs object stats calls failed for %q: %v", ci, errs)
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
	fmt.Println("File size:", prettySize(ar.Size))
	return nil
}

var siUnits = [...]string{"K", "M", "G", "T", "P", "E"}
var metricsUnits = [...]units.MetricBytes{units.KB, units.MB, units.GB, units.TB, units.PB, units.EB}

// prettySize returns the size in bytes as the largest metric unit limited to 2 decimals.
func prettySize(b int64) string {
	mb := units.MetricBytes(b)
	if mb < units.KB {
		return fmt.Sprintf("%dB", b)
	}
	for i, l := range metricsUnits[1:] { // start with MB
		if mb >= l {
			continue
		}
		// note that i now references the previous element, since we are ranging over a slice.
		f := big.NewRat(b, int64(metricsUnits[i])).FloatString(2)
		return fmt.Sprintf("%s%sB", f, siUnits[i])
	}
	f := big.NewRat(b, int64(units.EB)).FloatString(2)
	return fmt.Sprintf("%sEB", f)
}

func Cost(ctx context.Context, rpcURL string, contract common.Address, bh *big.Int) error {
	rate, err := gofs.Rate(ctx, rpcURL, contract)
	if err != nil {
		return err
	}
	cost := new(big.Int).Mul(rate, bh)
	fmt.Printf("%s bytes-hours: %s GO\n", bh, web3.WeiAsBase(cost))
	return nil
}

// prettyHours formats a duration like 1y2d3h.
func prettyHours(hours int64) string {
	days := hours / 24
	hours -= days * 24

	years := days / 365
	days -= years * 365

	var s strings.Builder
	if years > 0 {
		fmt.Fprintf(&s, "%dy", years)
	}
	if days > 0 {
		fmt.Fprintf(&s, "%dd", days)
	}
	if hours > 0 {
		fmt.Fprintf(&s, "%dh", hours)
	}
	return s.String()
}

// parseDuration returns years and days from a 1y2d formatted duration.
// 1y and 2d are also accepted.
func parseDuration(dur string) (int, int, error) {
	var y, d int
	if len(dur) == 0 {
		return 0, 0, errors.New("empty")
	}
	last := dur[len(dur)-1]
	switch last {
	case 'd':
		yi := strings.Index(dur, "y")
		ds := dur[yi+1 : len(dur)-1]
		d64, err := strconv.ParseInt(ds, 10, 0)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid day %q: %v", ds, err)
		}
		d = int(d64)
		dur = dur[:yi+1]
		fallthrough
	case 'y':
		if dur == "" {
			break
		}
		ys := dur[:len(dur)-1]
		y64, err := strconv.ParseInt(ys, 10, 0)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid year %q: %v", ys, err)
		}
		y = int(y64)
	default:
		return 0, 0, fmt.Errorf("unrecognized character: %c", last)
	}
	return y, d, nil
}

func Rate(ctx context.Context, rpcURL string, contract common.Address) error {
	rate, err := gofs.Rate(ctx, rpcURL, contract)
	if err != nil {
		return err
	}
	fmt.Printf("Current storage rate: %d attoGO per byte-hour.\n\n", rate)

	fmt.Println("Cost:")
	for _, vals := range []struct {
		bytes units.MetricBytes
		hrs   int64
	}{
		// 30 days
		{bytes: units.KB, hrs: 24 * 30},
		{bytes: units.MB, hrs: 24 * 30},
		{bytes: units.GB, hrs: 24 * 30},
		// 1 year
		{bytes: units.KB, hrs: 24 * 365},
		{bytes: units.MB, hrs: 24 * 365},
		{bytes: units.GB, hrs: 24 * 365},
	} {
		bh := new(big.Int).Mul(big.NewInt(int64(vals.bytes)), big.NewInt(vals.hrs))
		cost := bh.Mul(bh, rate)

		fmt.Println("\t", costStr(int64(vals.bytes), vals.hrs, cost))
	}
	return nil
}

func costStr(bytes int64, hrs int64, cost *big.Int) string {
	return fmt.Sprintf("%s for %s: %s GO", prettySize(bytes), prettyHours(hrs), web3.WeiAsBase(cost))
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
	fmt.Println("File size:", prettySize(st.Size))
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
