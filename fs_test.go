package bindatafs_test

import (
	"fmt"
	"io/ioutil"
	"log"

	bindatafs "github.com/jncornett/afero-bindatafs"
	"github.com/spf13/afero"
)

var (
	Asset     bindatafs.AssetFunc
	AssetDir  bindatafs.AssetDirFunc
	AssetInfo bindatafs.AssetInfoFunc
)

func ExampleFs() {
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
	fmt.Println("contents of file:", string(b))
}
