package flagx

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type Value interface {
	Set(string) error
}

func parseInt(v string, bitSize int, callback func(r int64)) error {
	r, err := strconv.ParseInt(v, 10, bitSize)
	if err != nil {
		return err
	}
	callback(r)
	return nil
}

type intValue int

func (i *intValue) Set(v string) error {
	return parseInt(v, 0, func(r int64) {
		*i = intValue(r)
	})
}

type int8Value int8

func (i *int8Value) Set(v string) error {
	return parseInt(v, 8, func(r int64) {
		*i = int8Value(r)
	})
}

type int16Value int16

func (i *int16Value) Set(v string) error {
	return parseInt(v, 16, func(r int64) {
		*i = int16Value(r)
	})
}

type int32Value int32

func (i *int32Value) Set(v string) error {
	return parseInt(v, 32, func(r int64) {
		*i = int32Value(r)
	})
}

type int64Value int64

func (i *int64Value) Set(v string) error {
	return parseInt(v, 64, func(r int64) {
		*i = int64Value(r)
	})
}

func parseUint(v string, bitSize int, callback func(r uint64)) error {
	r, err := strconv.ParseUint(v, 10, bitSize)
	if err != nil {
		return err
	}
	callback(r)
	return nil
}

type uintValue uint

func (i *uintValue) Set(v string) error {
	return parseUint(v, 0, func(r uint64) {
		*i = uintValue(r)
	})
}

type uint8Value uint8

func (i *uint8Value) Set(v string) error {
	return parseUint(v, 8, func(r uint64) {
		*i = uint8Value(r)
	})
}

type uint16Value uint16

func (i *uint16Value) Set(v string) error {
	return parseUint(v, 16, func(r uint64) {
		*i = uint16Value(r)
	})
}

type uint32Value uint32

func (i *uint32Value) Set(v string) error {
	return parseUint(v, 32, func(r uint64) {
		*i = uint32Value(r)
	})
}

type uint64Value uint64

func (i *uint64Value) Set(v string) error {
	return parseUint(v, 64, func(r uint64) {
		*i = uint64Value(r)
	})
}

func parseFloat(v string, bitSize int, callback func(r float64)) error {
	r, err := strconv.ParseFloat(v, bitSize)
	if err != nil {
		return err
	}
	callback(r)
	return nil
}

type float32Value float32

func (i *float32Value) Set(v string) error {
	return parseFloat(v, 32, func(r float64) {
		*i = float32Value(r)
	})
}

type float64Value float64

func (i *float64Value) Set(v string) error {
	return parseFloat(v, 64, func(r float64) {
		*i = float64Value(r)
	})
}

type boolValue bool

func (b *boolValue) Set(v string) error {
	bl, err := strconv.ParseBool(v)
	if err != nil {
		return err
	}
	*b = boolValue(bl)
	return nil
}

type stringValue string

func (s *stringValue) Set(v string) error {
	*s = stringValue(v)
	return nil
}

type durationValue time.Duration

func (d *durationValue) Set(v string) error {
	r, err := time.ParseDuration(v)
	if err != nil {
		return err
	}
	*d = durationValue(r)
	return nil
}

type fileValue os.File

func (f *fileValue) Set(v string) error {
	file, err := os.Open(v)
	if err != nil {
		return err
	}
	*f = fileValue(*file)
	return nil
}

type listValue interface {
	IsList() bool
}

func parseList(v string, callback func(r string) error) error {
	if strings.Contains(v, ",") {
		for _, s := range strings.Split(v, ",") {
			if err := callback(s); err != nil {
				return err
			}
		}
		return nil
	}

	return callback(v)
}

type intList []int

func (l *intList) Set(v string) error {
	return parseList(v, func(s string) error {
		return parseInt(s, 0, func(r int64) {
			*l = append(*l, int(r))
		})
	})
}

func (l *intList) IsList() bool {
	return true
}

type int8List []int8

func (l *int8List) Set(v string) error {
	return parseList(v, func(s string) error {
		return parseInt(s, 8, func(r int64) {
			*l = append(*l, int8(r))
		})
	})
}

func (l *int8List) IsList() bool {
	return true
}

type int16List []int16

func (l *int16List) Set(v string) error {
	return parseList(v, func(s string) error {
		return parseInt(s, 16, func(r int64) {
			*l = append(*l, int16(r))
		})
	})
}

func (l *int16List) IsList() bool {
	return true
}

type int32List []int32

func (l *int32List) Set(v string) error {
	return parseList(v, func(s string) error {
		return parseInt(s, 32, func(r int64) {
			*l = append(*l, int32(r))
		})
	})
}

func (l *int32List) IsList() bool {
	return true
}

type int64List []int64

func (l *int64List) Set(v string) error {
	return parseList(v, func(s string) error {
		return parseInt(s, 64, func(r int64) {
			*l = append(*l, r)
		})
	})
}

func (l *int64List) IsList() bool {
	return true
}

type uintList []uint

func (l *uintList) Set(v string) error {
	return parseList(v, func(s string) error {
		return parseUint(s, 0, func(r uint64) {
			*l = append(*l, uint(r))
		})
	})
}

func (l *uintList) IsList() bool {
	return true
}

type uint8List []uint8

func (l *uint8List) Set(v string) error {
	return parseList(v, func(s string) error {
		return parseUint(s, 8, func(r uint64) {
			*l = append(*l, uint8(r))
		})
	})
}

func (l *uint8List) IsList() bool {
	return true
}

type uint16List []uint16

func (l *uint16List) Set(v string) error {
	return parseList(v, func(s string) error {
		return parseUint(s, 16, func(r uint64) {
			*l = append(*l, uint16(r))
		})
	})
}

func (l *uint16List) IsList() bool {
	return true
}

type uint32List []uint32

func (l *uint32List) Set(v string) error {
	return parseList(v, func(s string) error {
		return parseUint(s, 32, func(r uint64) {
			*l = append(*l, uint32(r))
		})
	})
}

func (l *uint32List) IsList() bool {
	return true
}

type uint64List []uint64

func (l *uint64List) Set(v string) error {
	return parseList(v, func(s string) error {
		return parseUint(s, 64, func(r uint64) {
			*l = append(*l, r)
		})
	})
}

func (l *uint64List) IsList() bool {
	return true
}

type float32List []float32

func (l *float32List) Set(v string) error {
	return parseList(v, func(s string) error {
		return parseFloat(s, 32, func(r float64) {
			*l = append(*l, float32(r))
		})
	})
}

func (l *float32List) IsList() bool {
	return true
}

type float64List []float64

func (l *float64List) Set(v string) error {
	return parseList(v, func(s string) error {
		return parseFloat(s, 64, func(r float64) {
			*l = append(*l, r)
		})
	})
}

func (l *float64List) IsList() bool {
	return true
}

type boolList []bool

func (l *boolList) Set(v string) error {
	return parseList(v, func(s string) error {
		b, err := strconv.ParseBool(s)
		if err != nil {
			return err
		}
		*l = append(*l, b)
		return nil
	})
}

func (l *boolList) IsList() bool {
	return true
}

type stringList []string

func (l *stringList) Set(v string) error {
	return parseList(v, func(s string) error {
		*l = append(*l, s)
		return nil
	})
}

func (l *stringList) IsList() bool {
	return true
}

type durationList []time.Duration

func (l *durationList) Set(v string) error {
	return parseList(v, func(s string) error {
		d, err := time.ParseDuration(s)
		if err != nil {
			return err
		}
		*l = append(*l, d)
		return nil
	})
}

func (l *durationList) IsList() bool {
	return true
}

type fileList []*os.File

func (l *fileList) Set(v string) error {
	return parseList(v, func(s string) error {
		file, err := os.Open(s)
		if err != nil {
			return err
		}
		*l = append(*l, file)
		return nil
	})
}

func (l *fileList) IsList() bool {
	return true
}
