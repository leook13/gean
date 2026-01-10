package forkchoice

import (
	"github.com/devlongs/gean/common/types"
)

// ForkChoice handles the fork choice rule for selecting the head of the chain
// This package contains:
// - Fork choice rule implementation
// - Chain head selection logic
// - Block scoring algorithms

// Error represents errors that can occur during fork choice
type Error struct {
	Message string
}

func (e Error) Error() string {
	return e.Message
}

// ForkChoiceRule represents the fork choice rule implementation
type ForkChoiceRule struct {
	// TODO: Add fork choice state
}

// NewForkChoiceRule creates a new fork choice rule instance
func NewForkChoiceRule() *ForkChoiceRule {
	return &ForkChoiceRule{}
}

// SelectHead selects the head of the chain based on the fork choice rule
func (fc *ForkChoiceRule) SelectHead(blocks []*types.Block, attestations []*types.Attestation) (*types.Block, error) {
	// TODO: Implement fork choice logic
	if len(blocks) == 0 {
		return nil, Error{Message: "no blocks available"}
	}
	return blocks[0], nil
}
