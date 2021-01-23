package mingit

import (
	"bytes"
	"fmt"
	"os"
)

// Tree represents the data for a git tree, which you can then encode into
// a Git object using the Object method.
type Tree []TreeItem

// Object encodes the data from the Tree into a Git object ready to be
// stored into a repository.
func (t Tree) Object() Object {
	var raw []byte
	for _, item := range t {
		raw = item.appendRaw(raw)
	}
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "tree %d\x00%s", len(raw), raw)
	return Object{
		raw: buf.Bytes(),
	}
}

// TreeItem represents a single item in a tree.
type TreeItem struct {
	Mode     os.FileMode
	Name     string
	TargetID ObjectID
}

func (ti *TreeItem) appendRaw(to []byte) []byte {
	buf := bytes.NewBuffer(to)
	fmt.Fprintf(buf, "%o %s\x00%s", ti.Mode, ti.Name, ti.TargetID.hash)
	return buf.Bytes()
}

const (
	// ModeRegular is the file mode Git uses for regular (non-executable) files.
	// The target must be a blob in this case.
	ModeRegular os.FileMode = 0100644

	// ModeExecutable is the file mode Git uses for executable files.
	// The target must be a blob in this case.
	ModeExecutable os.FileMode = 0100755

	// ModeDir is the file mode Git uses for directories.
	// The target must be a tree in this case.
	ModeDir os.FileMode = 040000
)
