package mingit

import (
	"bytes"
	"fmt"
)

// Commit represents the data for a commit, which you can then encode into
// a Git object using the Object method.
type Commit struct {
	TreeID    ObjectID
	ParentIDs []ObjectID
	Author    *Identity
	Committer *Identity
	Message   string
}

// Object encodes the data from the Commit into a Git object ready to be
// stored into a repository.
func (c *Commit) Object() Object {
	content := c.ForStorage()
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "commit %d\x00%s", len(content), content)
	return Object{
		raw: buf.Bytes(),
	}
}

// ForStorage returns a byte slice containing the commit data in the raw
// format that git would use inside a generated object.
func (c *Commit) ForStorage() []byte {
	return c.appendRaw(nil)
}

func (c *Commit) appendRaw(to []byte) []byte {
	buf := bytes.NewBuffer(to)
	fmt.Fprintf(buf, "tree %s\n", c.TreeID)
	for _, id := range c.ParentIDs {
		fmt.Fprintf(buf, "parent %s\n", id)
	}
	fmt.Fprintf(buf, "author %s\n", c.Author)
	fmt.Fprintf(buf, "committer %s\n\n", c.Committer)
	buf.WriteString(c.Message)
	return buf.Bytes()
}
