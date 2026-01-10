package storage

import (
	"github.com/devlongs/gean/common/types"
)

// Storage layer for the Lean Ethereum client
// This package provides:
// - Key-value store interface
// - Blockchain data persistence
// - State storage
// - Database management

// Error represents errors that can occur in storage operations
type Error struct {
	Message string
}

func (e Error) Error() string {
	return e.Message
}

// Store represents the storage interface
type Store interface {
	// Block operations
	PutBlock(hash [32]byte, block *types.Block) error
	GetBlock(hash [32]byte) (*types.Block, error)

	// State operations
	PutState(hash [32]byte, state *types.State) error
	GetState(hash [32]byte) (*types.State, error)

	// TODO: Add additional storage operations
}

// MemoryStore is an in-memory implementation of Store
type MemoryStore struct {
	blocks map[[32]byte]*types.Block
	states map[[32]byte]*types.State
}

// NewMemoryStore creates a new in-memory store
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		blocks: make(map[[32]byte]*types.Block),
		states: make(map[[32]byte]*types.State),
	}
}

// PutBlock stores a block in memory
func (m *MemoryStore) PutBlock(hash [32]byte, block *types.Block) error {
	m.blocks[hash] = block
	return nil
}

// GetBlock retrieves a block from memory
func (m *MemoryStore) GetBlock(hash [32]byte) (*types.Block, error) {
	block, exists := m.blocks[hash]
	if !exists {
		return nil, Error{Message: "block not found"}
	}
	return block, nil
}

// PutState stores a state in memory
func (m *MemoryStore) PutState(hash [32]byte, state *types.State) error {
	m.states[hash] = state
	return nil
}

// GetState retrieves a state from memory
func (m *MemoryStore) GetState(hash [32]byte) (*types.State, error) {
	state, exists := m.states[hash]
	if !exists {
		return nil, Error{Message: "state not found"}
	}
	return state, nil
}
