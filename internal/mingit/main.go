// This is a scratchy test program that is here only for experimentation
// during development and may change arbitrarily over time, or be removed
// entirely.
package main

import (
	"log"
	"os"
	"time"

	"github.com/apparentlymart/go-mingit/mingit"
)

func main() {
	repo, err := mingit.NewRepository("/tmp/foo")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v", repo)

	blobID, err := repo.WriteBlob([]byte("hello world\n"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("created blob %s", blobID)

	treeID, err := repo.WriteTree(mingit.Tree{
		{
			Mode:     mingit.ModeRegular,
			Name:     "hello.txt",
			TargetID: blobID,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("created tree %s", treeID)

	author := &mingit.Identity{
		Name:  "Bob Bobbins",
		Email: "bbobbins@bobco.com",
		Time:  time.Now(),
	}
	commitID, err := repo.WriteCommit(&mingit.Commit{
		TreeID:    treeID,
		Message:   "Hello world",
		Author:    author,
		Committer: author,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("created commit %s", commitID)

	err = repo.SetRef("refs/heads/main", commitID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("updated refs/heads/main to refer to %s", commitID)

	commitID, err = repo.WriteCommit(&mingit.Commit{
		TreeID:    treeID,
		Message:   "Hello world 2",
		ParentIDs: []mingit.ObjectID{commitID},
		Author:    author,
		Committer: author,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("created another commit %s", commitID)

	err = repo.SetRef("refs/heads/main", commitID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("updated refs/heads/main to refer to %s", commitID)

	os.Exit(0)
}
