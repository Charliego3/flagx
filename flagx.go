package flagx

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

const (
	ContinueOnError   flag.ErrorHandling = 0b1
	SkipNoDeclared    flag.ErrorHandling = 0b10
	OverrideRedefined flag.ErrorHandling = 0b100
)

type Flagx struct {
	set    *flag.FlagSet
	desc   string
	flags  []*Flag
	args   []string
	output io.Writer

	UsageFn  func()
	handling flag.ErrorHandling
}

func init() {
	flag.CommandLine.Init(os.Args[0], flag.ExitOnError)
}

var (
	CommandLine = newfgx("", flag.CommandLine)
	patchOSExit = os.Exit
)

func NewNamedFlagx(name, description string) *Flagx {
	return newfgx(description, flag.NewFlagSet(name, flag.ExitOnError))
}

func newfgx(description string, set *flag.FlagSet) *Flagx {
	set.SetOutput(io.Discard)
	lipgloss.SetColorProfile(termenv.TrueColor)
	f := &Flagx{
		desc: description,
		set:  set,
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

func (f *Flagx) SetErrorHandling(h flag.ErrorHandling) {
	f.handling = h
}

func SetErrorHandling(h flag.ErrorHandling) {
	CommandLine.handling = h
}

func (f *Flagx) ErrorHandling() flag.ErrorHandling {
	return f.handling
}

func ErrorHandling() flag.ErrorHandling {
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

func (f *Flagx) Arg(i int) string {
	if i < 0 || i >= len(f.args) {
		return ""
	}
	return f.args[i]
}

func Arg(i int) string {
	return CommandLine.Arg(i)
}

func (f *Flagx) Args() []string {
	return f.args
}

func Args() []string {
	return CommandLine.args
}

func (f *Flagx) Name() string {
	return f.set.Name()
}

func Name() string {
	return CommandLine.Name()
}

func (f *Flagx) Lookup(name string) *flag.Flag {
	return f.set.Lookup(name)
}

func Lookup(name string) *flag.Flag {
	return CommandLine.Lookup(name)
}

func (f *Flagx) NArg() int {
	return len(f.args)
}

func NArg() int {
	return CommandLine.NArg()
}

func (f *Flagx) NFlag() int {
	return len(f.flags)
}

func NFlag() int {
	return CommandLine.NFlag()
}

func (f *Flagx) Visit(fn func(*Flag)) {
	for _, fg := range f.flags {
		fn(fg)
	}
}

func Visit(fn func(*Flag)) {
	CommandLine.Visit(fn)
}

func (f *Flagx) Set(name, value string) error {
	var fg *Flag
	for _, g := range f.flags {
		if g.Lname == name || g.Sname == name {
			fg = g
			break
		}
	}

	if fg == nil {
		return fmt.Errorf("no such flag -%v", name)
	}

	_, isList := fg.Value.(listValue)
	if fg.immutable && fg.Parsed && !isList {
		return fmt.Errorf("flag is immutable, can not be modtify")
	}

	return f.set.Set(name, value)
}

func Set(name, value string) {
	CommandLine.Set(name, value)
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

		f.report(err.Error())
		patchOSExit(0)
		break
	}

	for _, fg := range f.flags {
		if fg.require && !fg.Parsed {
			f.report("flag is required: ", fg.showFlag())
			patchOSExit(0)
		}
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
