package keeper

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	StoreKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	Cdc      *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the magpie Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		StoreKey: storeKey,
		Cdc:      cdc,
	}
}

// -------------
// --- Posts
// -------------

func (k Keeper) getPostStoreKey(postID types.PostID) []byte {
	return []byte(types.PostStorePrefix + postID.String())
}

// GetLastPostID returns the last post id that has been used
func (k Keeper) GetLastPostID(ctx sdk.Context) types.PostID {
	store := ctx.KVStore(k.StoreKey)
	if !store.Has([]byte(types.LastPostIDStoreKey)) {
		return types.PostID(0)
	}

	var id types.PostID
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastPostIDStoreKey)), &id)
	return id
}

// SavePost allows to save the given post inside the current context.
// It assumes that the given post has already been validated.
// If another post has the same ID of the given post, the old post will be overridden
func (k Keeper) SavePost(ctx sdk.Context, post types.Post) {
	store := ctx.KVStore(k.StoreKey)

	// Save the post
	store.Set([]byte(types.PostStorePrefix+post.PostID.String()), k.Cdc.MustMarshalBinaryBare(&post))

	// Set the last post id
	store.Set([]byte(types.LastPostIDStoreKey), k.Cdc.MustMarshalBinaryBare(&post.PostID))

	// Save the comments to the parent post, if it is valid
	if post.ParentID.Valid() {
		parentCommentsKey := []byte(types.PostCommentsStorePrefix + post.ParentID.String())

		var commentsIDs []types.PostID
		k.Cdc.MustUnmarshalBinaryBare(store.Get(parentCommentsKey), &commentsIDs)

		commentsIDs = append(commentsIDs, post.PostID)

		store.Set(parentCommentsKey, k.Cdc.MustMarshalBinaryBare(&commentsIDs))
	}
}

// GetPost returns the post having the given id inside the current context.
// If no post having the given id can be found inside the current context, false will be returned.
func (k Keeper) GetPost(ctx sdk.Context, id types.PostID) (post types.Post, found bool) {
	store := ctx.KVStore(k.StoreKey)

	key := k.getPostStoreKey(id)
	if !store.Has(key) {
		return types.Post{}, false
	}

	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &post)
	return post, true
}

// GetPostChildrenIDs returns the IDs of all the children posts associated to the post
// having the given postID
// nolint: interfacer
func (k Keeper) GetPostChildrenIDs(ctx sdk.Context, postID types.PostID) types.PostIDs {
	store := ctx.KVStore(k.StoreKey)

	var postIDs types.PostIDs
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostCommentsStorePrefix+postID.String())), &postIDs)
	return postIDs
}

// GetPosts returns the list of all the posts that are stored into the current state.
func (k Keeper) GetPosts(ctx sdk.Context) []types.Post {
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.PostStorePrefix))

	var posts []types.Post
	for ; iterator.Valid(); iterator.Next() {
		var post types.Post
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &post)
		posts = append(posts, post)
	}

	return posts
}

// -------------
// --- Reactions
// -------------

// SaveReaction allows to save the given reaction inside the store.
// It assumes that the given reaction is valid.
// If another reaction from the same user for the same post and with the same value exists, returns an expError.
// nolint: interfacer
func (k Keeper) SaveReaction(ctx sdk.Context, postID types.PostID, reaction types.Reaction) sdk.Error {
	store := ctx.KVStore(k.StoreKey)
	key := []byte(types.PostReactionsStorePrefix + postID.String())

	// Get the existent reactions
	var reactions types.Reactions
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &reactions)

	// Check for double reactions
	if reactions.ContainsReactionFrom(reaction.Owner, reaction.Value) {
		msg := fmt.Sprintf("%s has already reacted with %s to the post with id %s",
			reaction.Owner, reaction.Value, postID)
		return sdk.ErrUnknownRequest(msg)
	}

	// Save the new reaction
	reactions = append(reactions, reaction)
	store.Set(key, k.Cdc.MustMarshalBinaryBare(&reactions))

	return nil
}

// RemoveReaction removes the reaction from the given user from the post having the
// given postID. If no reaction with the same value was previously added from the given user, an expError
// is returned.
// nolint: interfacer
func (k Keeper) RemoveReaction(ctx sdk.Context, postID types.PostID, user sdk.AccAddress, value string) sdk.Error {
	store := ctx.KVStore(k.StoreKey)
	key := []byte(types.PostReactionsStorePrefix + postID.String())

	// Get the existing reactions
	var reactions types.Reactions
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &reactions)

	// Check if the user exists
	if !reactions.ContainsReactionFrom(user, value) {
		msg := fmt.Sprintf("Cannot remove the reaction with value %s from user %s as it does not exist",
			value, user)
		return sdk.ErrUnauthorized(msg)
	}

	// Remove and save the reactions list
	if newLikes, edited := reactions.RemoveReaction(user, value); edited {
		if len(newLikes) == 0 {
			store.Delete(key)
		} else {
			store.Set(key, k.Cdc.MustMarshalBinaryBare(&newLikes))
		}
	}

	return nil
}

// GetPostReactions returns the list of reactions that has been associated to the post having the given id
// nolint: interfacer
func (k Keeper) GetPostReactions(ctx sdk.Context, postID types.PostID) types.Reactions {
	store := ctx.KVStore(k.StoreKey)

	var reactions types.Reactions
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostReactionsStorePrefix+postID.String())), &reactions)
	return reactions
}

// GetReactions allows to returns the list of reactions that have been stored inside the given context
func (k Keeper) GetReactions(ctx sdk.Context) map[types.PostID]types.Reactions {
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.PostReactionsStorePrefix))

	reactionsData := map[types.PostID]types.Reactions{}
	for ; iterator.Valid(); iterator.Next() {
		var postLikes types.Reactions
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &postLikes)
		postID, _ := types.ParsePostID(strings.TrimPrefix(string(iterator.Key()), types.PostReactionsStorePrefix))
		reactionsData[postID] = postLikes
	}

	return reactionsData
}