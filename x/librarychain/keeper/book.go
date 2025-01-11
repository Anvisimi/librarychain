package keeper

import (
	"context"
	"encoding/binary"

	"librarychain/x/librarychain/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetBookCount get the total number of book
func (k Keeper) GetBookCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.BookCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetBookCount set the total number of book
func (k Keeper) SetBookCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.BookCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendBook appends a book in the store with a new id and update the count
func (k Keeper) AppendBook(
	ctx context.Context,
	book types.Book,
) uint64 {
	// Create the book
	count := k.GetBookCount(ctx)

	// Set the ID of the appended value
	book.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BookKey))
	appendedValue := k.cdc.MustMarshal(&book)
	store.Set(GetBookIDBytes(book.Id), appendedValue)

	// Update book count
	k.SetBookCount(ctx, count+1)

	return count
}

// SetBook set a specific book in the store
func (k Keeper) SetBook(ctx context.Context, book types.Book) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BookKey))
	b := k.cdc.MustMarshal(&book)
	store.Set(GetBookIDBytes(book.Id), b)
}

// GetBook returns a book from its id
func (k Keeper) GetBook(ctx context.Context, id uint64) (val types.Book, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BookKey))
	b := store.Get(GetBookIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBook removes a book from the store
func (k Keeper) RemoveBook(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BookKey))
	store.Delete(GetBookIDBytes(id))
}

// GetAllBook returns all book
func (k Keeper) GetAllBook(ctx context.Context) (list []types.Book) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BookKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Book
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetBookIDBytes returns the byte representation of the ID
func GetBookIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.BookKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
