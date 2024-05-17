package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// CreateFiles создаёт поддиректории и файлы во временной директории.
func CreateFiles() (string, string) {
	var ind int

	curdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// создаём временную директорию
	tmp, err := os.MkdirTemp("", "tmp")
	if err != nil {
		log.Fatal(err)
	}
	if err = os.Chdir(tmp); err != nil {
		log.Fatal(err)
	}
	createFile := func(subdir string) {
		// создаём файл и записываем туда его имя
		ind++
		name := filepath.Join(subdir, fmt.Sprintf("file-%d.txt", ind))
		if err = os.WriteFile(name, []byte(name), 0666); err != nil {
			log.Fatal(err)
		}
	}
	// создаём поддиректории с двумя файлами в каждой
	for _, name := range []string{"folder-a", "folder-b/folder_c"} {
		os.MkdirAll(name, 0755)
		createFile(name)
		createFile(name)
	}
	createFile("")
	return tmp, curdir
}

// MyWalkDir рекурсивно проходит по всем поддиректориям.
func MyWalkDir(dir string, shift string) {
	list, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range list {
		finfo, err := item.Info()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s%s %v", shift, item.Name(), finfo.Mode())
		if item.IsDir() {
			fmt.Println()
			// рекурсивно заходим в поддиректорию
			MyWalkDir(filepath.Join(dir, item.Name()), shift+"   ")
		} else {
			fmt.Printf(" %dB %s\n", finfo.Size(), finfo.ModTime().Format("2006/01/02 15:04:05"))
		}
	}
}

func main() {
	// создаём набор файлов во временной директории
	dir, curdir := CreateFiles()
	MyWalkDir(dir, "")
	os.Chdir(curdir)
	// можем удалять временную директорию
	os.RemoveAll(dir)
}

func MyWalkDir2(dir string) {
	dircount := len(strings.Split(dir, string(os.PathSeparator)))
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == dir {
			// пропускаем начальную директорию
			return nil
		}
		shift := strings.Repeat(` `, len(strings.Split(path,
			string(os.PathSeparator)))-dircount-1)
		fmt.Printf("%s%s %v", shift, info.Name(), info.Mode())
		if info.IsDir() {
			// рекурсивный вызов не нужен, так как filepath.Walk
			// проходит по всем поддиректориям
			fmt.Println()
		} else {
			fmt.Printf(" %dB %s\n", info.Size(),
				info.ModTime().Format("2006/01/02 15:04:05"))
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
