package flagx

import (
	"flag"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Flag struct {
	lname     string
	sname     string
	defVal    string
	usage     string // options
	require   bool   // options
	immutable bool   // options
	parsed    bool
	value     flag.Value
}

func (fg *Flag) isHelp() bool {
	return fg.lname == "help" && fg.sname == "h"
}

func (fg *Flag) getType() string {
	if fg.isHelp() {
		return ""
	}

	switch fg.value.(type) {
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
	switch fg.value.(type) {
	case *stringValue, *durationValue, *fileValue:
		return "\"" + fg.defVal + "\""
	default:
		return fg.defVal
	}
}

func (f *Flagx) addHelp() {
	f.append(new(stringValue), "help,h", WithDescription("Help for "+f.name))
}

func (f *Flagx) append(v flag.Value, name string, opts ...Option) {
	fg := &Flag{value: v, defVal: v.String()}
	for _, opt := range opts {
		opt(fg)
	}

	keys := strings.Split(name, ",")
	keyn := len(keys)
	if keyn > 2 || keyn == 0 {
		f.badSyntax(name, fg.usage)
	} else if keyn == 2 {
		fg.lname = keys[0]
		fg.sname = keys[1]
	} else if len(keys[0]) == 1 {
		fg.sname = keys[0]
	} else {
		fg.lname = keys[0]
	}

	var short bool
	defer func() {
		err := recover()
		if err == nil {
			return
		}
		if strings.Contains(err.(string), "redefined") { // redefined
			var msg string
			if short {
				msg = fg.getShortName()
			} else {
				msg = fg.getLongName()
			}
			msg = "[" + msg + "] " + fg.usage
			def := v.String()
			if len(def) > 0 {
				msg += ` (default: ` + fg.getDef() + `)`
			}
			f.report("redefined:", msg)
		} else { // bad syntax
			f.badSyntax("["+name+"]", fg.usage)
		}
	}()

	if len(fg.lname) > 0 {
		f.set.Var(v, fg.lname, fg.usage)
	}
	if len(fg.sname) > 0 {
		short = true
		f.set.Var(v, fg.sname, fg.usage)
	}

	f.flags = append(f.flags, fg)
}

func (fg *Flag) getShortName() string {
	if len(fg.sname) == 0 {
		return ""
	}
	return "-" + fg.sname
}

func (fg *Flag) getLongName() string {
	if len(fg.lname) == 0 {
		return ""
	}
	return "--" + fg.lname
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
	namen := len(f.name)
	n := namen + len(msgs) + 1
	for i := 0; i < len(msgs); i++ {
		n += len(msgs[i])
	}

	var b strings.Builder
	b.Grow(n)
	if namen > 0 {
		b.WriteString(f.name)
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
