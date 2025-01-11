package types

const (
	// ModuleName defines the module name
	ModuleName = "librarychain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_librarychain"
)

var (
	ParamsKey = []byte("p_librarychain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	BookKey      = "Book/value/"
	BookCountKey = "Book/count/"
)
