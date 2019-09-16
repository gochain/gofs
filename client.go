package gofs

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gochain/gochain/v3/common"
	cid "github.com/ipfs/go-cid"
)

const (
	APIURL        = "https://api.gofs.io"
	MainnetRPCURL = "https://rpc.gochain.io"
)

var MainnetAddress = common.HexToAddress("0x545a6A1AE20a6091898a5cEe2D7b43A04D77d4C0")

type API interface {
	// Add puts new data and temporarily pins.
	Add(context.Context, io.ReadCloser) (AddResponse, error)
	// Status returns the status of a CID.
	Status(ctx context.Context, ci cid.Cid) (StatusResponse, error)
}
type AddResponse struct {
	CID        string `json:"cid"`
	Expiration int64  `json:"expiration"` // Unix TS
	Size       int64  `json:"size"`       // File size in bytes.
}

type StatusResponse struct {
	Expiration int64 `json:"expiration,omitempty"` // Unix TS
	Size       int64 `json:"size"`                 // File size in bytes.
}

// DefaultClient uses the default APIURL: https://api.gofs.io
var DefaultClient = NewClient(APIURL)

// NewClient returns a new client backed by url.
func NewClient(url string) API {
	return &client{
		url: url,
	}
}

type client struct {
	url string
}

func (c *client) Add(ctx context.Context, r io.ReadCloser) (AddResponse, error) {
	req, err := http.NewRequest("PUT", c.url+"/v0/add", r)
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
	resp, err := http.Get(c.url + "/v0/status/" + ci.String())
	if err != nil {
		return StatusResponse{}, err
	}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return StatusResponse{}, nil
		}
		return StatusResponse{}, fmt.Errorf("http error: %d - %s", resp.StatusCode, resp.Status)
	}
	var sr StatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&sr); err != nil {
		return StatusResponse{}, fmt.Errorf("failed to unmarshal json response: %v", err)
	}
	return sr, nil
}
