package blockchain

import (
	"math/rand"
	"time"
)

// Validator represents a validator participating in the consensus process.
type Validator struct {
	Address       string
	Stake         float64
	StakingPeriod time.Duration
}

// SelectValidator selects the next validator for block creation based on the weighted staking mechanism.
// This is a placeholder function and should be replaced with the actual weighted staking implementation.
func SelectValidator(validators []*Validator) *Validator {
	return validators[rand.Intn(len(validators))]
}
