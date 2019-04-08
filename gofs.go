package gofs

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/gochain-io/gochain/v3"

	"github.com/gochain-io/gochain/v3/crypto"

	"github.com/gochain-io/gochain/v3/accounts/abi/bind"
	"github.com/gochain-io/gochain/v3/common"
	"github.com/gochain-io/gochain/v3/core/types"
	"github.com/gochain-io/gochain/v3/goclient"
	cid "github.com/ipfs/go-cid"
)

func Rate(ctx context.Context, rpcURL string, contract common.Address) (*big.Int, error) {
	gc, err := goclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	p, err := NewPinner(contract, gc)
	if err != nil {
		return nil, err
	}
	return p.Rate(&bind.CallOpts{Context: ctx})
}

// Cost calculates the storage cost of size GigaBytes for dur hours at the current rate.
func Cost(ctx context.Context, rpcURL string, contract common.Address, size, dur uint64) (rate *big.Int, cost *big.Int, err error) {
	gc, err := goclient.Dial(rpcURL)
	if err != nil {
		return
	}
	p, err := NewPinner(contract, gc)
	if err != nil {
		return
	}
	rate, err = p.Rate(&bind.CallOpts{Context: ctx})
	if err != nil {
		return
	}
	gbhs := size * dur
	cost = new(big.Int).Mul(rate, big.NewInt(int64(gbhs)))
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

func AddFile(ctx context.Context, url, path string) (AddResponse, error) {
	f, err := os.Open(path)
	if err != nil {
		return AddResponse{}, fmt.Errorf("failed to open file %q: %v", path, err)
	}
	defer f.Close()
	return NewClient(url).Add(ctx, f)
}

func Pin(ctx context.Context, url string, contract common.Address, pk *ecdsa.PrivateKey, ci string, dur uint64) (common.Hash, *types.Receipt, error) {
	cid, err := cid.Parse(ci)
	if err != nil {
		return common.Hash{}, nil, fmt.Errorf("invalid cid %q: %v", ci, err)
	}
	if cid.Version() == 0 {
		return common.Hash{}, nil, errors.New("version 0 CID not supported")
	}
	gc, err := goclient.Dial(url)
	if err != nil {
		return common.Hash{}, nil, err
	}
	p, err := NewPinner(contract, gc)
	if err != nil {
		return common.Hash{}, nil, err
	}
	rate, err := p.Rate(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Hash{}, nil, err
	}
	cost := new(big.Int).Mul(rate, big.NewInt(int64(dur)))
	opts := &bind.TransactOpts{
		Context: ctx,
		From:    crypto.PubkeyToAddress(pk.PublicKey),
		Signer: func(s types.Signer, _ common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(tx, s, pk)
		},
		GasLimit: 50000,
		Value:    cost,
	}
	tx, err := p.Pin(opts, cid.Bytes(), big.NewInt(int64(dur)))
	if err != nil {
		return common.Hash{}, nil, fmt.Errorf("failed to pin %q: %v", cid, err)
	}
	r, err := WaitForReceipt(ctx, gc, tx.Hash())
	if err != nil {
		return common.Hash{}, nil, fmt.Errorf("failed to get receipt for tx %s: %v", tx.Hash().Hex(), err)
	}
	return tx.Hash(), r, nil
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
