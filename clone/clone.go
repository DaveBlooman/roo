package clone

import (
	"fmt"
	"log"
	"os"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func Fetch(directory, hash string) error {

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:      "https://github.com/DaveBlooman/go-app",
		Progress: os.Stdout,
	})
	if err != nil {
		log.Println("error cloning")
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		fmt.Println("error")
		return err
	}

	err = w.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash(hash),
	})
	if err != nil {
		fmt.Println("error checking out hash")
		return err
	}

	return nil
}
