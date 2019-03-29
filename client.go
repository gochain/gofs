package gofs

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gochain-io/gochain/v3/common"
	cid "github.com/ipfs/go-cid"
)

const (
	APIURL        = "https://gofs.io/api/v0/" //TODO
	MainnetRPCURL = "https://rpc.gochain.io"
)

var MainnetAddress = common.HexToAddress("0x1234") //TODO

type API interface {
	// Add puts new data and temporarily pins.
	Add(context.Context, io.Reader) (AddResponse, error)
	// Status returns the status of a CID.
	Status(ctx context.Context, ci cid.Cid) (StatusResponse, error)
}

type AddResponse struct {
	CID        string `json:"cid"`
	Expiration int64  `json:"expiration"` // Unix timestamp.
}

type StatusResponse struct {
	Pinned     bool  `json:"pinned"`
	Expiration int64 `json:"expiration,omitempty"` // Unix timestamp.
}

func NewClient(url string) API {
	return &client{
		url: url,
	}
}

type client struct {
	url string
}

func (c *client) Add(ctx context.Context, r io.Reader) (AddResponse, error) {
	req, err := http.NewRequest("PUT", c.url+"add", r)
	if err != nil {
		return AddResponse{}, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return AddResponse{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return AddResponse{}, fmt.Errorf("http error: %d - %s", resp.StatusCode, resp.Status)
	}
	var ar AddResponse
	if err := json.NewDecoder(resp.Body).Decode(&ar); err != nil {
		return AddResponse{}, fmt.Errorf("failed to unmarshal json response: %v", err)
	}
	return ar, nil
}

func (c *client) Status(ctx context.Context, ci cid.Cid) (StatusResponse, error) {
	resp, err := http.Get(c.url + "status")
	if err != nil {
		return StatusResponse{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return StatusResponse{}, fmt.Errorf("http error: %d - %s", resp.StatusCode, resp.Status)
	}
	var sr StatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&sr); err != nil {
		return StatusResponse{}, fmt.Errorf("failed to unmarshal json response: %v", err)
	}
	return sr, nil
}