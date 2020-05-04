package store

import (
	dbm "github.com/tendermint/tm-db"

	"github.com/kava-labs/cosmos-sdk/store/cache"
	"github.com/kava-labs/cosmos-sdk/store/rootmulti"
	"github.com/kava-labs/cosmos-sdk/store/types"
)

// Pruning strategies that may be provided to a KVStore to enable pruning.
const (
	PruningStrategyNothing    = "nothing"
	PruningStrategyEverything = "everything"
	PruningStrategySyncable   = "syncable"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}

func NewPruningOptionsFromString(strategy string) (opt PruningOptions) {
	switch strategy {
	case PruningStrategyNothing:
		opt = PruneNothing
	case PruningStrategyEverything:
		opt = PruneEverything
	case PruningStrategySyncable:
		opt = PruneSyncable
	default:
		opt = PruneSyncable
	}
	return
}
