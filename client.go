package gofs

import (
	"context"
	"io"
)

type Client interface {
	Pin(ctx context.Context, cid string) (PinResponse, error)
	Add(context.Context, io.Reader) (AddResponse, error)
}

type PinResponse struct {
	Pinned     bool  // True if pinned. False if in progress.
	Expiration int64 // Unix timestamp.
}

type AddResponse struct {
	Expiration int64 // Unix timestamp.
}

const APIURL = "https://gofs.io/api/v0/"

func NewClient(url string) Client {
	return &client{
		url: url,
	}
}

type client struct {
	url string
}

func (c *client) Pin(ctx context.Context, cid string) (PinResponse, error) {
	panic("unimplemented")
}

func (c *client) Add(ctx context.Context, r io.Reader) (AddResponse, error) {
	panic("unimplemented")
}
