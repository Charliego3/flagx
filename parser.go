package flagx

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/charmbracelet/lipgloss"
)

type noDeclaredErr error

func (f *Flagx) parseOne() (bool, error) {
	if len(f.args) == 0 {
		return false, nil
	}

	original, name, long := f.nextArg()
	if original[0] != '-' || len(name) == 0 ||
		name[0] == '-' || name[0] == '=' {
		return false, f.failf("bad flag syntax: %s", original)
	}

	var fg *Flag
	var value string
	var err error
	idx := strings.Index(name, "=")
	if idx > 0 { // eg: --name=value
		value = name[idx+1:]
		name = name[:idx]
	}

	for _, g := range f.flags {
		if long && g.Lname == name {
			fg = g
			break
		} else if g.Sname == name {
			fg = g
			break
		}
	}

	if fg == nil {
		return false, noDeclaredErr(f.failf("flag not declared: %s", original))
	}

	if fg.isHelp() {
		f.Usage()
		patchOSExit(0)
	}

	_, isList := fg.Value.(listValue)
	if fg.immutable && fg.Parsed && !isList {
		return false, f.failf("flag multiple: %s", fg.showFlag())
	}

	// instead of providing value via =
	// eg: --name [value], bool can be `--name` only
	if idx <= 0 {
		vo, vv, _ := f.nextArg()
		if len(vv) > 0 && vo[0] == '-' {
			if _, ok := fg.Value.(*boolValue); ok {
				value = "true"
				f.args = append([]string{vo}, f.args...)
			} else {
				return false, f.failf("flag needs an argument: %s", fg.showFlag())
			}
		} else {
			value = vv
		}
	}

	if err = fg.Value.Set(value); err != nil {
		return false, f.failf("invalid value %q for flag %s: %v", value, fg.showFlag(), err)
	}
	fg.Parsed = true
	return true, nil
}

func (f *Flagx) nextArg() (string, string, bool) {
	if len(f.args) == 0 {
		return "", "", false
	}

	arg := f.args[0]
	if len(arg) == 0 {
		f.args = f.args[1:]
		return f.nextArg()
	}

	numMinus := 0
	var long bool
	if arg[0] == '-' {
		numMinus++
		if len(arg) == 1 {
			f.args = f.args[1:]
			return f.nextArg()
		}

		if arg[1] == '-' {
			numMinus++
			if len(arg) == 2 {
				f.args = f.args[1:]
				return f.nextArg()
			}
			long = true
		}
	}
	f.args = f.args[1:]
	return arg, arg[numMinus:], long
}

func (f *Flagx) failf(format string, v ...any) error {
	msg := fmt.Sprintf(format, v...)
	return errors.New(msg)
}

func (f *Flagx) println(msgs ...string) {
	_, _ = fmt.Fprintln(f.Output(), strings.Join(msgs, ""))
}

// defaultUsage print the Usage to console
func (f *Flagx) defaultUsage() {
	padding := lipgloss.NewStyle().PaddingLeft(4).Bold(true)
	usageStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#ffa400"))
	nameStyle := padding.Copy().Underline(true).Foreground(lipgloss.Color("#0aa344"))
	optsStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#75878a"))
	flagStyle := usageStyle.Copy().Foreground(lipgloss.Color("#00F5FF"))
	snameStyle := padding.Copy().Foreground(lipgloss.Color("#0c8918"))
	lnameStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#4682B4"))
	defStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#7D26CD"))
	typStyle := lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#CFCFCF", Dark: "#4F4F4F"})
	descStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#808080"))
	name := f.set.Name()
	if len(name) == 0 {
		name = os.Args[0]
	}
	f.println(usageStyle.Render("Usage:"))
	f.println(nameStyle.Render(name), optsStyle.Render(" [flags]...\n"))
	f.println(flagStyle.Render("Flags:"))
	w := tabwriter.NewWriter(f.Output(), 10, 0, 0, ' ', 0)

	for _, fg := range f.flags {
		var name, usage string
		sname := fg.getShortName()
		lname := fg.getLongName()
		typ := fg.getType()
		if len(typ) > 0 {
			typ = "<" + typ + ">"
		}

		name = snameStyle.Render(sname)
		if len(lname) > 0 {
			if len(sname) > 0 {
				name += typStyle.Render(", ")
			} else {
				name += typStyle.String()
			}
			name += lnameStyle.Render(lname)
		}
		usage = fg.Usage
		if len(fg.DefVal) > 0 {
			usage += defStyle.Render(fg.getDef())
		}
		_, _ = fmt.Fprintln(w, name, "\t", typStyle.Render(typ), "\t", descStyle.Render("\t\t"+usage))
	}
	_ = w.Flush()
}
