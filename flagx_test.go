package flagx

import (
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/suite"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

type obj struct {
	int   *int
	int8  *int8
	int16 *int16
	int32 *int32
	int64 *int64

	uint   *uint
	uint8  *uint8
	uint16 *uint16
	uint32 *uint32
	uint64 *uint64

	float32  *float32
	float64  *float64
	bool     *bool
	string   *string
	duration *time.Duration
	file     *os.File
}

type BasicSuite struct {
	suite.Suite
	*obj
}

func (b *BasicSuite) SetupTest() {
	registerPatch(b.T())
	rand.Seed(time.Now().UnixMilli())
	os.Args = []string{"flagx_tester"}
	b.obj = randomObj()
}

func (b *BasicSuite) TestBasic() {
	fakeArgs(
		newArg('a', "int", *b.int),
		newArg('b', "int8", *b.int8),
		newArg('c', "int16", *b.int16),
		newArg('d', "int32", *b.int32),
		newArg('e', "int64", *b.int64),
		newArg('f', "uint", *b.uint),
		newArg('g', "uint8", *b.uint8),
		newArg('i', "uint16", *b.uint16),
		newArg('j', "uint32", *b.uint32),
		newArg('k', "uint64", *b.uint64),
		newArg('l', "float32", *b.float32),
		newArg('m', "float64", *b.float64),
		newArg('n', "bool", *b.bool),
		newArg('o', "string", *b.string),
		newArg('p', "duration", *b.duration),
		newArg('q', "file", b.file),
	)

	flagx := NewFlagx()
	v := &obj{}
	v.int = flagx.Int('a', "int", 0, WithDescription("expect int value"))
	v.int8 = flagx.Int8('b', "int8", 0, WithDescription("expect int8 value"))
	v.int16 = flagx.Int16('c', "int16", 0, WithDescription("expect int16 value"))
	v.int32 = flagx.Int32('d', "int32", 0, WithDescription("expect int32 value"))
	v.int64 = flagx.Int64('e', "int64", 0, WithDescription("expect int64 value"))
	v.uint = flagx.Uint('f', "uint", 0, WithDescription("expect uint value"))
	v.uint8 = flagx.Uint8('g', "uint8", 0, WithDescription("expect uint8 value"))
	v.uint16 = flagx.Uint16('i', "uint16", 0, WithDescription("expect uint16 value"))
	v.uint32 = flagx.Uint32('j', "uint32", 0, WithDescription("expect uint32 value"))
	v.uint64 = flagx.Uint64('k', "uint64", 0, WithDescription("expect uint64 value"))
	v.float32 = flagx.Float32('l', "float32", 0, WithDescription("expect float32 value"))
	v.float64 = flagx.Float64('m', "float64", 0, WithDescription("expect float64 value"))
	v.bool = flagx.Bool('n', "bool", false, WithDescription("expect bool value"))
	v.string = flagx.String('o', "string", "", WithDescription("expect string value"))
	v.duration = flagx.Duration('p', "duration", time.Millisecond, WithDescription("expect string value"))
	v.file = flagx.File('q', "file", nil, WithDescription("expect file value"))
	err := flagx.Parse()
	b.NoError(err)

	b.Equal(*v.int, *b.int, "int value not equal")
	b.Equal(*v.int8, *b.int8, "int8 value not equal")
	b.Equal(*v.int16, *b.int16, "int16 value not equal")
	b.Equal(*v.int32, *b.int32, "int32 value not equal")
	b.Equal(*v.int64, *b.int64, "int64 value not equal")
	b.Equal(*v.uint, *b.uint, "uint value not equal")
	b.Equal(*v.uint8, *b.uint8, "uint8 value not equal")
	b.Equal(*v.uint16, *b.uint16, "uint16 value not equal")
	b.Equal(*v.uint32, *b.uint32, "uint32 value not equal")
	b.Equal(*v.uint64, *b.uint64, "uint64 value not equal")
	b.Equal(*v.float32, *b.float32, "float32 value not equal")
	b.Equal(*v.float64, *b.float64, "float64 value not equal")
	b.Equal(*v.bool, *b.bool, "bool value not equal")
	b.Equal(*v.string, *b.string, "string value not equal")
	b.Equal(*v.duration, *b.duration, "duration value not equal")
	b.NotNil(*v.file, "flag file is nil")
	b.Equal((*v.file).Name(), (*b.file).Name(), "file value not equal")
}

func (b *BasicSuite) TestVar() {

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(BasicSuite))
}

func fakeArgs(args ...*arg) {
	for _, a := range args {
		os.Args = append(os.Args, a.randomFlag()...)
	}
}

func randomObj() *obj {
	obj := new(obj)
	obj.int = random[int]()
	obj.int8 = random[int8]()
	obj.int16 = random[int16]()
	obj.int32 = random[int32]()
	obj.int64 = random[int64]()
	obj.uint = random[uint]()
	obj.uint8 = random[uint8]()
	obj.uint16 = random[uint16]()
	obj.uint32 = random[uint32]()
	obj.uint64 = random[uint64]()
	obj.float32 = random[float32]()
	obj.float64 = random[float64]()
	obj.bool = random[bool]()
	obj.string = random[string]()
	obj.duration = random[time.Duration]()
	obj.file = random[os.File]()
	return obj
}

func random[T any]() *T {
	var t T
	switch v := (any)(&t).(type) {
	case *int:
		*v = rand.Int()
	case *int8:
		*v = int8(rand.Intn(math.MaxInt8))
	case *int16:
		*v = int16(rand.Intn(math.MaxInt16))
	case *int32:
		*v = rand.Int31()
	case *int64:
		*v = rand.Int63()
	case *uint:
		*v = uint(rand.Int())
	case *uint8:
		*v = uint8(rand.Intn(math.MaxUint8))
	case *uint16:
		*v = uint16(rand.Intn(math.MaxInt16))
	case *uint32:
		*v = rand.Uint32()
	case *uint64:
		*v = rand.Uint64()
	case *float32:
		*v = rand.Float32()
	case *float64:
		*v = rand.Float64()
	case *bool:
		*v = rand.Float64() > 0.3
	case *string:
		buf := make([]byte, 16)
		_, _ = rand.Read(buf)
		*v = base64.StdEncoding.EncodeToString(buf)
	case *time.Duration:
		d, err := time.ParseDuration(strconv.Itoa(*random[int]()) + "s")
		if err != nil {
			d = time.Second
		}
		*v = d
	case *os.File:
		dirs, err := os.ReadDir(".")
		if err != nil {
			return nil
		}

		var entries []os.DirEntry
		for _, de := range dirs {
			if de.IsDir() || strings.HasPrefix(de.Name(), ".") {
				continue
			}
			entries = append(entries, de)
		}

		if len(entries) > 0 {
			entry := entries[rand.Intn(len(entries))]
			file, err := os.Open(entry.Name())
			if err != nil {
				panic(err)
			}
			*v = *file
		}
	}
	return &t
}

func registerPatch(t *testing.T) {
	patchOSExit = func(code int) {
		t.Log("exited")
		t.FailNow()
	}
}

type arg struct {
	s rune
	l string
	v any
}

func (a *arg) randomFlag() []string {
	key := "-"
	if *random[bool]() {
		key += "-" + a.l
	} else {
		key += string(a.s)
	}

	if d, ok := a.v.(bool); ok && d && *random[bool]() {
		return []string{key}
	}

	var v string
	if f, ok := a.v.(*os.File); ok {
		v = f.Name()
	} else {
		v = fmt.Sprintf("%v", a.v)
	}
	if *random[bool]() {
		key += "=" + v
		return []string{key}
	} else {
		return []string{key, v}
	}
}

func newArg(s rune, l string, v any) *arg {
	return &arg{s: s, l: l, v: v}
}
