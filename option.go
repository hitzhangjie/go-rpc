package gorpc

type options struct {
	configfile string
}

// Option gorpc.AddServerModule optionns
type Option func(*options)

// WithConfigfile specify config path
func WithConfig(fpath string) Option {

	return func(opts *options) {
		opts.configfile = fpath
	}
}