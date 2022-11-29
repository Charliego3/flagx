package flagx

import (
	"encoding/base64"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
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

var (
	_     = flag.NewFlagSet("tester", flag.PanicOnError)
	value = flag.String("str", "string", "string value")
	xval  = String("xval,xx", "def", WithDescription("flagx based on flag string"))
	_     = String("xval1", "def", WithDescription("1 flagx based on flag string"))
	_     = String("xval2", "def", WithDescription("2 flagx based on flag string"))
	_     = String("xval3", "def", WithDescription("3 flagx based on flag string"))
	_     = String("xval4", "def", WithDescription("4 flagx based on flag string"))
	_     = String("xval5", "def", WithDescription("5 flagx based on flag string"))
	// _     = String("xval1,xx1", "def", WithDescription("1 flagx based on flag string"))
	// _     = String("xval2,xx2", "def", WithDescription("2 flagx based on flag string"))
	// _     = String("xval3,xx3", "def", WithDescription("3 flagx based on flag string"))
	// _     = String("xval4,xx4", "def", WithDescription("4 flagx based on flag string"))
	// _     = String("xval5,xx5", "def", WithDescription("5 flagx based on flag string"))
	_ = String("xvals,xx", "xvalue", WithDescription("xval flag"))
)

func TestList(t *testing.T) {
	flag.CommandLine.Init("nnnn", flag.PanicOnError)
	t.Log(*value)
	t.Log(*xval)
	list := intList{1, 2, 3, 4, 5}
	t.Log(list.String())
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
		newArg("a", "int", *b.int),
		newArg("b", "int8", *b.int8),
		newArg("c", "int16", *b.int16),
		newArg("d", "int32", *b.int32),
		newArg("e", "int64", *b.int64),
		newArg("f", "uint", *b.uint),
		newArg("g", "uint8", *b.uint8),
		newArg("i", "uint16", *b.uint16),
		newArg("j", "uint32", *b.uint32),
		newArg("k", "uint64", *b.uint64),
		newArg("l", "float32", *b.float32),
		newArg("m", "float64", *b.float64),
		newArg("n", "bool", *b.bool),
		newArg("o", "string", *b.string),
		newArg("p", "duration", *b.duration),
		newArg("q", "file", b.file),
	)

	flagx := NewFlagx()
	v := &obj{}
	v.int = flagx.Int("int,a", 0, WithDescription("expect int value"))
	v.int8 = flagx.Int8("int8,b", 0, WithDescription("expect int8 value"))
	v.int16 = flagx.Int16("int16,c", 0, WithDescription("expect int16 value"))
	v.int32 = flagx.Int32("int32,d", 0, WithDescription("expect int32 value"))
	v.int64 = flagx.Int64("int64,e", 0, WithDescription("expect int64 value"))
	v.uint = flagx.Uint("uint,f", 0, WithDescription("expect uint value"))
	v.uint8 = flagx.Uint8("uint8,g", 0, WithDescription("expect uint8 value"))
	v.uint16 = flagx.Uint16("uint16,i", 0, WithDescription("expect uint16 value"))
	v.uint32 = flagx.Uint32("uint32,j", 0, WithDescription("expect uint32 value"))
	v.uint64 = flagx.Uint64("uint64,k", 0, WithDescription("expect uint64 value"))
	v.float32 = flagx.Float32("float32,l", 0, WithDescription("expect float32 value"))
	v.float64 = flagx.Float64("float64,m", 0, WithDescription("expect float64 value"))
	v.bool = flagx.Bool("bool,n", false, WithDescription("expect bool value"))
	v.string = flagx.String("string,o", "", WithDescription("expect string value"))
	v.duration = flagx.Duration("duration,p", time.Millisecond, WithDescription("expect string value"))
	v.file = flagx.File("file,q", nil, WithDescription("expect file value"))
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
	s string
	l string
	v any
}

func (a *arg) randomFlag() []string {
	key := "-"
	if len(strings.TrimSpace(a.s)) == 0 || *random[bool]() {
		key += "-" + a.l
	} else {
		key += a.s
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

func newArg(sname string, lname string, v any) *arg {
	return &arg{s: sname, l: lname, v: v}
}
