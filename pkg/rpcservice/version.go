package rpcservice

import (
	"context"

	"github.com/HiteshDatt/stripe-cli/pkg/version"
	"github.com/HiteshDatt/stripe-cli/rpc"
)

// Version returns the version of the Stripe CLI
func (srv *RPCService) Version(ctx context.Context, req *rpc.VersionRequest) (*rpc.VersionResponse, error) {
	return &rpc.VersionResponse{Version: version.Version}, nil
}
