package logfs

import (
	"os"
	"time"

	"github.com/spf13/afero"
)

// Fs wraps an afero.Fs and adds a logging layer
type Fs struct {
	Logger
	Fs afero.Fs
}

func (fs Fs) Create(name string) (f afero.File, err error) {
	start := time.Now()
	defer func() { fs.Log(newRecord("create", err, start, name)) }()
	f, err = fs.Fs.Create(name)
	return
}

func (fs Fs) Mkdir(name string, mode os.FileMode) (err error) {
	start := time.Now()
	defer func() { fs.Log(newRecord("mkdir", err, start, name, mode)) }()
	err = fs.Fs.Mkdir(name, mode)
	return
}

func (fs Fs) MkdirAll(name string, mode os.FileMode) (err error) {
	start := time.Now()
	defer func() { fs.Log(newRecord("mkdirall", err, start, name, mode)) }()
	err = fs.Fs.MkdirAll(name, mode)
	return
}

func (fs Fs) Open(name string) (f afero.File, err error) {
	start := time.Now()
	defer func() { fs.Log(newRecord("open", err, start, name)) }()
	f, err = fs.Fs.Open(name)
	return
}

func (fs Fs) OpenFile(
	name string,
	flag int,
	perm os.FileMode,
) (f afero.File, err error) {
	start := time.Now()
	defer func() { fs.Log(newRecord("openfile", err, start, name, flag, perm)) }()
	f, err = fs.Fs.OpenFile(name, flag, perm)
	return
}

func (fs Fs) Remove(name string) (err error) {
	start := time.Now()
	defer func() { fs.Log(newRecord("remove", err, start, name)) }()
	err = fs.Fs.Remove(name)
	return
}

func (fs Fs) RemoveAll(name string) (err error) {
	start := time.Now()
	defer func() { fs.Log(newRecord("removeall", err, start, name)) }()
	err = fs.Fs.RemoveAll(name)
	return
}

func (fs Fs) Rename(oldname, newname string) (err error) {
	start := time.Now()
	defer func() { fs.Log(newRecord("rename", err, start, oldname, newname)) }()
	err = fs.Fs.Rename(oldname, newname)
	return
}

func (fs Fs) Stat(name string) (fi os.FileInfo, err error) {
	start := time.Now()
	defer func() { fs.Log(newRecord("stat", err, start, name)) }()
	fi, err = fs.Fs.Stat(name)
	return
}

func (fs Fs) Chmod(name string, mode os.FileMode) (err error) {
	start := time.Now()
	defer func() { fs.Log(newRecord("chmod", err, start, name, mode)) }()
	err = fs.Fs.Chmod(name, mode)
	return
}

func (fs Fs) Chtimes(name string, atime, mtime time.Time) (err error) {
	start := time.Now()
	defer func() { fs.Log(newRecord("chmod", err, start, name, atime, mtime)) }()
	err = fs.Fs.Chtimes(name, atime, mtime)
	return
}

func (fs Fs) Name() string {
	return fs.Fs.Name()
}

func newRecord(op string, err error, start time.Time, args ...interface{}) Record {
	return Record{
		Op:   op,
		Args: args,
		Err:  err,
		Time: time.Since(start),
	}
}
