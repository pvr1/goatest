// Code generated by goa v3.5.4, DO NOT EDIT.
//
// hellosvc client
//
// Command:
// $ goa gen github.com/pvr1/goatest/design

package hellosvc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "hellosvc" service client.
type Client struct {
	GreetEndpoint goa.Endpoint
}

// NewClient initializes a "hellosvc" service client given the endpoints.
func NewClient(greet goa.Endpoint) *Client {
	return &Client{
		GreetEndpoint: greet,
	}
}

// Greet calls the "greet" endpoint of the "hellosvc" service.
func (c *Client) Greet(ctx context.Context, p *GreetPayload) (res string, err error) {
	var ires interface{}
	ires, err = c.GreetEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}
