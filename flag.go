package flagx

type Flag struct {
	short     rune
	long      string
	usage     string
	def       any // using default value when haven't provide
	require   bool
	immutable bool
	parsed    bool
	value     Value
}

func (f *Flag) getKey() string {
	key := "-" + string(f.short)
	if len(f.long) > 0 {
		key += ", --" + f.long
	}
	return key
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

func (f *Flagx) append(v Value, short rune, long string, def any, opts ...Option) *Flag {
	arg := &Flag{short: short, long: long, value: v, def: def}
	for _, opt := range opts {
		opt(arg)
	}

	sn := string(short)
	f.check(v, sn, "-")
	f.flags[sn] = arg
	if len(long) > 0 {
		f.check(v, long, "--")
		f.flags[long] = arg
	}
	return arg
}

func (f *Flagx) check(v Value, name, dashes string) {
	_, exists := f.flags[name]
	if _, ok := v.(listValue); !ok && exists {
		f.report(dashes + name)
	}
}

func (f *Flagx) report(name string) {
	var msg string
	if f.name == "" {
		msg = f.sprintf("flag redefined: %s", name)
	} else {
		msg = f.sprintf("%s flag redefined: %s", f.name, name)
	}
	panic(msg)
}
