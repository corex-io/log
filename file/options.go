package file

// Options options
type Options struct {
	dir  string // 日志目录
	path string // 文件
	Size int64  // 最大size
}

// Option func options
type Option func(*Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		dir:  ".",
		path: "logs/%Y%M%D.log",
		Size: 1024 * 1024, // 1GB
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// Dir dir
func Dir(dir string) Option {
	return func(o *Options) {
		o.dir = dir
	}
}

// Path path
func Path(path string) Option {
	return func(o *Options) {
		o.path = path
	}
}

// MaxSize max size (bytes)
func MaxSize(size int64) Option {
	return func(o *Options) {
		o.Size = size
	}
}
