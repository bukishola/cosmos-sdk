package keeper

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

type VestingKeeper struct {
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	key           storetypes.StoreKey
}

func NewVestingKeeper(
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	key storetypes.StoreKey,
) VestingKeeper {
	return VestingKeeper{
		accountKeeper,
		bankKeeper,
		key,
	}
}

// AddVestingAccount adds the address of vesting account to store.
// The caller should check the account type to make sure it's a vesting account type.
func (vk VestingKeeper) AddVestingAccount(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(vk.key)
	store.Set(types.VestingAccountStoreKey(addr), []byte{})
}

// IterateVestingAccounts iterates over all the stored accounts and performs a callback function.
// Stops iteration when callback returns true.
func (vk VestingKeeper) IterateVestingAccounts(ctx sdk.Context, cb func(account exported.VestingAccount) (stop bool)) {
	store := ctx.KVStore(vk.key)
	iterator := sdk.KVStorePrefixIterator(store, types.VestingAccountStoreKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		addr := types.AddressFromVestingAccountKey(iterator.Key())

		acct := vk.accountKeeper.GetAccount(ctx, addr)
		vestingAcct, ok := acct.(exported.VestingAccount)
		if !ok {
			// not vesting account
			continue
		}
		if cb(vestingAcct) {
			break
		}
	}
}
