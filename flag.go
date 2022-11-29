package flagx

import (
	"flag"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Flag struct {
	Lname  string
	Sname  string
	DefVal string
	Value  flag.Value
	Parsed bool
	Usage  string // options

	require   bool // options
	immutable bool // options
}

func (fg *Flag) isHelp() bool {
	return fg.Lname == "help" && fg.Sname == "h"
}

func (fg *Flag) getType() string {
	if fg.isHelp() {
		return ""
	}

	switch fg.Value.(type) {
	case *intValue:
		return "int"
	case *int8Value:
		return "int8"
	case *int16Value:
		return "int16"
	case *int32Value:
		return "int32"
	case *int64Value:
		return "int64"
	case *uintValue:
		return "uint"
	case *uint8Value:
		return "uint8"
	case *uint16Value:
		return "uint16"
	case *uint32Value:
		return "uint32"
	case *uint64Value:
		return "uint64"
	case *float32Value:
		return "float32"
	case *float64Value:
		return "float64"
	case *boolValue:
		return "bool"
	case *stringValue:
		return "string"
	case *durationValue:
		return "duration"
	case *fileValue:
		return "file"
	case *funcValue:
		return "func"
	case *intList:
		return "[]int"
	case *int8List:
		return "[]int8"
	case *int16List:
		return "[]int16"
	case *int32List:
		return "[]int32"
	case *int64List:
		return "[]int64"
	case *uintList:
		return "[]uint"
	case *uint8List:
		return "[]uint8"
	case *uint16List:
		return "[]uint16"
	case *uint32List:
		return "[]uint32"
	case *uint64List:
		return "[]uint64"
	case *float32List:
		return "[]float32"
	case *float64List:
		return "[]float64"
	case *boolList:
		return "[]bool"
	case *stringList:
		return "[]string"
	case *durationList:
		return "[]duration"
	case *fileList:
		return "[]file"
	default:
		return ""
	}
}

func (fg *Flag) getDef() string {
	if len(fg.DefVal) == 0 {
		return ""
	}

	def := fg.DefVal
	switch fg.Value.(type) {
	case *stringValue, *durationValue, *fileValue:
		def = `"` + def + `"`
	}
	return ` (default: ` + def + `)`
}

func (f *Flagx) addHelp() {
	f.append(new(stringValue), "help,h", WithDescription("Help for "+f.set.Name()))
}

func (f *Flagx) append(v flag.Value, name string, opts ...Option) {
	fg := &Flag{Value: v, DefVal: v.String()}
	for _, opt := range opts {
		opt(fg)
	}

	keys := strings.Split(name, ",")
	keyn := len(keys)
	if keyn > 2 || keyn == 0 {
		f.badSyntax(name, fg.Usage)
	} else if keyn == 2 {
		fg.Lname = keys[0]
		fg.Sname = keys[1]
	} else if len(keys[0]) == 1 {
		fg.Sname = keys[0]
	} else {
		fg.Lname = keys[0]
	}

	var short bool
	defer func() {
		err := recover()
		if err == nil {
			return
		}
		if strings.Contains(err.(string), "redefined") { // redefined
			if f.handling&OverrideRedefined == OverrideRedefined {
				lffg := f.Lookup(fg.Lname)
				if lffg != nil {
					lffg.Name = fg.Lname
					lffg.Usage = fg.Usage
					lffg.DefValue = fg.DefVal
					lffg.Value = v
				}

				sffg := f.Lookup(fg.Sname)
				if sffg != nil {
					sffg.Name = fg.Sname
					sffg.Usage = fg.Usage
					sffg.DefValue = fg.DefVal
					sffg.Value = v
				}

				if lffg != nil || sffg != nil {
					for _, g := range f.flags {
						if g.Lname == fg.Lname || g.Sname == fg.Sname {
							g.Value = v
						}
					}
				}
				return
			}

			var msg string
			if short {
				msg = fg.getShortName()
			} else {
				msg = fg.getLongName()
			}
			msg = "[" + msg + "] " + fg.Usage
			msg += fg.getDef()
			f.report("redefined:", msg)
		} else { // bad syntax
			f.badSyntax("["+name+"]", fg.Usage)
		}
	}()

	if len(fg.Lname) > 0 {
		f.set.Var(v, fg.Lname, fg.Usage)
	}
	if len(fg.Sname) > 0 {
		short = true
		f.set.Var(v, fg.Sname, fg.Usage)
	}

	f.flags = append(f.flags, fg)
}

func (fg *Flag) getShortName() string {
	if len(fg.Sname) == 0 {
		return ""
	}
	return "-" + fg.Sname
}

func (fg *Flag) getLongName() string {
	if len(fg.Lname) == 0 {
		return ""
	}
	return "--" + fg.Lname
}

func (fg *Flag) showFlag() string {
	var name string
	sname := fg.getShortName()
	lname := fg.getLongName()
	if len(sname) > 0 {
		name = sname
		if len(lname) > 0 {
			name += ", " + lname
		}
	} else {
		name = lname
	}

	name = name + " " + fg.getType()
	return name
}

func (f *Flagx) badSyntax(msg ...string) {
	f.report(append([]string{"bad syntax:"}, msg...)...)
}

func (f *Flagx) report(msgs ...string) {
	namen := len(f.set.Name())
	n := namen + len(msgs) + 1
	for i := 0; i < len(msgs); i++ {
		n += len(msgs[i])
	}

	var b strings.Builder
	b.Grow(n)
	if namen > 0 {
		b.WriteString(f.set.Name())
	}
	for _, msg := range msgs {
		b.WriteString(" " + msg)
	}
	b.WriteString("\n")

	style := lipgloss.NewStyle().Bold(true).Underline(true).Foreground(lipgloss.Color("#ff4c00"))
	f.println(style.Render(b.String()))
	f.Usage()
	patchOSExit(2)
}
