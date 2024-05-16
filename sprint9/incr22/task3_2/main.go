package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func LastModified(dir string, hours int) error {
	list, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	t := time.Now().Add(-time.Hour * time.Duration(hours))
	for _, item := range list {
		finfo, err := item.Info()
		if err != nil {
			return err
		}
		if item.IsDir() {
			if item.IsDir() {
				LastModified(filepath.Join(dir, item.Name()), hours)
			}
		} else {
			if finfo.ModTime().After(t) {
				fmt.Printf("%s %s\n", filepath.Join(dir, item.Name()),
					finfo.ModTime().Format("2006/01/02 15:04:05"))
			}
		}
	}
	return nil
}

func main() {
	LastModified(`/home/ak/.cinnamon`, 72)
}
