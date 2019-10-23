package log

import "io"

// Options options
type Options struct {
	dateFormat string
	writers    []io.Writer
}

// Option Options function
type Option func(*Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		dateFormat: "2006/01/02 15:04:05.000",
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// Writer writer
func Writer(writer io.Writer) Option {
	return func(o *Options) {
		o.writers = append(o.writers, writer)
	}
}
