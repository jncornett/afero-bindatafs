# afero-bindatafs
An afero filesystem for go-bindata.

## usage
```Go
var fs afero.Fs = bindatafs.Fs{
    AssetFunc: Asset,
    AssetDirFunc: AssetDir,
    AssetDirInfo: AssetInfo,
}
f, err := fs.Open("data/a.txt")
if err != nil {
    log.Fatal(err)
}
b, err := ioutil.ReadAll(f)
if err != nil {
    log.Fatal(err)
}
fmt.Println("contents:", string(b))
```

## related projects
- [afero](https://github.com/spf13/afero) - A filesystem abstraction system for Go
- [go-bindata](https://github.com/jteeuwen/go-bindata) - A small utility which generates Go code from any file. Useful for embedding binary data in a Go program.
