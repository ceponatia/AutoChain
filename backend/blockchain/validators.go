package pos

import "time"

type Validator struct {
	Address       string
	Stake         float64
	StakingPeriod time.Duration
}

func NewValidator(address string, stake float64, stakingPeriod time.Duration) *Validator {
	return &Validator{Address: address, Stake: stake, StakingPeriod: stakingPeriod}
}
