// This is a scratchy test program that is here only for experimentation
// during development and may change arbitrarily over time, or be removed
// entirely.
package main

import (
	"log"
	"os"

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

	os.Exit(0)
}
