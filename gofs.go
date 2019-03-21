package gofs

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/gochain-io/gochain/v3/accounts/abi/bind"
	"github.com/gochain-io/gochain/v3/common"
	"github.com/gochain-io/gochain/v3/goclient"
	"github.com/gochain-io/web3"
)

func Pin(ctx context.Context, cid string, contract common.Address) error {
	//TODO validate CID
	gc, err := goclient.Dial(MainnetRPCURL)
	if err != nil {
		return err
	}
	p, err := NewPinner(contract, gc)
	if err != nil {
		return err
	}
	opts := &bind.TransactOpts{
		Context: ctx,
		//TODO set value
	}
	_, err = p.Pin(opts, cid)
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
	ar, err := NewClient(url).Add(ctx, f)
	if err != nil {
		return fmt.Errorf("failed to add file %q: %v", path, err)
	}
	fmt.Println("File uploaded and pinned.")
	fmt.Println("Pinned until:", time.Unix(ar.Expiration, 0))
	return nil
}

func Rate(ctx context.Context, rpcurl string, contract common.Address) error {
	gc, err := goclient.Dial(rpcurl)
	if err != nil {
		return err
	}
	p, err := NewPinner(contract, gc)
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
		gbs int64
		hrs int64
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
		gbh := big.NewInt(vals.gbs * vals.hrs)
		cost := new(big.Int).Mul(gbh, rate)

		fmt.Printf("%d GBs for %s: %s", vals.gbs, time.Duration(vals.hrs)*time.Hour, web3.WeiAsBase(cost))
	}
	return nil
}

func Status(ctx context.Context, url, cid string) error {
	//TODO validate CID
	st, err := NewClient(url).Status(ctx, cid)
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
