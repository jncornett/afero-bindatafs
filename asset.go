package bindatafs

import "os"

type AssetFunc func(path string) ([]byte, error)
type AssetDirFunc func(path string) ([]string, error)
type AssetInfoFunc func(path string) (os.FileInfo, error)

type Api interface {
	Asset(string) ([]byte, error)
	AssetDir(string) ([]string, error)
	AssetInfo(string) (os.FileInfo, error)
}
