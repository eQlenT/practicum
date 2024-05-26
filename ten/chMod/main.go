package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fname := "/home/ivan/scripts/readme.txt"
	err := os.Chmod(fname, 0777)
	if err != nil {
		log.Fatal(err)
	}
	fi, err := os.Stat(fname)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Права доступа: 0%o\n", fi.Mode().Perm())
}
