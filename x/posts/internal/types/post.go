package types

import (
	"encoding/json"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ---------------
// --- Post id
// ---------------

// PostID represents a unique post id
type PostID uint64

// Valid tells if the id can be used safely
func (id PostID) Valid() bool {
	return id != 0
}

// Next returns the subsequent id to this one
func (id PostID) Next() PostID {
	return id + 1
}

// String implements fmt.Stringer
func (id PostID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

// Equals compares two PostID instances
func (id PostID) Equals(other PostID) bool {
	return id == other
}

// MarshalJSON implements Marshaler
func (id PostID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

// UnmarshalJSON implements Unmarshaler
func (id *PostID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	postID, err := ParsePostID(s)
	if err != nil {
		return err
	}

	*id = postID
	return nil
}

// ParsePostID returns the PostID represented inside the provided
// value, or an error if no id could be parsed properly
func ParsePostID(value string) (PostID, error) {
	intVal, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return PostID(0), err
	}

	return PostID(intVal), err
}

// ----------------
// --- Post IDs
// ----------------

// PostIDs represents a slice of PostID objects
type PostIDs []PostID

// Equals returns true iff the ids slice and the other
// one contain the same data in the same order
func (ids PostIDs) Equals(other PostIDs) bool {
	if len(ids) != len(other) {
		return false
	}

	for index, id := range ids {
		if id != other[index] {
			return false
		}
	}

	return true
}

// ---------------
// --- Post
// ---------------

// Post is a struct of a Magpie post
type Post struct {
	PostID            PostID         `json:"id"`                 // Unique id
	ParentID          PostID         `json:"parent_id"`          // Post of which this one is a comment
	Message           string         `json:"message"`            // Message contained inside the post
	Created           sdk.Int        `json:"created"`            // Block height at which the post has been created
	LastEdited        sdk.Int        `json:"last_edited"`        // Block height at which the post has been edited the last time
	AllowsComments    bool           `json:"allows_comments"`    // Tells if users can reference this PostID as the parent
	ExternalReference string         `json:"external_reference"` // Used to know when to display this post
	Owner             sdk.AccAddress `json:"owner"`              // Creator of the Post
}

func NewPost(id, parentID PostID, message string, allowsComments bool, externalReference string, created int64, owner sdk.AccAddress) Post {
	return Post{
		PostID:            id,
		ParentID:          parentID,
		Message:           message,
		Created:           sdk.NewInt(created),
		LastEdited:        sdk.ZeroInt(),
		AllowsComments:    allowsComments,
		ExternalReference: externalReference,
		Owner:             owner,
	}
}

// String implements fmt.Stringer
func (p Post) String() string {
	bytes, err := json.Marshal(&p)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// Validate implements validator
func (p Post) Validate() error {
	if !p.PostID.Valid() {
		return fmt.Errorf("invalid post id: %s", p.PostID)
	}

	if p.Owner == nil {
		return fmt.Errorf("invalid post owner: %s", p.Owner)
	}

	if p.Message == "" {
		return fmt.Errorf("invalid post message: %s", p.Message)
	}

	if sdk.ZeroInt().Equal(p.Created) {
		return fmt.Errorf("invalid post creation block height: %s", p.Created)
	}

	if p.Created.GT(p.LastEdited) {
		return fmt.Errorf("invalid post last edit block height: %s", p.LastEdited)
	}

	return nil
}

func (p Post) Equals(other Post) bool {
	return p.PostID.Equals(other.PostID) &&
		p.ParentID.Equals(other.ParentID) &&
		p.Message == other.Message &&
		p.Created.Equal(other.Created) &&
		p.LastEdited.Equal(other.LastEdited) &&
		p.AllowsComments == other.AllowsComments &&
		p.Owner.Equals(other.Owner)
}

// -------------
// --- Posts
// -------------

// Posts represents a slice of Post objects
type Posts []Post

// Equals returns true iff the p slice contains the same
// data in the same order of the other slice
func (p Posts) Equals(other Posts) bool {
	if len(p) != len(other) {
		return false
	}

	for index, post := range p {
		if !post.Equals(other[index]) {
			return false
		}
	}

	return true
}