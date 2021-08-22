// Code generated by goa v3.4.3, DO NOT EDIT.
//
// orders HTTP client encoders and decoders
//
// Command:
// $ goa gen algodexidx/design

package client

import (
	orders "algodexidx/gen/orders"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildGetRequest instantiates a HTTP request object with method and path set
// to call the "orders" service "get" endpoint
func (c *Client) BuildGetRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetOrdersPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("orders", "get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetRequest returns an encoder for requests sent to the orders get
// server.
func EncodeGetRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*orders.GetPayload)
		if !ok {
			return goahttp.ErrInvalidType("orders", "get", "*orders.GetPayload", v)
		}
		values := req.URL.Query()
		if p.AssetID != nil {
			values.Add("assetId", fmt.Sprintf("%v", *p.AssetID))
		}
		for _, value := range p.OwnerAddr {
			values.Add("ownerAddr", value)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeGetResponse returns a decoder for responses returned by the orders get
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
// DecodeGetResponse may return the following errors:
//	- "access_denied" (type *goa.ServiceError): http.StatusUnauthorized
//	- error: internal error
func DecodeGetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("orders", "get", err)
			}
			err = ValidateGetResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("orders", "get", err)
			}
			res := NewGetOrdersOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body GetAccessDeniedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("orders", "get", err)
			}
			err = ValidateGetAccessDeniedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("orders", "get", err)
			}
			return nil, NewGetAccessDenied(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("orders", "get", resp.StatusCode, string(body))
		}
	}
}

// unmarshalOrderResponseBodyToOrdersOrder builds a value of type *orders.Order
// from a value of type *OrderResponseBody.
func unmarshalOrderResponseBodyToOrdersOrder(v *OrderResponseBody) *orders.Order {
	if v == nil {
		return nil
	}
	res := &orders.Order{
		AssetLimitPriceInAlgos:     *v.AssetLimitPriceInAlgos,
		AsaPrice:                   *v.AsaPrice,
		AssetLimitPriceD:           *v.AssetLimitPriceD,
		AssetLimitPriceN:           *v.AssetLimitPriceN,
		AlgoAmount:                 *v.AlgoAmount,
		AsaAmount:                  *v.AsaAmount,
		AssetID:                    *v.AssetID,
		AppID:                      *v.AppID,
		EscrowAddress:              *v.EscrowAddress,
		OwnerAddress:               *v.OwnerAddress,
		MinimumExecutionSizeInAlgo: *v.MinimumExecutionSizeInAlgo,
		Round:                      *v.Round,
		UnixTime:                   *v.UnixTime,
		FormattedPrice:             *v.FormattedPrice,
		FormattedASAAmount:         *v.FormattedASAAmount,
		Decimals:                   *v.Decimals,
	}

	return res
}
