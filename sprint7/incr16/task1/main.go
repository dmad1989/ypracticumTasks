package main

import "os"

var dumpFile *os.File

func Dump(data []byte) error {
	if dumpFile == nil {
		// создаём файл при первом обращении
		fname, err := os.Executable()
		if err == nil {
			dumpFile, err = os.Create(fname + `.dump`)
		}
		if err != nil {
			panic(err)
		}
	}
	_, err := dumpFile.Write(data)
	return err
}
