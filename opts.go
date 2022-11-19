package flagx

type Option func(*Flag)

func WithDescription(desc string) Option {
	return func(arg *Flag) {
		arg.usage = desc
	}
}

func WithRequire(require bool) Option {
	return func(arg *Flag) {
		arg.require = require
	}
}

func WithUnique(unique bool) Option {
	return func(arg *Flag) {
		arg.immutable = unique
	}
}
