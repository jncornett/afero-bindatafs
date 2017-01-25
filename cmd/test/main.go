package main

import (
	"fmt"
	"io/ioutil"
	"log"

	bindatafs "github.com/jncornett/afero-bindatafs"
	"github.com/spf13/afero"
)

//go:generate go-bindata data/...

func main() {
	var fs afero.Fs = &bindatafs.Fs{
		AssetFunc:     Asset,
		AssetInfoFunc: AssetInfo,
		AssetDirFunc:  AssetDir,
	}
	f, err := fs.Open("data/a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
