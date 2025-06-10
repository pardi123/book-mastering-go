package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument")
		os.Exit(1) // keluar dengan status 1 = error
	}

	file := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	found := false

	for _, directory := range pathSplit {
		fullpath := filepath.Join(directory, file)
		fileInfo, err := os.Stat(fullpath)
		if err != nil {
			continue
		}

		mode := fileInfo.Mode()
		if !mode.IsRegular() {
			continue
		}
		if mode&0111 != 0 {
			fmt.Println(fullpath)
			found = true
			// Hapus return agar bisa lanjut mencari semua lokasi
		}
	}

	if found {
		os.Exit(0) // Berhasil
	} else {
		os.Exit(1) // Tidak ditemukan
	}
}
