// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/kava-labs/cosmos-sdk/x/auth/vesting/types/
package vesting

import (
	"github.com/kava-labs/cosmos-sdk/x/auth/vesting/types"
)

var (
	// functions aliases
	RegisterCodec                  = types.RegisterCodec
	NewBaseVestingAccount          = types.NewBaseVestingAccount
	NewContinuousVestingAccountRaw = types.NewContinuousVestingAccountRaw
	NewContinuousVestingAccount    = types.NewContinuousVestingAccount
	NewPeriodicVestingAccountRaw   = types.NewPeriodicVestingAccountRaw
	NewPeriodicVestingAccount      = types.NewPeriodicVestingAccount
	NewDelayedVestingAccountRaw    = types.NewDelayedVestingAccountRaw
	NewDelayedVestingAccount       = types.NewDelayedVestingAccount

	// variable aliases
	VestingCdc = types.VestingCdc
)

type (
	BaseVestingAccount       = types.BaseVestingAccount
	ContinuousVestingAccount = types.ContinuousVestingAccount
	PeriodicVestingAccount   = types.PeriodicVestingAccount
	DelayedVestingAccount    = types.DelayedVestingAccount
	Period                   = types.Period
	Periods                  = types.Periods
)
