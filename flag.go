package flagx

import (
	"strings"
)

type Flag struct {
	name      string
	usage     string
	def       any // using default value when haven't provide
	require   bool
	immutable bool
	parsed    bool
	value     Value
}

func (f *Flag) getKey() string {
	keys := strings.Split(f.name, ",")
	if len(keys) == 1 {
		return "--" + keys[0]
	}

	return "-" + keys[1] + ", --" + keys[0]
}

func (f *Flag) getValueType() string {
	switch f.value.(type) {
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

func (f *Flagx) append(v Value, name string, def any, opts ...Option) {
	fg := &Flag{name: name, value: v, def: def}
	for _, opt := range opts {
		opt(fg)
	}

	if len(fg.name) == 0 {
		f.report("empty name with: " + fg.getValueType())
	}

	var sname, lname string
	keys := strings.Split(fg.name, ",")
	kl := len(keys)
	invalid := false
	if kl == 1 {
		lname = keys[0]
	} else if kl == 2 {
		if len(keys[1]) != 1 {
			invalid = true
		} else {
			sname = keys[1]
		}
		lname = keys[0]
	} else {
		invalid = true
	}

	if invalid {
		f.report("invalid name syntax: " + fg.name)
	}

	if _, ok := f.sflags[sname]; ok {
		f.report("redefined: -" + sname)
	}

	if _, ok := f.lflags[lname]; ok {
		f.report("redefined: --" + lname)
	}

	f.lflags[lname] = fg
	if len(sname) > 0 {
		f.sflags[sname] = fg
	}
}

func (f *Flagx) report(name string) {
	var msg string
	if f.name == "" {
		msg = f.sprintf("flag %s", name)
	} else {
		msg = f.sprintf("%s flag %s", f.name, name)
	}
	panic(msg)
}
