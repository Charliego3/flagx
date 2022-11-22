package flagx

import (
	"io"
	"os"
)

const (
	ContinueOnError uint8 = 0b1
	SkipNoDeclared  uint8 = 0b10
	ClearAfterParse uint8 = 0b100
)

type Flagx struct {
	name   string
	desc   string
	sflags map[string]*Flag
	lflags map[string]*Flag
	args   []string
	output io.Writer

	UsageFn  func()
	handling uint8
}

var (
	CommandLine = NewFlagx()
	patchOSExit = os.Exit
)

func NewNamedFlagx(name, description string) *Flagx {
	if len(name) == 0 {
		name = os.Args[0]
	}
	f := &Flagx{
		name:   name,
		desc:   description,
		sflags: make(map[string]*Flag),
		lflags: make(map[string]*Flag),
	}
	f.addHelp()
	return f
}

func NewFlagx() *Flagx {
	return NewNamedFlagx(os.Args[0], "")
}

func (f *Flagx) SetOutput(w io.Writer) {
	f.output = w
}

func (f *Flagx) SetErrorHandling(h uint8) {
	f.handling = h
}

func SetErrorHandling(h uint8) {
	CommandLine.handling = h
}

func (f *Flagx) GetErrorHandling() uint8 {
	return f.handling
}

func GetErrorHandling() uint8 {
	return CommandLine.handling
}

func SetOutput(w io.Writer) {
	CommandLine.output = w
}

func (f *Flagx) Output() io.Writer {
	if f.output == nil {
		return os.Stderr
	}
	return f.output
}

func Output() io.Writer {
	return CommandLine.Output()
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
		} else if f.handling&ContinueOnError == ContinueOnError {
			continue
		} else if _, ok := err.(noDeclaredErr); ok &&
			f.handling&SkipNoDeclared == SkipNoDeclared {
			continue
		}

		patchOSExit(0)
		break
	}

	for _, fg := range f.lflags {
		if fg.require && !fg.parsed {
			_ = f.failf("flag is required: %s %s", fg.getKey(), fg.getValueType())
			patchOSExit(0)
		}
	}
	if f.handling&ClearAfterParse == ClearAfterParse {
		f.lflags = make(map[string]*Flag)
		f.sflags = make(map[string]*Flag)
		f.addHelp()
	}
	return nil
}

func Parse() error {
	return CommandLine.Parse()
}

func (f *Flagx) MustParse() {
	err := f.Parse()
	if err != nil {
		panic(err)
	}
}

func MustParse() {
	CommandLine.MustParse()
}

func (f *Flagx) Usage() {
	if f.UsageFn == nil {
		f.defaultUsage()
	} else {
		f.UsageFn()
	}
}

func Usage() {
	CommandLine.Usage()
}
