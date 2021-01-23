package mingit

import (
	"time"
)

// Identity represents a name, address, and timestamp that can appear as
// either the author or the committer of a commit object.
type Identity struct {
	Name  string
	Email string
	Time  time.Time
}
