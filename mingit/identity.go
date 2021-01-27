package mingit

import (
	"fmt"
	"time"
)

// Identity represents a name, address, and timestamp that can appear as
// either the author or the committer of a commit object.
type Identity struct {
	Name  string
	Email string
	Time  time.Time
}

func (i Identity) String() string {
	return fmt.Sprintf("%s <%s> %d +0000", i.Name, i.Email, i.Time.Unix())
}
