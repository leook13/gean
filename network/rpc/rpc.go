package rpc

import (
	"context"

	"github.com/devlongs/gean/common/types"
)

// RPC handles the JSON-RPC interface for the Lean Ethereum client
// This package contains:
// - JSON-RPC server
// - API endpoints
// - Client interface

// Error represents errors that can occur in RPC operations
type Error struct {
	Message string
	Code    int
}

func (e Error) Error() string {
	return e.Message
}

// RPCServer represents the JSON-RPC server
type RPCServer struct {
	// TODO: Add RPC server fields
}

// NewRPCServer creates a new RPC server instance
func NewRPCServer() *RPCServer {
	return &RPCServer{}
}

// Start starts the RPC server
func (rpc *RPCServer) Start(ctx context.Context, addr string) error {
	// TODO: Implement RPC server startup
	return nil
}

// GetBlock returns a block by hash or number
func (rpc *RPCServer) GetBlock(ctx context.Context, blockHashOrNumber interface{}) (*types.Block, error) {
	// TODO: Implement get block logic
	return nil, Error{Message: "not implemented", Code: -1}
}

// GetState returns the current state
func (rpc *RPCServer) GetState(ctx context.Context) (*types.State, error) {
	// TODO: Implement get state logic
	return nil, Error{Message: "not implemented", Code: -1}
}
