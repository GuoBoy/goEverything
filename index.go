package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// pre check
func preCheck(temp string) error {
	stat, err := os.Stat(temp)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New(temp + " is not found, please check input")
		}
		return err
	}
	if !stat.IsDir() {
		return errors.New(temp + " is a file")
	}
	return nil
}

// IndexFile index file by input
func IndexFile(inputPath string) error {
	if err := preCheck(inputPath); err != nil {
		return err
	}
	// store disk
	id := ToMd5String(inputPath)
	diskStore := DiskStore{D: &Disk{
		ID:   id,
		Name: inputPath,
	}}
	if err := diskStore.Add(); err != nil {
		return err
	}
	store := ItemStore{}
	if err := store.Delete(id); err != nil {
		return err
	}
	// start
	var items []*Item
	if err := filepath.Walk(inputPath, func(ph string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(info)
		items = append(items, &Item{
			ID:     ToMd5String(ph),
			Path:   ph,
			Name:   info.Name(),
			Size:   info.Size(),
			Type:   Bool2Type(info.IsDir()),
			DiskID: id,
		})
		return nil
	}); err != nil {
		return err
	}
	store.Add(items)
	return nil
}
