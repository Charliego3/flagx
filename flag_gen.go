package flagx

import (
	"os"
	"time"
)

// Int start ========================================

func (f *Flagx) Int(short rune, long string, def int, opts ...Option) *int {
	v := new(int)
	f.IntVar(v, short, long, def, opts...)
	return v
}

func Int(short rune, long string, def int, opts ...Option) *int {
	return CommandLine.Int(short, long, def, opts...)
}

func (f *Flagx) IntVar(p *int, short rune, long string, def int, opts ...Option) {
	*p = def
	f.append((*intValue)(p), short, long, def, opts...)
}

func IntVar(p *int, short rune, long string, def int, opts ...Option) {
	CommandLine.IntVar(p, short, long, def, opts...)
}

// Int ended ========================================

// Int8 start ========================================

func (f *Flagx) Int8(short rune, long string, def int8, opts ...Option) *int8 {
	v := new(int8)
	f.Int8Var(v, short, long, def, opts...)
	return v
}

func Int8(short rune, long string, def int8, opts ...Option) *int8 {
	return CommandLine.Int8(short, long, def, opts...)
}

func (f *Flagx) Int8Var(p *int8, short rune, long string, def int8, opts ...Option) {
	*p = def
	f.append((*int8Value)(p), short, long, def, opts...)
}

func Int8Var(p *int8, short rune, long string, def int8, opts ...Option) {
	CommandLine.Int8Var(p, short, long, def, opts...)
}

// Int8 ended ========================================

// Int16 start ========================================

func (f *Flagx) Int16(short rune, long string, def int16, opts ...Option) *int16 {
	v := new(int16)
	f.Int16Var(v, short, long, def, opts...)
	return v
}

func Int16(short rune, long string, def int16, opts ...Option) *int16 {
	return CommandLine.Int16(short, long, def, opts...)
}

func (f *Flagx) Int16Var(p *int16, short rune, long string, def int16, opts ...Option) {
	*p = def
	f.append((*int16Value)(p), short, long, def, opts...)
}

func Int16Var(p *int16, short rune, long string, def int16, opts ...Option) {
	CommandLine.Int16Var(p, short, long, def, opts...)
}

// Int16 ended ========================================

// Int32 start ========================================

func (f *Flagx) Int32(short rune, long string, def int32, opts ...Option) *int32 {
	v := new(int32)
	f.Int32Var(v, short, long, def, opts...)
	return v
}

func Int32(short rune, long string, def int32, opts ...Option) *int32 {
	return CommandLine.Int32(short, long, def, opts...)
}

func (f *Flagx) Int32Var(p *int32, short rune, long string, def int32, opts ...Option) {
	*p = def
	f.append((*int32Value)(p), short, long, def, opts...)
}

func Int32Var(p *int32, short rune, long string, def int32, opts ...Option) {
	CommandLine.Int32Var(p, short, long, def, opts...)
}

// Int32 ended ========================================

// Int64 start ========================================

func (f *Flagx) Int64(short rune, long string, def int64, opts ...Option) *int64 {
	v := new(int64)
	f.Int64Var(v, short, long, def, opts...)
	return v
}

func Int64(short rune, long string, def int64, opts ...Option) *int64 {
	return CommandLine.Int64(short, long, def, opts...)
}

func (f *Flagx) Int64Var(p *int64, short rune, long string, def int64, opts ...Option) {
	*p = def
	f.append((*int64Value)(p), short, long, def, opts...)
}

func Int64Var(p *int64, short rune, long string, def int64, opts ...Option) {
	CommandLine.Int64Var(p, short, long, def, opts...)
}

// Int64 ended ========================================

// Uint start ========================================

func (f *Flagx) Uint(short rune, long string, def uint, opts ...Option) *uint {
	v := new(uint)
	f.UintVar(v, short, long, def, opts...)
	return v
}

func Uint(short rune, long string, def uint, opts ...Option) *uint {
	return CommandLine.Uint(short, long, def, opts...)
}

func (f *Flagx) UintVar(p *uint, short rune, long string, def uint, opts ...Option) {
	*p = def
	f.append((*uintValue)(p), short, long, def, opts...)
}

func UintVar(p *uint, short rune, long string, def uint, opts ...Option) {
	CommandLine.UintVar(p, short, long, def, opts...)
}

// Uint ended ========================================

// Uint8 start ========================================

func (f *Flagx) Uint8(short rune, long string, def uint8, opts ...Option) *uint8 {
	v := new(uint8)
	f.Uint8Var(v, short, long, def, opts...)
	return v
}

func Uint8(short rune, long string, def uint8, opts ...Option) *uint8 {
	return CommandLine.Uint8(short, long, def, opts...)
}

func (f *Flagx) Uint8Var(p *uint8, short rune, long string, def uint8, opts ...Option) {
	*p = def
	f.append((*uint8Value)(p), short, long, def, opts...)
}

func Uint8Var(p *uint8, short rune, long string, def uint8, opts ...Option) {
	CommandLine.Uint8Var(p, short, long, def, opts...)
}

// Uint8 ended ========================================

// Uint16 start ========================================

func (f *Flagx) Uint16(short rune, long string, def uint16, opts ...Option) *uint16 {
	v := new(uint16)
	f.Uint16Var(v, short, long, def, opts...)
	return v
}

func Uint16(short rune, long string, def uint16, opts ...Option) *uint16 {
	return CommandLine.Uint16(short, long, def, opts...)
}

func (f *Flagx) Uint16Var(p *uint16, short rune, long string, def uint16, opts ...Option) {
	*p = def
	f.append((*uint16Value)(p), short, long, def, opts...)
}

func Uint16Var(p *uint16, short rune, long string, def uint16, opts ...Option) {
	CommandLine.Uint16Var(p, short, long, def, opts...)
}

// Uint16 ended ========================================

// Uint32 start ========================================

func (f *Flagx) Uint32(short rune, long string, def uint32, opts ...Option) *uint32 {
	v := new(uint32)
	f.Uint32Var(v, short, long, def, opts...)
	return v
}

func Uint32(short rune, long string, def uint32, opts ...Option) *uint32 {
	return CommandLine.Uint32(short, long, def, opts...)
}

func (f *Flagx) Uint32Var(p *uint32, short rune, long string, def uint32, opts ...Option) {
	*p = def
	f.append((*uint32Value)(p), short, long, def, opts...)
}

func Uint32Var(p *uint32, short rune, long string, def uint32, opts ...Option) {
	CommandLine.Uint32Var(p, short, long, def, opts...)
}

// Uint32 ended ========================================

// Uint64 start ========================================

func (f *Flagx) Uint64(short rune, long string, def uint64, opts ...Option) *uint64 {
	v := new(uint64)
	f.Uint64Var(v, short, long, def, opts...)
	return v
}

func Uint64(short rune, long string, def uint64, opts ...Option) *uint64 {
	return CommandLine.Uint64(short, long, def, opts...)
}

func (f *Flagx) Uint64Var(p *uint64, short rune, long string, def uint64, opts ...Option) {
	*p = def
	f.append((*uint64Value)(p), short, long, def, opts...)
}

func Uint64Var(p *uint64, short rune, long string, def uint64, opts ...Option) {
	CommandLine.Uint64Var(p, short, long, def, opts...)
}

// Uint64 ended ========================================

// Float32 start ========================================

func (f *Flagx) Float32(short rune, long string, def float32, opts ...Option) *float32 {
	v := new(float32)
	f.Float32Var(v, short, long, def, opts...)
	return v
}

func Float32(short rune, long string, def float32, opts ...Option) *float32 {
	return CommandLine.Float32(short, long, def, opts...)
}

func (f *Flagx) Float32Var(p *float32, short rune, long string, def float32, opts ...Option) {
	*p = def
	f.append((*float32Value)(p), short, long, def, opts...)
}

func Float32Var(p *float32, short rune, long string, def float32, opts ...Option) {
	CommandLine.Float32Var(p, short, long, def, opts...)
}

// Float32 ended ========================================

// Float64 start ========================================

func (f *Flagx) Float64(short rune, long string, def float64, opts ...Option) *float64 {
	v := new(float64)
	f.Float64Var(v, short, long, def, opts...)
	return v
}

func Float64(short rune, long string, def float64, opts ...Option) *float64 {
	return CommandLine.Float64(short, long, def, opts...)
}

func (f *Flagx) Float64Var(p *float64, short rune, long string, def float64, opts ...Option) {
	*p = def
	f.append((*float64Value)(p), short, long, def, opts...)
}

func Float64Var(p *float64, short rune, long string, def float64, opts ...Option) {
	CommandLine.Float64Var(p, short, long, def, opts...)
}

// Float64 ended ========================================

// Bool start ========================================

func (f *Flagx) Bool(short rune, long string, def bool, opts ...Option) *bool {
	v := new(bool)
	f.BoolVar(v, short, long, def, opts...)
	return v
}

func Bool(short rune, long string, def bool, opts ...Option) *bool {
	return CommandLine.Bool(short, long, def, opts...)
}

func (f *Flagx) BoolVar(p *bool, short rune, long string, def bool, opts ...Option) {
	*p = def
	f.append((*boolValue)(p), short, long, def, opts...)
}

func BoolVar(p *bool, short rune, long string, def bool, opts ...Option) {
	CommandLine.BoolVar(p, short, long, def, opts...)
}

// Bool ended ========================================

// String start ========================================

func (f *Flagx) String(short rune, long string, def string, opts ...Option) *string {
	v := new(string)
	f.StringVar(v, short, long, def, opts...)
	return v
}

func String(short rune, long string, def string, opts ...Option) *string {
	return CommandLine.String(short, long, def, opts...)
}

func (f *Flagx) StringVar(p *string, short rune, long string, def string, opts ...Option) {
	*p = def
	f.append((*stringValue)(p), short, long, def, opts...)
}

func StringVar(p *string, short rune, long string, def string, opts ...Option) {
	CommandLine.StringVar(p, short, long, def, opts...)
}

// String ended ========================================

// Duration start ========================================

func (f *Flagx) Duration(short rune, long string, def time.Duration, opts ...Option) *time.Duration {
	v := new(time.Duration)
	f.DurationVar(v, short, long, def, opts...)
	return v
}

func Duration(short rune, long string, def time.Duration, opts ...Option) *time.Duration {
	return CommandLine.Duration(short, long, def, opts...)
}

func (f *Flagx) DurationVar(p *time.Duration, short rune, long string, def time.Duration, opts ...Option) {
	*p = def
	f.append((*durationValue)(p), short, long, def, opts...)
}

func DurationVar(p *time.Duration, short rune, long string, def time.Duration, opts ...Option) {
	CommandLine.DurationVar(p, short, long, def, opts...)
}

// Duration ended ========================================

// File start ========================================

func (f *Flagx) File(short rune, long string, def *os.File, opts ...Option) *os.File {
	v := new(os.File)
	f.FileVar(v, short, long, def, opts...)
	return v
}

func File(short rune, long string, def *os.File, opts ...Option) *os.File {
	return CommandLine.File(short, long, def, opts...)
}

func (f *Flagx) FileVar(p *os.File, short rune, long string, def *os.File, opts ...Option) {
	if def != nil {
		*p = *def
	}
	f.append((*fileValue)(p), short, long, def, opts...)
}

func FileVar(p *os.File, short rune, long string, def *os.File, opts ...Option) {
	CommandLine.FileVar(p, short, long, def, opts...)
}

// File ended ========================================

// IntList start ========================================

func (f *Flagx) IntList(short rune, long string, def []int, opts ...Option) *[]int {
	v := new([]int)
	f.IntListVar(v, short, long, def, opts...)
	return v
}

func IntList(short rune, long string, def []int, opts ...Option) *[]int {
	return CommandLine.IntList(short, long, def, opts...)
}

func (f *Flagx) IntListVar(p *[]int, short rune, long string, def []int, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*intList)(p), short, long, def, opts...)
}

func IntListVar(p *[]int, short rune, long string, def []int, opts ...Option) {
	CommandLine.IntListVar(p, short, long, def, opts...)
}

// IntList ended ========================================

// Int8List start ========================================

func (f *Flagx) Int8List(short rune, long string, def []int8, opts ...Option) *[]int8 {
	v := new([]int8)
	f.Int8ListVar(v, short, long, def, opts...)
	return v
}

func Int8List(short rune, long string, def []int8, opts ...Option) *[]int8 {
	return CommandLine.Int8List(short, long, def, opts...)
}

func (f *Flagx) Int8ListVar(p *[]int8, short rune, long string, def []int8, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*int8List)(p), short, long, def, opts...)
}

func Int8ListVar(p *[]int8, short rune, long string, def []int8, opts ...Option) {
	CommandLine.Int8ListVar(p, short, long, def, opts...)
}

// Int8List ended ========================================

// Int16List start ========================================

func (f *Flagx) Int16List(short rune, long string, def []int16, opts ...Option) *[]int16 {
	v := new([]int16)
	f.Int16ListVar(v, short, long, def, opts...)
	return v
}

func Int16List(short rune, long string, def []int16, opts ...Option) *[]int16 {
	return CommandLine.Int16List(short, long, def, opts...)
}

func (f *Flagx) Int16ListVar(p *[]int16, short rune, long string, def []int16, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*int16List)(p), short, long, def, opts...)
}

func Int16ListVar(p *[]int16, short rune, long string, def []int16, opts ...Option) {
	CommandLine.Int16ListVar(p, short, long, def, opts...)
}

// Int16List ended ========================================

// Int32List start ========================================

func (f *Flagx) Int32List(short rune, long string, def []int32, opts ...Option) *[]int32 {
	v := new([]int32)
	f.Int32ListVar(v, short, long, def, opts...)
	return v
}

func Int32List(short rune, long string, def []int32, opts ...Option) *[]int32 {
	return CommandLine.Int32List(short, long, def, opts...)
}

func (f *Flagx) Int32ListVar(p *[]int32, short rune, long string, def []int32, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*int32List)(p), short, long, def, opts...)
}

func Int32ListVar(p *[]int32, short rune, long string, def []int32, opts ...Option) {
	CommandLine.Int32ListVar(p, short, long, def, opts...)
}

// Int32List ended ========================================

// Int64List start ========================================

func (f *Flagx) Int64List(short rune, long string, def []int64, opts ...Option) *[]int64 {
	v := new([]int64)
	f.Int64ListVar(v, short, long, def, opts...)
	return v
}

func Int64List(short rune, long string, def []int64, opts ...Option) *[]int64 {
	return CommandLine.Int64List(short, long, def, opts...)
}

func (f *Flagx) Int64ListVar(p *[]int64, short rune, long string, def []int64, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*int64List)(p), short, long, def, opts...)
}

func Int64ListVar(p *[]int64, short rune, long string, def []int64, opts ...Option) {
	CommandLine.Int64ListVar(p, short, long, def, opts...)
}

// Int64List ended ========================================

// UintList start ========================================

func (f *Flagx) UintList(short rune, long string, def []uint, opts ...Option) *[]uint {
	v := new([]uint)
	f.UintListVar(v, short, long, def, opts...)
	return v
}

func UintList(short rune, long string, def []uint, opts ...Option) *[]uint {
	return CommandLine.UintList(short, long, def, opts...)
}

func (f *Flagx) UintListVar(p *[]uint, short rune, long string, def []uint, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*uintList)(p), short, long, def, opts...)
}

func UintListVar(p *[]uint, short rune, long string, def []uint, opts ...Option) {
	CommandLine.UintListVar(p, short, long, def, opts...)
}

// UintList ended ========================================

// Uint8List start ========================================

func (f *Flagx) Uint8List(short rune, long string, def []uint8, opts ...Option) *[]uint8 {
	v := new([]uint8)
	f.Uint8ListVar(v, short, long, def, opts...)
	return v
}

func Uint8List(short rune, long string, def []uint8, opts ...Option) *[]uint8 {
	return CommandLine.Uint8List(short, long, def, opts...)
}

func (f *Flagx) Uint8ListVar(p *[]uint8, short rune, long string, def []uint8, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*uint8List)(p), short, long, def, opts...)
}

func Uint8ListVar(p *[]uint8, short rune, long string, def []uint8, opts ...Option) {
	CommandLine.Uint8ListVar(p, short, long, def, opts...)
}

// Uint8List ended ========================================

// Uint16List start ========================================

func (f *Flagx) Uint16List(short rune, long string, def []uint16, opts ...Option) *[]uint16 {
	v := new([]uint16)
	f.Uint16ListVar(v, short, long, def, opts...)
	return v
}

func Uint16List(short rune, long string, def []uint16, opts ...Option) *[]uint16 {
	return CommandLine.Uint16List(short, long, def, opts...)
}

func (f *Flagx) Uint16ListVar(p *[]uint16, short rune, long string, def []uint16, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*uint16List)(p), short, long, def, opts...)
}

func Uint16ListVar(p *[]uint16, short rune, long string, def []uint16, opts ...Option) {
	CommandLine.Uint16ListVar(p, short, long, def, opts...)
}

// Uint16List ended ========================================

// Uint32List start ========================================

func (f *Flagx) Uint32List(short rune, long string, def []uint32, opts ...Option) *[]uint32 {
	v := new([]uint32)
	f.Uint32ListVar(v, short, long, def, opts...)
	return v
}

func Uint32List(short rune, long string, def []uint32, opts ...Option) *[]uint32 {
	return CommandLine.Uint32List(short, long, def, opts...)
}

func (f *Flagx) Uint32ListVar(p *[]uint32, short rune, long string, def []uint32, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*uint32List)(p), short, long, def, opts...)
}

func Uint32ListVar(p *[]uint32, short rune, long string, def []uint32, opts ...Option) {
	CommandLine.Uint32ListVar(p, short, long, def, opts...)
}

// Uint32List ended ========================================

// Uint64List start ========================================

func (f *Flagx) Uint64List(short rune, long string, def []uint64, opts ...Option) *[]uint64 {
	v := new([]uint64)
	f.Uint64ListVar(v, short, long, def, opts...)
	return v
}

func Uint64List(short rune, long string, def []uint64, opts ...Option) *[]uint64 {
	return CommandLine.Uint64List(short, long, def, opts...)
}

func (f *Flagx) Uint64ListVar(p *[]uint64, short rune, long string, def []uint64, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*uint64List)(p), short, long, def, opts...)
}

func Uint64ListVar(p *[]uint64, short rune, long string, def []uint64, opts ...Option) {
	CommandLine.Uint64ListVar(p, short, long, def, opts...)
}

// Uint64List ended ========================================

// Float32List start ========================================

func (f *Flagx) Float32List(short rune, long string, def []float32, opts ...Option) *[]float32 {
	v := new([]float32)
	f.Float32ListVar(v, short, long, def, opts...)
	return v
}

func Float32List(short rune, long string, def []float32, opts ...Option) *[]float32 {
	return CommandLine.Float32List(short, long, def, opts...)
}

func (f *Flagx) Float32ListVar(p *[]float32, short rune, long string, def []float32, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*float32List)(p), short, long, def, opts...)
}

func Float32ListVar(p *[]float32, short rune, long string, def []float32, opts ...Option) {
	CommandLine.Float32ListVar(p, short, long, def, opts...)
}

// Float32List ended ========================================

// Float64List start ========================================

func (f *Flagx) Float64List(short rune, long string, def []float64, opts ...Option) *[]float64 {
	v := new([]float64)
	f.Float64ListVar(v, short, long, def, opts...)
	return v
}

func Float64List(short rune, long string, def []float64, opts ...Option) *[]float64 {
	return CommandLine.Float64List(short, long, def, opts...)
}

func (f *Flagx) Float64ListVar(p *[]float64, short rune, long string, def []float64, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*float64List)(p), short, long, def, opts...)
}

func Float64ListVar(p *[]float64, short rune, long string, def []float64, opts ...Option) {
	CommandLine.Float64ListVar(p, short, long, def, opts...)
}

// Float64List ended ========================================

// BoolList start ========================================

func (f *Flagx) BoolList(short rune, long string, def []bool, opts ...Option) *[]bool {
	v := new([]bool)
	f.BoolListVar(v, short, long, def, opts...)
	return v
}

func BoolList(short rune, long string, def []bool, opts ...Option) *[]bool {
	return CommandLine.BoolList(short, long, def, opts...)
}

func (f *Flagx) BoolListVar(p *[]bool, short rune, long string, def []bool, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*boolList)(p), short, long, def, opts...)
}

func BoolListVar(p *[]bool, short rune, long string, def []bool, opts ...Option) {
	CommandLine.BoolListVar(p, short, long, def, opts...)
}

// BoolList ended ========================================

// StringList start ========================================

func (f *Flagx) StringList(short rune, long string, def []string, opts ...Option) *[]string {
	v := new([]string)
	f.StringListVar(v, short, long, def, opts...)
	return v
}

func StringList(short rune, long string, def []string, opts ...Option) *[]string {
	return CommandLine.StringList(short, long, def, opts...)
}

func (f *Flagx) StringListVar(p *[]string, short rune, long string, def []string, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*stringList)(p), short, long, def, opts...)
}

func StringListVar(p *[]string, short rune, long string, def []string, opts ...Option) {
	CommandLine.StringListVar(p, short, long, def, opts...)
}

// StringList ended ========================================

// DurationList start ========================================

func (f *Flagx) DurationList(short rune, long string, def []time.Duration, opts ...Option) *[]time.Duration {
	v := new([]time.Duration)
	f.DurationListVar(v, short, long, def, opts...)
	return v
}

func DurationList(short rune, long string, def []time.Duration, opts ...Option) *[]time.Duration {
	return CommandLine.DurationList(short, long, def, opts...)
}

func (f *Flagx) DurationListVar(p *[]time.Duration, short rune, long string, def []time.Duration, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*durationList)(p), short, long, def, opts...)
}

func DurationListVar(p *[]time.Duration, short rune, long string, def []time.Duration, opts ...Option) {
	CommandLine.DurationListVar(p, short, long, def, opts...)
}

// DurationList ended ========================================

// FileList start ========================================

func (f *Flagx) FileList(short rune, long string, def []*os.File, opts ...Option) *[]*os.File {
	v := new([]*os.File)
	f.FileListVar(v, short, long, def, opts...)
	return v
}

func FileList(short rune, long string, def []*os.File, opts ...Option) *[]*os.File {
	return CommandLine.FileList(short, long, def, opts...)
}

func (f *Flagx) FileListVar(p *[]*os.File, short rune, long string, def []*os.File, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*fileList)(p), short, long, def, opts...)
}

func FileListVar(p *[]*os.File, short rune, long string, def []*os.File, opts ...Option) {
	CommandLine.FileListVar(p, short, long, def, opts...)
}

// FileList ended ========================================
