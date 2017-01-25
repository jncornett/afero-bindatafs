package bindatafs

import (
	"bytes"
	"os"
	"syscall"
	"time"

	"github.com/spf13/afero"
)

type Fs struct {
	AssetFunc
	AssetDirFunc
	AssetInfoFunc
}

func (fs *Fs) Chmod(string, os.FileMode) error            { return syscall.EPERM }
func (fs *Fs) Chtimes(string, time.Time, time.Time) error { return syscall.EPERM }
func (fs *Fs) Create(string) (afero.File, error)          { return nil, syscall.EPERM }
func (fs *Fs) Mkdir(string, os.FileMode) error            { return syscall.EPERM }
func (fs *Fs) MkdirAll(string, os.FileMode) error         { return syscall.EPERM }
func (fs *Fs) Remove(string) error                        { return syscall.EPERM }
func (fs *Fs) RemoveAll(string) error                     { return syscall.EPERM }
func (fs *Fs) Rename(string, string) error                { return syscall.EPERM }

func (fs *Fs) OpenFile(
	name string,
	flag int,
	mode os.FileMode,
) (afero.File, error) {
	if flag&(os.O_WRONLY|syscall.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_TRUNC) != 0 {
		return nil, syscall.EPERM
	}
	data, err := fs.Asset(name)
	if err != nil {
		return nil, err
	}
	return &File{name: name, Reader: bytes.NewReader(data), Api: fs}, nil
}

func (fs *Fs) Open(name string) (afero.File, error) {
	data, err := fs.Asset(name)
	if err != nil {
		return nil, err
	}
	return &File{name: name, Reader: bytes.NewReader(data), Api: fs}, nil
}

func (fs *Fs) Stat(name string) (os.FileInfo, error) {
	return fs.AssetInfo(name)
}

func (fs *Fs) Name() string {
	return "Bindata"
}

var _ afero.Fs = &Fs{}

func (fs *Fs) Asset(name string) ([]byte, error) {
	return fs.AssetFunc(name)
}

func (fs *Fs) AssetDir(name string) ([]string, error) {
	return fs.AssetDirFunc(name)
}

func (fs *Fs) AssetInfo(name string) (os.FileInfo, error) {
	return fs.AssetInfoFunc(name)
}

var _ Api = &Fs{}
