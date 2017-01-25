package bindatafs

import (
	"bytes"
	"os"
	"syscall"

	"github.com/spf13/afero"
)

type File struct {
	name string
	*bytes.Reader
	Api
}

func (f *File) Close() error                       { return nil }
func (f *File) Write([]byte) (int, error)          { return 0, syscall.EPERM }
func (f *File) WriteAt([]byte, int64) (int, error) { return 0, syscall.EPERM }
func (f *File) Sync() error                        { return syscall.EPERM }
func (f *File) Truncate(int64) error               { return syscall.EPERM }
func (f *File) WriteString(string) (int, error)    { return 0, syscall.EPERM }

func (f *File) Name() string { return f.name }

func (f *File) Readdir(int) ([]os.FileInfo, error) {
	return nil, &os.SyscallError{
		Syscall: "readdirent",
		Err:     syscall.EINVAL,
	}
}

func (f *File) Readdirnames(int) ([]string, error) {
	return nil, &os.SyscallError{
		Syscall: "readdirent",
		Err:     syscall.EINVAL,
	}
}

func (f *File) Stat() (os.FileInfo, error) { return f.AssetInfo(f.name) }

var _ afero.File = &File{}
