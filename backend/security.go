package backend

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"algodexidx/gen/account"
)

// private type used to define context keys
type ctxKey int

// RemoteIPKey is the request context key used to store the remote IP address set by the SetRemoteIP middleware.
const RemoteIPKey ctxKey = 1

// SetRemoteIP is middleware we inject to set the calling ip into the context so we can fetch later in the implementations
// and use for things like network ACLs
func SetRemoteIP() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				var remoteAddr = r.RemoteAddr
				if f := r.Header.Get("X-Forwarded-For"); f != "" {
					remoteAddr = f
				}
				ctx := context.WithValue(r.Context(), RemoteIPKey, remoteAddr)
				h.ServeHTTP(w, r.WithContext(ctx))
			},
		)
	}
}

// GetRemoteIP returns the remote ip set into our context
func GetRemoteIP(ctx context.Context) string {
	if i := ctx.Value(RemoteIPKey); i != nil {
		return i.(string)
	}
	return ""
}

func IsAddressInAllowedSubnets(ctx context.Context, remoteAddress string) (bool, error) {
	allowList := os.Getenv("ALGODEX_SUBNET_WHITELIST")
	if allowList == "" {
		return false, errors.New("allow-list not defined, rejecting all")
	}
	// Allow multiple comma delimited subnets
	for _, cidrSpec := range strings.Split(allowList, ",") {
		_, subnet, err := net.ParseCIDR(cidrSpec)
		if err != nil {
			return false, fmt.Errorf("error in IsAddressInAllowedSubnets: %w", err)
		}
		host, _, err := net.SplitHostPort(remoteAddress)
		if host == "" {
			host = remoteAddress
		}
		if subnet.Contains(net.ParseIP(host)) {
			return true, nil
		}
	}
	return false, nil
}

func FailIfNotAuthorized(ctx context.Context) error {
	remoteIP := GetRemoteIP(ctx)
	allowed, err := IsAddressInAllowedSubnets(ctx, remoteIP)
	if err != nil {
		return account.MakeAccessDenied(err)
	}
	if !allowed {
		return account.MakeAccessDenied(fmt.Errorf("%v was blocked", remoteIP))
	}
	return nil
}
