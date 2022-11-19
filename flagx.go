package flagx

import (
	"io"
	"os"
)

type Flagx struct {
	name   string
	desc   string
	flags  map[string]*Flag
	args   []string
	output io.Writer

	UsageFn func()
}

var (
	CommandLine = NewFlagx()
	patchOSExit = os.Exit
)

func NewNamedFlagx(name, description string) *Flagx {
	if len(name) == 0 {
		name = os.Args[0]
	}
	return &Flagx{
		name: name,
		desc: description,
		flags: map[string]*Flag{
			"h": {short: 'h', long: "help", usage: "Help for " + name},
		},
	}
}

func NewFlagx() *Flagx {
	return NewNamedFlagx(os.Args[0], "")
}

func (f *Flagx) SetOutput(w io.Writer) {
	f.output = w
}

func (f *Flagx) Output() io.Writer {
	if f.output == nil {
		return os.Stderr
	}
	return f.output
}

func (f *Flagx) Parse() error {
	f.args = os.Args[1:]
	for {
		parsed, err := f.parseOne()
		if parsed {
			continue
		}
		if err == nil {
			break
		}

		patchOSExit(0)
		break
	}

	for _, fg := range f.flags {
		if fg.require && !fg.parsed {
			_ = f.failf("flag is required: %s %s", fg.getKey(), fg.getValueType())
			patchOSExit(0)
		}
	}
	return nil
}

func (f *Flagx) Usage() {
	if f.UsageFn == nil {
		f.defaultUsage()
	} else {
		f.UsageFn()
	}
}

func Parse() {
	_ = CommandLine.Parse()
}
