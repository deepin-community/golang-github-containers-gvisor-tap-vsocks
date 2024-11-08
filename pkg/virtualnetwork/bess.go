package virtualnetwork

import (
	"context"
	"net"

	"github.com/containers/gvisor-tap-vsock/pkg/types"
)

func (n *VirtualNetwork) AcceptBess(ctx context.Context, conn net.Conn) error {
	return n.networkSwitch.Accept(ctx, conn, types.BessProtocol)
}
