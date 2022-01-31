// Code generated by goa v3.5.4, DO NOT EDIT.
//
// hellosvc HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/pvr1/goatest/design

package server

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

// EncodeGreetResponse returns an encoder for responses returned by the
// hellosvc greet endpoint.
func EncodeGreetResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGreetRequest returns a decoder for requests sent to the hellosvc greet
// endpoint.
func DecodeGreetRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			name string

			params = mux.Vars(r)
		)
		name = params["name"]
		payload := NewGreetPayload(name)

		return payload, nil
	}
}