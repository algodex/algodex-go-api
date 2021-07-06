package algodexidx

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"

	inspect "algodexidx/gen/inspect"
	"github.com/algorand/go-algorand-sdk/encoding/msgpack"
	"github.com/algorand/go-algorand-sdk/types"
	"github.com/algorand/go-codec/codec"
)

// inspect service example implementation.
// The example methods log the requests and return zero values.
type inspectsrvc struct {
	logger *log.Logger
}

// NewInspect returns the inspect service implementation.
func NewInspect(logger *log.Logger) inspect.Service {
	return &inspectsrvc{logger}
}

// Unpack a msgpack body (base64 encoded)
func (s *inspectsrvc) Unpack(ctx context.Context, p *inspect.UnpackPayload) (err error) {
	s.logger.Printf("inspect.unpack: body length:%d", len(*p.Msgpack))
	if p.Msgpack == nil {
		return errors.New("must provide msgpack data to unpack")
	}
	//err = backend.WatchAccount(ctx, *p.Account)
	//if err != nil {
	//	return fmt.Errorf("account watch add of address:%s, error:%w", p, err)
	//}
	msgPackData, err := base64.StdEncoding.DecodeString(*p.Msgpack)
	if err != nil {
		return fmt.Errorf("invalid msgpack base64, error: %w", err)
	}
	var (
		count int
		txn   types.SignedTxn
		dec   *codec.Decoder = msgpack.NewDecoder(bytes.NewReader(msgPackData))
	)
	for {
		err = dec.Decode(&txn)
		if errors.Is(io.EOF, err) {
			break
		}
		if err != nil {
			return fmt.Errorf("error in msgpack decode: %w", err)
		}
		fmt.Printf("[%d] %#v\n\n", count, txn)
		if count++; count > 10 {
			break
		}
	}
	//err = nil
	return nil
	/*
		dec := protocol.NewDecoderBytes(data)
		count := 0
		for {
			var txn transactions.SignedTxn
			err = dec.Decode(&txn)
			if err == io.EOF {
				break
			}
			if err != nil {
				reportErrorf(txDecodeError, txFilename, err)
			}
			sti, err := inspectTxn(txn)
			if err != nil {
				reportErrorf(txDecodeError, txFilename, err)
			}
			fmt.Printf("%s[%d]\n%s\n\n", txFilename, count, string(protocol.EncodeJSON(sti)))
			count++
		}
	*/
	return
}
