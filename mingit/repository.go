package mingit

import (
	"compress/zlib"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Repository represents a git repository under construction.
type Repository struct {
	gitDir string
}

// NewRepository creates a new directory at the given path and populates it
// with just enough for git to consider it to be an empty repository.
//
// The given directory must not already exist. Note that this creates a bare
// repository, like "git init --bare", so the repository won't have an
// associated work tree. You can clone the generated repository directory
// with "git clone" in order to create a copy with a work tree where you
// could potentially create new commits, etc.
func NewRepository(gitDir string) (*Repository, error) {
	err := os.Mkdir(gitDir, 0755)
	if err != nil {
		return nil, err
	}

	refsDir := filepath.Join(gitDir, "refs")
	tagsDir := filepath.Join(refsDir, "tags")
	headsDir := filepath.Join(refsDir, "heads")
	objectsDir := filepath.Join(gitDir, "objects")
	configPath := filepath.Join(gitDir, "config")
	headPath := filepath.Join(gitDir, "HEAD")

	err = os.Mkdir(refsDir, 0755)
	if err != nil {
		return nil, err
	}
	err = os.Mkdir(tagsDir, 0755)
	if err != nil {
		return nil, err
	}
	err = os.Mkdir(headsDir, 0755)
	if err != nil {
		return nil, err
	}
	err = os.Mkdir(objectsDir, 0755)
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(configPath, []byte(initialConfig), 0644)
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(headPath, []byte(initialHEAD), 0644)
	if err != nil {
		return nil, err
	}

	// We lock the path as given when we're instantiated, in case the
	// larger program switches to a new working directory later.
	absDir, err := filepath.Abs(gitDir)
	if err != nil {
		return nil, err
	}

	return &Repository{
		gitDir: absDir,
	}, nil
}

// WriteObject writes the given object into the object store of the repository.
//
// This is the low-level implementation of WriteBlob, WriteTree, and
// WriteCommit. You only need to do this if you're constructing your objects
// externally.
func (r *Repository) WriteObject(obj Object) (ObjectID, error) {
	id := obj.ID()
	idStr := id.String()
	bucketDir := filepath.Join(r.gitDir, "objects", idStr[:2])
	objFile := filepath.Join(bucketDir, idStr[2:])
	err := os.Mkdir(bucketDir, 0755)
	if err != nil {
		return id, err
	}
	f, err := os.Create(objFile)
	if err != nil {
		return id, err
	}
	wr := zlib.NewWriter(f)
	_, err = wr.Write(obj.raw)
	if err != nil {
		return id, err
	}
	err = wr.Close()
	if err != nil {
		return id, err
	}
	err = f.Close()
	if err != nil {
		return id, err
	}
	return id, nil
}

// WriteBlob writes the given byte slice into the repository as a blob object,
// and returns the id of the object that was created.
func (r *Repository) WriteBlob(data []byte) (ObjectID, error) {
	obj := NewBlob(data)
	return r.WriteObject(obj)
}

// WriteTree writes the given tree data into the repository as a tree object,
// and returns the id of the object that was created.
func (r *Repository) WriteTree(tree Tree) (ObjectID, error) {
	return r.WriteObject(tree.Object())
}

// WriteCommit writes the given commit data into the repository as a commit
// object, and returns the id of the object that was created.
func (r *Repository) WriteCommit(commit *Commit) (ObjectID, error) {
	return r.WriteObject(commit.Object())
}

const initialConfig = `[core]
	repositoryformatversion = 0
	bare = true
`

const initialHEAD = "ref: refs/heads/main"
