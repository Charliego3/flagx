package flagx

import (
	"errors"
	"fmt"
	"strings"
	"text/tabwriter"
)

func (f *Flagx) parseOne() (bool, error) {
	if len(f.args) == 0 {
		return false, nil
	}

	original, name := f.nextArg()
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

	fg, err = f.getFlag(original, name)
	if err != nil {
		return false, err
	}

	_, isList := fg.value.(listValue)
	if fg.immutable && fg.parsed && !isList {
		return false, f.failf("flag multiple: %s", fg.getKey())
	}

	_, isBool := fg.value.(*boolValue)
	if idx <= 0 { // eg: --name [value], bool can be `--name` only
		vo, vv := f.nextArg()
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

func (f *Flagx) getFlag(original, name string) (*Flag, error) {
	arg := f.flags[name]
	if arg == nil {
		return nil, f.failf("flag not declared: %s", original)
	}
	return arg, nil
}

func (f *Flagx) nextArg() (string, string) {
	if len(f.args) == 0 {
		return "", ""
	}

	arg := f.args[0]
	if len(arg) == 0 {
		f.args = f.args[1:]
		return f.nextArg()
	}

	numMinus := 0
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
		}
	}
	f.args = f.args[1:]
	return arg, arg[numMinus:]
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
	ignore := make(map[string]struct{})
	w := tabwriter.NewWriter(f.Output(), 10, 0, 3, ' ', 0)
	for _, fg := range f.flags {
		key := fg.getKey()
		if _, ok := ignore[key]; ok {
			continue
		}

		ignore[key] = struct{}{}
		key = "    " + key
		valueType := fg.getValueType()
		if _, ok := fg.value.(*boolValue); ok {
			valueType = "[" + valueType + "]"
		}
		_, _ = fmt.Fprintln(w, key, valueType, "\t", fg.usage)
	}
	_ = w.Flush()
}
