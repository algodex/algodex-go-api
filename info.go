package algodexidx

import (
	"context"
	"log"

	info "algodexidx/gen/info"
)

// info service example implementation.
// The example methods log the requests and return zero values.
type infosrvc struct {
	logger         *log.Logger
	versionSummary string
}

// NewInfo returns the info service implementation.
func NewInfo(logger *log.Logger, versionSummary string) info.Service {
	return &infosrvc{logger, versionSummary}
}

// Returns version information for the service
func (s *infosrvc) Version(ctx context.Context) (res string, err error) {
	s.logger.Print("info.version")
	return s.versionSummary, nil
}

// Simple health check
func (s *infosrvc) Live(ctx context.Context) (err error) {
	s.logger.Print("info.live")
	return
}
