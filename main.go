package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

func main() {
	files, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	var ll []fs.DirEntry
	ll = append(ll, files...)

	sort.Slice(ll, func(i, j int) bool {
		infoA, err := ll[i].Info()
		if err != nil {
			log.Fatal(err)
		}

		infoB, err := ll[j].Info()
		if err != nil {
			log.Fatal(err)
		}

		return infoA.ModTime().Before(infoB.ModTime())
	})

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	for i, f := range ll {
		if filepath.Join(exPath, f.Name()) != os.Args[0] {
			os.Rename(filepath.Join(exPath, f.Name()), filepath.Join(exPath, strconv.Itoa(i+1)+". "+f.Name()))
		}
	}

}
