package types

const (
	// ModuleName defines the module name
	ModuleName = "cosa"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	AuctionCountKey = "cosa/count"

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cosa"
)

var (
	ParamsKey = []byte("p_cosa")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
