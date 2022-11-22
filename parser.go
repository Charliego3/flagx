package flagx

import (
	"errors"
	"fmt"
	"strings"
	"text/tabwriter"
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

	if name == "help" || name == "h" {
		f.Usage()
		return false, nil
	}

	fg, err = f.getFlag(original, name, long)
	if err != nil {
		return false, err
	}

	_, isList := fg.value.(listValue)
	if fg.immutable && fg.parsed && !isList {
		return false, f.failf("flag multiple: %s", fg.getKey())
	}

	_, isBool := fg.value.(*boolValue)
	if idx <= 0 { // eg: --name [value], bool can be `--name` only
		vo, vv, _ := f.nextArg()
		if isBool && len(vo) > 0 && vo[0] == '-' {
			value = "true"
			f.args = append([]string{vo}, f.args...)
		} else if len(vv) > 0 && vv[0] == '-' {
			return false, f.failf("flag needs an argument: %s", fg.getKey())
		} else {
			value = vv
		}
	}

	if err = fg.value.Set(value); err != nil {
		return false, f.failf("invalid value %q for flag %s: %v", value, fg.getKey(), err)
	}
	fg.parsed = true
	return true, nil
}

func (f *Flagx) getFlag(original, name string, long bool) (*Flag, error) {
	var fg *Flag
	if long {
		fg = f.lflags[name]
	} else {
		fg = f.sflags[name]
	}
	if fg == nil {
		return nil, noDeclaredErr(f.failf("flag not declared: %s", original))
	}
	return fg, nil
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
	msg := f.sprintf(format+"\n", v...)
	f.Usage()
	return errors.New(msg)
}

func (f *Flagx) sprintf(format string, v ...any) string {
	msg := fmt.Sprintf(format, v...)
	_, _ = fmt.Fprintln(f.Output(), msg)
	return msg
}

func (f *Flagx) defaultUsage() {
	f.sprintf("Usage:")
	if f.name != "" {
		f.sprintf("    %s [flags]...\n", f.name)
	}

	f.sprintf("Flags:")
	w := tabwriter.NewWriter(f.Output(), 10, 0, 3, ' ', 0)
	for _, fg := range f.lflags {
		key := fg.getKey()
		key = "    " + key
		valueType := fg.getValueType()
		if _, ok := fg.value.(*boolValue); ok {
			valueType = "[" + valueType + "]"
		}
		_, _ = fmt.Fprintln(w, key, valueType, "\t", fg.usage)
	}
	_ = w.Flush()
}
