package algodexidx

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	inspect "algodexidx/gen/inspect"
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

// Unpack a msgpack body (base64 encoded) and spwan 'goal clerk inspect' to process
// It's a hideous hack but this is just for debugging purposes and the algorand sdks don't provide a means for
// doing this cleanly and without bringing in substantial dependencies (including cgo) from the main Algorand node code.
func (s *inspectsrvc) Unpack(ctx context.Context, p *inspect.UnpackPayload) (res string, err error) {
	s.logger.Printf("inspect.unpack: body length:%d", len(*p.Msgpack))
	if p.Msgpack == nil {
		return "", errors.New("must provide msgpack data to unpack")
	}
	msgPackData, err := base64.StdEncoding.DecodeString(*p.Msgpack)
	if err != nil {
		return "", fmt.Errorf("invalid msgpack base64, error: %w", err)
	}
	tmpfile, err := ioutil.TempFile("", "msgpack")
	if err != nil {
		return "", fmt.Errorf("couldn't create temp file: %w", err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(msgPackData); err != nil {
		return "", fmt.Errorf("couldn't write to temp file: %w", err)
	}
	if err := tmpfile.Close(); err != nil {
		return "", fmt.Errorf("couldn't close temp file: %w", err)
	}
	// TODO: Hardcoded for docker execution environment at this point
	cmd := exec.Command("/node/goal", "clerk", "inspect", tmpfile.Name())
	inspectOutput, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("couldn't execute goal command: %w", err)
	}
	return string(inspectOutput), nil
}
