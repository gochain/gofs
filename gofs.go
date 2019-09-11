package gofs

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/gochain/gochain/v3"
	"github.com/gochain/gochain/v3/accounts/abi"
	"github.com/gochain/gochain/v3/accounts/abi/bind"
	"github.com/gochain/gochain/v3/common"
	"github.com/gochain/gochain/v3/common/hexutil"
	"github.com/gochain/gochain/v3/core/types"
	"github.com/gochain/gochain/v3/crypto"
	"github.com/gochain/gochain/v3/goclient"
	cid "github.com/ipfs/go-cid"
)

var (
	pinnerABI abi.ABI
	pinMethod abi.Method
)

func init() {
	parsed, err := abi.JSON(strings.NewReader(GOFSABI))
	if err != nil {
		panic(fmt.Sprintf("failed to parse generated abi: %v", err))
	}
	pinnerABI = parsed
	pinMethod = pinnerABI.Methods["pin"]
}

// Rate returns the storage cost rate in attoGO per byte-hour.
func Rate(ctx context.Context, rpcURL string, contract common.Address) (*big.Int, error) {
	gc, err := goclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	p, err := NewIGOFS(contract, gc)
	if err != nil {
		return nil, err
	}
	return p.Rate(&bind.CallOpts{Context: ctx})
}

// Cost calculates the storage cost at the current rate.
func Cost(ctx context.Context, rpcURL string, contract common.Address, bytes, hrs int64) (rate *big.Int, cost *big.Int, err error) {
	rate, err = Rate(ctx, rpcURL, contract)
	if err != nil {
		return
	}
	bh := new(big.Int).Mul(big.NewInt(bytes), big.NewInt(hrs))
	cost = bh.Mul(bh, rate)
	return
}

func Status(ctx context.Context, apiURL, ci string) (StatusResponse, error) {
	cid, err := cid.Decode(ci)
	if err != nil {
		return StatusResponse{}, err
	}
	if cid.Version() == 0 {
		return StatusResponse{}, errors.New("version 0 CID not supported")
	}
	return NewClient(apiURL).Status(ctx, cid)
}

func AddFile(ctx context.Context, apiURL, path string) (AddResponse, error) {
	f, err := os.Open(path)
	if err != nil {
		return AddResponse{}, fmt.Errorf("failed to open file %q: %v", path, err)
	}
	defer f.Close()
	return NewClient(apiURL).Add(ctx, f)
}

func Pin(ctx context.Context, rpcURL string, contract common.Address, pk *ecdsa.PrivateKey, ci string, bh *big.Int) (*types.Receipt, error) {
	cid, err := cid.Parse(ci)
	if err != nil {
		return nil, fmt.Errorf("invalid cid %q: %v", ci, err)
	}
	if cid.Version() == 0 {
		return nil, errors.New("version 0 CID not supported")
	}
	gc, err := goclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	p, err := NewIGOFS(contract, gc)
	if err != nil {
		return nil, err
	}
	rate, err := p.Rate(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	cost := new(big.Int).Mul(rate, bh)
	opts := &bind.TransactOpts{
		Context: ctx,
		From:    crypto.PubkeyToAddress(pk.PublicKey),
		Signer: func(s types.Signer, _ common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(tx, s, pk)
		},
		GasLimit: 100000,
		Value:    cost,
	}
	tx, err := p.Pin(opts, cid.Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to pin %q: %v", cid, err)
	}
	r, err := WaitForReceipt(ctx, gc, tx.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to get receipt for tx %s: %v", tx.Hash().Hex(), err)
	}
	return r, nil
}

// Wallet returns the deposit wallet address for a CID, or an empty address if none exists.
func Wallet(ctx context.Context, rpcURL string, contract common.Address, ci string) (common.Address, error) {
	cid, err := cid.Parse(ci)
	if err != nil {
		return common.Address{}, fmt.Errorf("invalid cid %q: %v", ci, err)
	}
	if cid.Version() == 0 {
		return common.Address{}, errors.New("version 0 CID not supported")
	}
	gc, err := goclient.Dial(rpcURL)
	if err != nil {
		return common.Address{}, err
	}
	p, err := NewIGOFS(contract, gc)
	if err != nil {
		return common.Address{}, err
	}
	return p.Wallet(&bind.CallOpts{Context: ctx}, cid.Bytes())
}

// NewWallet creates a new deposit wallet for a CID.
func NewWallet(ctx context.Context, rpcURL string, contract common.Address, pk *ecdsa.PrivateKey, ci string) (*types.Receipt, error) {
	cid, err := cid.Parse(ci)
	if err != nil {
		return nil, fmt.Errorf("invalid cid %q: %v", ci, err)
	}
	if cid.Version() == 0 {
		return nil, errors.New("version 0 CID not supported")
	}
	gc, err := goclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	p, err := NewIGOFS(contract, gc)
	if err != nil {
		return nil, err
	}
	tx, err := p.NewWallet(&bind.TransactOpts{
		Context: ctx,
		From:    crypto.PubkeyToAddress(pk.PublicKey),
		Signer: func(s types.Signer, _ common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(tx, s, pk)
		},
	}, cid.Bytes())
	r, err := WaitForReceipt(ctx, gc, tx.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to get receipt for tx %s: %v", tx.Hash().Hex(), err)
	}
	return r, err
}

// WaitForReceipt polls for a transaction receipt until it is available, or ctx is cancelled.
func WaitForReceipt(ctx context.Context, client *goclient.Client, hash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(ctx, hash)
		if err == nil && receipt != nil {
			return receipt, nil
		} else if err != nil && err != gochain.NotFound {
			return nil, err
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}
}

type EventFilter struct {
	// Hash for single block to match. From and To must be nil.
	Hash *common.Hash

	// Block number range.
	From *big.Int
	To   *big.Int

	Users []common.Address
	CIDs  []cid.Cid
}

type PinInputs struct {
	CID []byte `abi:"cid"`
}

// UnpackPinInputs returns arguments from a pin call.
// Handles full tx data, or plain inputs.
func UnpackPinInputs(data []byte) (pi PinInputs, err error) {
	data = bytes.TrimPrefix(data, pinMethod.Id())
	err = pinMethod.Inputs.Unpack(&pi, data)
	return
}

type Receipt struct {
	User    common.Address
	CID     cid.Cid
	BH      *big.Int
	Tx      *types.Transaction
	BlNum   uint64
	TxNum   uint
	LogNum  uint
	Removed bool
}

// Receipts returns Pinned events for the given EventFilter.
func Receipts(ctx context.Context, rpcURL string, contract common.Address, filter EventFilter) ([]Receipt, error) {
	gc, err := goclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	var q gochain.FilterQuery
	q.Addresses = []common.Address{contract}
	if filter.Hash != nil {
		if filter.From != nil || filter.To != nil {
			return nil, fmt.Errorf("cannot filter on both hash and from/to")
		}
		q.BlockHash = filter.Hash
	} else {
		q.FromBlock = filter.From
		q.ToBlock = filter.To
	}
	var userTopic []common.Hash
	if len(filter.Users) > 0 {
		userTopic = make([]common.Hash, len(filter.Users))
		for i := range filter.Users {
			userTopic[i] = filter.Users[i].Hash()
		}

	}
	var cidTopic []common.Hash
	if len(filter.CIDs) > 0 {
		cidTopic = make([]common.Hash, len(filter.CIDs))
		for i, ci := range filter.CIDs {
			if ci.Version() == 0 {
				return nil, errors.New("version 0 CID not supported")
			}
			cidTopic[i] = crypto.Keccak256Hash(ci.Bytes())
		}
	}
	q.Topics = [][]common.Hash{{pinnerABI.Events["Pinned"].Id()}, userTopic, cidTopic}
	logs, err := gc.FilterLogs(ctx, q)
	if err != nil {
		return nil, err
	}

	bc := bind.NewBoundContract(contract, pinnerABI, nil, nil, nil)
	igofs, err := NewIGOFS(contract, gc)
	var receipts []Receipt
	for _, l := range logs {
		var event IGOFSPinned
		if err := bc.UnpackLog(&event, "Pinned", l); err != nil {
			return nil, err
		}
		tx, _, err := gc.TransactionByHash(ctx, l.TxHash)
		if err != nil {
			return nil, fmt.Errorf("failed to get tx %s: %v", l.TxHash, err)
		}
		from, err := gc.TransactionSender(ctx, tx, l.BlockHash, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to look up tx sender %s: %v", l.TxHash, err)
		}
		cidBytes, err := igofs.CidByHash(&bind.CallOpts{Context: ctx}, event.Cid)
		if err != nil {
			return nil, fmt.Errorf("failed to get CID for hash %s: %v", event.Cid.Hex(), err)
		}
		ci, err := cid.Parse(cidBytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse CID %s: %v", hexutil.Encode(cidBytes), err)
		}
		receipts = append(receipts, Receipt{
			User:    from,
			CID:     ci,
			BH:      event.Bh,
			Tx:      tx,
			BlNum:   l.BlockNumber,
			TxNum:   l.TxIndex,
			LogNum:  l.Index,
			Removed: l.Removed,
		})
	}

	return receipts, nil
}
