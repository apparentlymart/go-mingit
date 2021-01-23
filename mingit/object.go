package mingit

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

// Object represents an object that can be stored in a git repository.
type Object struct {
	raw []byte
}

// ID returns the id of the object, which is currently always a SHA-1 hash
// of the object's encoded content.
func (o Object) ID() ObjectID {
	hash := sha1.Sum(o.raw)
	return ObjectID{hash[:]}
}

// NewBlob encodes and returns a new blob object.
func NewBlob(content []byte) Object {
	raw := []byte(fmt.Sprintf("blob %d\x00%s", len(content), content))
	return Object{raw: raw}
}

// ObjectID represents the ID of a Git object
type ObjectID struct {
	hash []byte
}

func (oid ObjectID) String() string {
	return hex.EncodeToString(oid.hash)
}
