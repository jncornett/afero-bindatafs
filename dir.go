package bindatafs

import (
	"os"
	"path/filepath"
	"syscall"

	"github.com/spf13/afero"
)

type Dir struct {
	name string
	Api
}

func (d *Dir) Close() error                       { return nil }
func (d *Dir) Write([]byte) (int, error)          { return 0, syscall.EPERM }
func (d *Dir) WriteAt([]byte, int64) (int, error) { return 0, syscall.EPERM }
func (d *Dir) Sync() error                        { return syscall.EPERM }
func (d *Dir) Truncate(int64) error               { return syscall.EPERM }
func (d *Dir) WriteString(string) (int, error)    { return 0, syscall.EPERM }

func (d *Dir) Read([]byte) (int, error) {
	return 0, &os.PathError{
		Op:   "read",
		Path: d.name,
		Err:  syscall.EISDIR,
	}
}

func (d *Dir) ReadAt([]byte, int64) (int, error) {
	return 0, &os.PathError{
		Op:   "read",
		Path: d.name,
		Err:  syscall.EISDIR,
	}
}

func (d *Dir) Seek(int64, int) (int64, error) {
	return 0, &os.PathError{
		Op:   "seek",
		Path: d.name,
		Err:  syscall.EISDIR,
	}
}

func (d *Dir) Name() string                       { return d.name }
func (d *Dir) Readdirnames(int) ([]string, error) { return d.AssetDir(d.name) }
func (d *Dir) Stat() (os.FileInfo, error)         { return d.AssetInfo(d.name) }

func (d *Dir) Readdir(n int) (out []os.FileInfo, err error) {
	var names []string
	if names, err = d.AssetDir(d.name); err != nil {
		return
	}
	for _, name := range names {
		var fi os.FileInfo
		if fi, err = d.AssetInfo(filepath.Join(d.name, name)); err != nil {
			return
		}
		out = append(out, fi)
	}
	return
}

var _ afero.File = &Dir{}
