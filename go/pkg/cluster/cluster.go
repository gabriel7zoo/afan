package cluster

import (
	"context"
	"fmt"

	"github.com/unkeyed/unkey/go/pkg/events"
	"github.com/unkeyed/unkey/go/pkg/logging"
	"github.com/unkeyed/unkey/go/pkg/membership"
	"github.com/unkeyed/unkey/go/pkg/ring"
)

// Config configures a new cluster instance with the necessary components
// for node management and distributed operations.
type Config struct {
	// Self contains information about the local node
	Self Node

	// Membership provides the underlying node discovery and failure detection
	Membership membership.Membership

	// Logger for cluster operations
	Logger logging.Logger

	// RpcPort is the port used for RPC communication between nodes
	RpcPort int
}

// New creates a new cluster instance with the provided configuration.
// It initializes the consistent hash ring and sets up event listeners
// for membership changes.
//
// Returns an error if the hash ring cannot be created or if the configuration
// is invalid.
func New(config Config) (*cluster, error) {

	r, err := ring.New[Node](ring.Config{
		TokensPerNode: 256,
		Logger:        config.Logger,
	})

	if err != nil {
		return nil, fmt.Errorf("unable to create hash ring: %w", err)
	}

	c := &cluster{
		self:        config.Self,
		membership:  config.Membership,
		ring:        r,
		logger:      config.Logger,
		joinEvents:  events.NewTopic[Node](),
		leaveEvents: events.NewTopic[Node](),

		rpcPort: config.RpcPort,
	}

	go c.keepInSync()

	return c, nil
}

// cluster implements the Cluster interface, combining membership information
// with consistent hashing to provide distributed operations.
type cluster struct {
	self       Node
	membership membership.Membership
	ring       *ring.Ring[Node]
	logger     logging.Logger

	joinEvents  events.Topic[Node]
	leaveEvents events.Topic[Node]

	rpcPort int
}

// Ensure cluster implements the Cluster interface
var _ Cluster = (*cluster)(nil)

// Self returns information about the local node.
func (c *cluster) Self() Node {
	return c.self
}

// SubscribeJoin returns a channel that receives Node events
// whenever a new node joins the cluster.
func (c *cluster) SubscribeJoin() <-chan Node {
	return c.joinEvents.Subscribe("cluster.joinEvents")
}

// SubscribeLeave returns a channel that receives Node events
// whenever a node leaves the cluster.
func (c *cluster) SubscribeLeave() <-chan Node {
	return c.leaveEvents.Subscribe("cluster.leaveEvents")
}

// keepInSync listens to membership changes and updates the hash ring accordingly.
// This is run in a separate goroutine to handle cluster topology changes.
func (c *cluster) keepInSync() {
	joins := c.membership.SubscribeJoinEvents()
	leaves := c.membership.SubscribeLeaveEvents()

	for {
		select {
		case node := <-joins:
			{
				ctx := context.Background()
				c.logger.Info("node joined", "nodeID", node.NodeID)

				err := c.ring.AddNode(ctx, ring.Node[Node]{
					ID: node.NodeID,
					Tags: Node{
						RpcAddr: fmt.Sprintf("%s:%d", node.Addr, c.rpcPort),
						ID:      node.NodeID,
						Addr:    node.Addr,
					},
				})
				if err != nil {
					c.logger.Error("failed to add node to ring", "error", err.Error())
				}

			}
		case node := <-leaves:
			{
				ctx := context.Background()
				c.logger.Info("node left", "nodeID", node.NodeID)
				err := c.ring.RemoveNode(ctx, node.NodeID)
				if err != nil {
					c.logger.Error("failed to remove node from ring", "error", err.Error())
				}
			}
		}

	}

}

// FindNode determines which node is responsible for a given key
// based on consistent hashing.
func (c *cluster) FindNode(ctx context.Context, key string) (Node, error) {
	node, err := c.ring.FindNode(key)
	if err != nil {
		return Node{}, err
	}
	return node.Tags, nil

}

// Shutdown gracefully exits the cluster, notifying other nodes
// and cleaning up resources.
func (c *cluster) Shutdown(ctx context.Context) error {
	return c.membership.Leave()
}
