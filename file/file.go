package file

import (
	"fmt"
	"os"
	"path/filepath"
)

// File file impl io.Writer
type File struct {
	opts Options
	fd   *os.File
	size int64
	line int64
}

// New file
func New(opts ...Option) *File {
	options := newOptions(opts...)
	return &File{
		opts: options,
	}
}

// Init init
func (f *File) Init(opts ...Option) {
	for _, o := range opts {
		o(&f.opts)
	}
}

func (f *File) Write(b []byte) (int, error) {
	if err := f.rorate(); err != nil {
		return 0, err
	}
	n, err := f.fd.Write(b)
	f.size += int64(n)
	return n, err
}

func (f *File) reset() {
	f.size = 0
	f.line = 0
}

func (f *File) openFile(file string) error {
	var err error
	if err = os.MkdirAll(filepath.Dir(file), 0755); err != nil {
		return err
	}
	if f.fd, err = os.OpenFile(file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644); err != nil {
		return fmt.Errorf("open %s err: %v", file, err)
	}
	return nil
}

func (f *File) rorate() error {
	var err error
	file := WithTime(filepath.Join(f.opts.dir, f.opts.path))
	if f.fd == nil {
		if err = f.openFile(file); err != nil {
			return err
		}
	}
	if f.fd.Name() != file {
		if err = f.fd.Close(); err != nil {
			return fmt.Errorf("close %s err: %v", f.fd.Name(), err)
		}
		f.reset()
		if err = f.openFile(file); err != nil {
			return err
		}
	}
	// info, err := f.fd.Stat()
	// if err != nil {
	// 	return err
	// }
	// if info.Size() > f.opts.Size {
	// 	if err = f.fd.Close(); err != nil {
	// 		return fmt.Errorf("close %s err: %v", f.fd.Name(), err)
	// 	}
	// 	f.reset()
	// 	if err = f.openFile(file); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}
