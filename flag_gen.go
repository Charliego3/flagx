package flagx

import (
	"os"
	"time"
)

// Int start ========================================

func (f *Flagx) Int(name string, def int, opts ...Option) *int {
	v := new(int)
	f.IntVar(v, name, def, opts...)
	return v
}

func Int(name string, def int, opts ...Option) *int {
	return CommandLine.Int(name, def, opts...)
}

func (f *Flagx) IntVar(p *int, name string, def int, opts ...Option) {
	*p = def
	f.append((*intValue)(p), name, opts...)
}

func IntVar(p *int, name string, def int, opts ...Option) {
	CommandLine.IntVar(p, name, def, opts...)
}

// Int ended ========================================

// Int8 start ========================================

func (f *Flagx) Int8(name string, def int8, opts ...Option) *int8 {
	v := new(int8)
	f.Int8Var(v, name, def, opts...)
	return v
}

func Int8(name string, def int8, opts ...Option) *int8 {
	return CommandLine.Int8(name, def, opts...)
}

func (f *Flagx) Int8Var(p *int8, name string, def int8, opts ...Option) {
	*p = def
	f.append((*int8Value)(p), name, opts...)
}

func Int8Var(p *int8, name string, def int8, opts ...Option) {
	CommandLine.Int8Var(p, name, def, opts...)
}

// Int8 ended ========================================

// Int16 start ========================================

func (f *Flagx) Int16(name string, def int16, opts ...Option) *int16 {
	v := new(int16)
	f.Int16Var(v, name, def, opts...)
	return v
}

func Int16(name string, def int16, opts ...Option) *int16 {
	return CommandLine.Int16(name, def, opts...)
}

func (f *Flagx) Int16Var(p *int16, name string, def int16, opts ...Option) {
	*p = def
	f.append((*int16Value)(p), name, opts...)
}

func Int16Var(p *int16, name string, def int16, opts ...Option) {
	CommandLine.Int16Var(p, name, def, opts...)
}

// Int16 ended ========================================

// Int32 start ========================================

func (f *Flagx) Int32(name string, def int32, opts ...Option) *int32 {
	v := new(int32)
	f.Int32Var(v, name, def, opts...)
	return v
}

func Int32(name string, def int32, opts ...Option) *int32 {
	return CommandLine.Int32(name, def, opts...)
}

func (f *Flagx) Int32Var(p *int32, name string, def int32, opts ...Option) {
	*p = def
	f.append((*int32Value)(p), name, opts...)
}

func Int32Var(p *int32, name string, def int32, opts ...Option) {
	CommandLine.Int32Var(p, name, def, opts...)
}

// Int32 ended ========================================

// Int64 start ========================================

func (f *Flagx) Int64(name string, def int64, opts ...Option) *int64 {
	v := new(int64)
	f.Int64Var(v, name, def, opts...)
	return v
}

func Int64(name string, def int64, opts ...Option) *int64 {
	return CommandLine.Int64(name, def, opts...)
}

func (f *Flagx) Int64Var(p *int64, name string, def int64, opts ...Option) {
	*p = def
	f.append((*int64Value)(p), name, opts...)
}

func Int64Var(p *int64, name string, def int64, opts ...Option) {
	CommandLine.Int64Var(p, name, def, opts...)
}

// Int64 ended ========================================

// Uint start ========================================

func (f *Flagx) Uint(name string, def uint, opts ...Option) *uint {
	v := new(uint)
	f.UintVar(v, name, def, opts...)
	return v
}

func Uint(name string, def uint, opts ...Option) *uint {
	return CommandLine.Uint(name, def, opts...)
}

func (f *Flagx) UintVar(p *uint, name string, def uint, opts ...Option) {
	*p = def
	f.append((*uintValue)(p), name, opts...)
}

func UintVar(p *uint, name string, def uint, opts ...Option) {
	CommandLine.UintVar(p, name, def, opts...)
}

// Uint ended ========================================

// Uint8 start ========================================

func (f *Flagx) Uint8(name string, def uint8, opts ...Option) *uint8 {
	v := new(uint8)
	f.Uint8Var(v, name, def, opts...)
	return v
}

func Uint8(name string, def uint8, opts ...Option) *uint8 {
	return CommandLine.Uint8(name, def, opts...)
}

func (f *Flagx) Uint8Var(p *uint8, name string, def uint8, opts ...Option) {
	*p = def
	f.append((*uint8Value)(p), name, opts...)
}

func Uint8Var(p *uint8, name string, def uint8, opts ...Option) {
	CommandLine.Uint8Var(p, name, def, opts...)
}

// Uint8 ended ========================================

// Uint16 start ========================================

func (f *Flagx) Uint16(name string, def uint16, opts ...Option) *uint16 {
	v := new(uint16)
	f.Uint16Var(v, name, def, opts...)
	return v
}

func Uint16(name string, def uint16, opts ...Option) *uint16 {
	return CommandLine.Uint16(name, def, opts...)
}

func (f *Flagx) Uint16Var(p *uint16, name string, def uint16, opts ...Option) {
	*p = def
	f.append((*uint16Value)(p), name, opts...)
}

func Uint16Var(p *uint16, name string, def uint16, opts ...Option) {
	CommandLine.Uint16Var(p, name, def, opts...)
}

// Uint16 ended ========================================

// Uint32 start ========================================

func (f *Flagx) Uint32(name string, def uint32, opts ...Option) *uint32 {
	v := new(uint32)
	f.Uint32Var(v, name, def, opts...)
	return v
}

func Uint32(name string, def uint32, opts ...Option) *uint32 {
	return CommandLine.Uint32(name, def, opts...)
}

func (f *Flagx) Uint32Var(p *uint32, name string, def uint32, opts ...Option) {
	*p = def
	f.append((*uint32Value)(p), name, opts...)
}

func Uint32Var(p *uint32, name string, def uint32, opts ...Option) {
	CommandLine.Uint32Var(p, name, def, opts...)
}

// Uint32 ended ========================================

// Uint64 start ========================================

func (f *Flagx) Uint64(name string, def uint64, opts ...Option) *uint64 {
	v := new(uint64)
	f.Uint64Var(v, name, def, opts...)
	return v
}

func Uint64(name string, def uint64, opts ...Option) *uint64 {
	return CommandLine.Uint64(name, def, opts...)
}

func (f *Flagx) Uint64Var(p *uint64, name string, def uint64, opts ...Option) {
	*p = def
	f.append((*uint64Value)(p), name, opts...)
}

func Uint64Var(p *uint64, name string, def uint64, opts ...Option) {
	CommandLine.Uint64Var(p, name, def, opts...)
}

// Uint64 ended ========================================

// Float32 start ========================================

func (f *Flagx) Float32(name string, def float32, opts ...Option) *float32 {
	v := new(float32)
	f.Float32Var(v, name, def, opts...)
	return v
}

func Float32(name string, def float32, opts ...Option) *float32 {
	return CommandLine.Float32(name, def, opts...)
}

func (f *Flagx) Float32Var(p *float32, name string, def float32, opts ...Option) {
	*p = def
	f.append((*float32Value)(p), name, opts...)
}

func Float32Var(p *float32, name string, def float32, opts ...Option) {
	CommandLine.Float32Var(p, name, def, opts...)
}

// Float32 ended ========================================

// Float64 start ========================================

func (f *Flagx) Float64(name string, def float64, opts ...Option) *float64 {
	v := new(float64)
	f.Float64Var(v, name, def, opts...)
	return v
}

func Float64(name string, def float64, opts ...Option) *float64 {
	return CommandLine.Float64(name, def, opts...)
}

func (f *Flagx) Float64Var(p *float64, name string, def float64, opts ...Option) {
	*p = def
	f.append((*float64Value)(p), name, opts...)
}

func Float64Var(p *float64, name string, def float64, opts ...Option) {
	CommandLine.Float64Var(p, name, def, opts...)
}

// Float64 ended ========================================

// Bool start ========================================

func (f *Flagx) Bool(name string, def bool, opts ...Option) *bool {
	v := new(bool)
	f.BoolVar(v, name, def, opts...)
	return v
}

func Bool(name string, def bool, opts ...Option) *bool {
	return CommandLine.Bool(name, def, opts...)
}

func (f *Flagx) BoolVar(p *bool, name string, def bool, opts ...Option) {
	*p = def
	f.append((*boolValue)(p), name, opts...)
}

func BoolVar(p *bool, name string, def bool, opts ...Option) {
	CommandLine.BoolVar(p, name, def, opts...)
}

// Bool ended ========================================

// String start ========================================

func (f *Flagx) String(name string, def string, opts ...Option) *string {
	v := new(string)
	f.StringVar(v, name, def, opts...)
	return v
}

func String(name string, def string, opts ...Option) *string {
	return CommandLine.String(name, def, opts...)
}

func (f *Flagx) StringVar(p *string, name string, def string, opts ...Option) {
	*p = def
	f.append((*stringValue)(p), name, opts...)
}

func StringVar(p *string, name string, def string, opts ...Option) {
	CommandLine.StringVar(p, name, def, opts...)
}

// String ended ========================================

// Duration start ========================================

func (f *Flagx) Duration(name string, def time.Duration, opts ...Option) *time.Duration {
	v := new(time.Duration)
	f.DurationVar(v, name, def, opts...)
	return v
}

func Duration(name string, def time.Duration, opts ...Option) *time.Duration {
	return CommandLine.Duration(name, def, opts...)
}

func (f *Flagx) DurationVar(p *time.Duration, name string, def time.Duration, opts ...Option) {
	*p = def
	f.append((*durationValue)(p), name, opts...)
}

func DurationVar(p *time.Duration, name string, def time.Duration, opts ...Option) {
	CommandLine.DurationVar(p, name, def, opts...)
}

// Duration ended ========================================

// File start ========================================

func (f *Flagx) File(name string, def *os.File, opts ...Option) *os.File {
	// v := new(os.File)
	v := os.NewFile(0, "")
	f.FileVar(v, name, def, opts...)
	return v
}

func File(name string, def *os.File, opts ...Option) *os.File {
	return CommandLine.File(name, def, opts...)
}

func (f *Flagx) FileVar(p *os.File, name string, def *os.File, opts ...Option) {
	if def != nil {
		*p = *def
	}
	f.append((*fileValue)(p), name, opts...)
}

func FileVar(p *os.File, name string, def *os.File, opts ...Option) {
	CommandLine.FileVar(p, name, def, opts...)
}

// File ended ========================================

// IntList start ========================================

func (f *Flagx) IntList(name string, def []int, opts ...Option) *[]int {
	v := new([]int)
	f.IntListVar(v, name, def, opts...)
	return v
}

func IntList(name string, def []int, opts ...Option) *[]int {
	return CommandLine.IntList(name, def, opts...)
}

func (f *Flagx) IntListVar(p *[]int, name string, def []int, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*intList)(p), name, opts...)
}

func IntListVar(p *[]int, name string, def []int, opts ...Option) {
	CommandLine.IntListVar(p, name, def, opts...)
}

// IntList ended ========================================

// Int8List start ========================================

func (f *Flagx) Int8List(name string, def []int8, opts ...Option) *[]int8 {
	v := new([]int8)
	f.Int8ListVar(v, name, def, opts...)
	return v
}

func Int8List(name string, def []int8, opts ...Option) *[]int8 {
	return CommandLine.Int8List(name, def, opts...)
}

func (f *Flagx) Int8ListVar(p *[]int8, name string, def []int8, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*int8List)(p), name, opts...)
}

func Int8ListVar(p *[]int8, name string, def []int8, opts ...Option) {
	CommandLine.Int8ListVar(p, name, def, opts...)
}

// Int8List ended ========================================

// Int16List start ========================================

func (f *Flagx) Int16List(name string, def []int16, opts ...Option) *[]int16 {
	v := new([]int16)
	f.Int16ListVar(v, name, def, opts...)
	return v
}

func Int16List(name string, def []int16, opts ...Option) *[]int16 {
	return CommandLine.Int16List(name, def, opts...)
}

func (f *Flagx) Int16ListVar(p *[]int16, name string, def []int16, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*int16List)(p), name, opts...)
}

func Int16ListVar(p *[]int16, name string, def []int16, opts ...Option) {
	CommandLine.Int16ListVar(p, name, def, opts...)
}

// Int16List ended ========================================

// Int32List start ========================================

func (f *Flagx) Int32List(name string, def []int32, opts ...Option) *[]int32 {
	v := new([]int32)
	f.Int32ListVar(v, name, def, opts...)
	return v
}

func Int32List(name string, def []int32, opts ...Option) *[]int32 {
	return CommandLine.Int32List(name, def, opts...)
}

func (f *Flagx) Int32ListVar(p *[]int32, name string, def []int32, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*int32List)(p), name, opts...)
}

func Int32ListVar(p *[]int32, name string, def []int32, opts ...Option) {
	CommandLine.Int32ListVar(p, name, def, opts...)
}

// Int32List ended ========================================

// Int64List start ========================================

func (f *Flagx) Int64List(name string, def []int64, opts ...Option) *[]int64 {
	v := new([]int64)
	f.Int64ListVar(v, name, def, opts...)
	return v
}

func Int64List(name string, def []int64, opts ...Option) *[]int64 {
	return CommandLine.Int64List(name, def, opts...)
}

func (f *Flagx) Int64ListVar(p *[]int64, name string, def []int64, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*int64List)(p), name, opts...)
}

func Int64ListVar(p *[]int64, name string, def []int64, opts ...Option) {
	CommandLine.Int64ListVar(p, name, def, opts...)
}

// Int64List ended ========================================

// UintList start ========================================

func (f *Flagx) UintList(name string, def []uint, opts ...Option) *[]uint {
	v := new([]uint)
	f.UintListVar(v, name, def, opts...)
	return v
}

func UintList(name string, def []uint, opts ...Option) *[]uint {
	return CommandLine.UintList(name, def, opts...)
}

func (f *Flagx) UintListVar(p *[]uint, name string, def []uint, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*uintList)(p), name, opts...)
}

func UintListVar(p *[]uint, name string, def []uint, opts ...Option) {
	CommandLine.UintListVar(p, name, def, opts...)
}

// UintList ended ========================================

// Uint8List start ========================================

func (f *Flagx) Uint8List(name string, def []uint8, opts ...Option) *[]uint8 {
	v := new([]uint8)
	f.Uint8ListVar(v, name, def, opts...)
	return v
}

func Uint8List(name string, def []uint8, opts ...Option) *[]uint8 {
	return CommandLine.Uint8List(name, def, opts...)
}

func (f *Flagx) Uint8ListVar(p *[]uint8, name string, def []uint8, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*uint8List)(p), name, opts...)
}

func Uint8ListVar(p *[]uint8, name string, def []uint8, opts ...Option) {
	CommandLine.Uint8ListVar(p, name, def, opts...)
}

// Uint8List ended ========================================

// Uint16List start ========================================

func (f *Flagx) Uint16List(name string, def []uint16, opts ...Option) *[]uint16 {
	v := new([]uint16)
	f.Uint16ListVar(v, name, def, opts...)
	return v
}

func Uint16List(name string, def []uint16, opts ...Option) *[]uint16 {
	return CommandLine.Uint16List(name, def, opts...)
}

func (f *Flagx) Uint16ListVar(p *[]uint16, name string, def []uint16, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*uint16List)(p), name, opts...)
}

func Uint16ListVar(p *[]uint16, name string, def []uint16, opts ...Option) {
	CommandLine.Uint16ListVar(p, name, def, opts...)
}

// Uint16List ended ========================================

// Uint32List start ========================================

func (f *Flagx) Uint32List(name string, def []uint32, opts ...Option) *[]uint32 {
	v := new([]uint32)
	f.Uint32ListVar(v, name, def, opts...)
	return v
}

func Uint32List(name string, def []uint32, opts ...Option) *[]uint32 {
	return CommandLine.Uint32List(name, def, opts...)
}

func (f *Flagx) Uint32ListVar(p *[]uint32, name string, def []uint32, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*uint32List)(p), name, opts...)
}

func Uint32ListVar(p *[]uint32, name string, def []uint32, opts ...Option) {
	CommandLine.Uint32ListVar(p, name, def, opts...)
}

// Uint32List ended ========================================

// Uint64List start ========================================

func (f *Flagx) Uint64List(name string, def []uint64, opts ...Option) *[]uint64 {
	v := new([]uint64)
	f.Uint64ListVar(v, name, def, opts...)
	return v
}

func Uint64List(name string, def []uint64, opts ...Option) *[]uint64 {
	return CommandLine.Uint64List(name, def, opts...)
}

func (f *Flagx) Uint64ListVar(p *[]uint64, name string, def []uint64, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*uint64List)(p), name, opts...)
}

func Uint64ListVar(p *[]uint64, name string, def []uint64, opts ...Option) {
	CommandLine.Uint64ListVar(p, name, def, opts...)
}

// Uint64List ended ========================================

// Float32List start ========================================

func (f *Flagx) Float32List(name string, def []float32, opts ...Option) *[]float32 {
	v := new([]float32)
	f.Float32ListVar(v, name, def, opts...)
	return v
}

func Float32List(name string, def []float32, opts ...Option) *[]float32 {
	return CommandLine.Float32List(name, def, opts...)
}

func (f *Flagx) Float32ListVar(p *[]float32, name string, def []float32, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*float32List)(p), name, opts...)
}

func Float32ListVar(p *[]float32, name string, def []float32, opts ...Option) {
	CommandLine.Float32ListVar(p, name, def, opts...)
}

// Float32List ended ========================================

// Float64List start ========================================

func (f *Flagx) Float64List(name string, def []float64, opts ...Option) *[]float64 {
	v := new([]float64)
	f.Float64ListVar(v, name, def, opts...)
	return v
}

func Float64List(name string, def []float64, opts ...Option) *[]float64 {
	return CommandLine.Float64List(name, def, opts...)
}

func (f *Flagx) Float64ListVar(p *[]float64, name string, def []float64, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*float64List)(p), name, opts...)
}

func Float64ListVar(p *[]float64, name string, def []float64, opts ...Option) {
	CommandLine.Float64ListVar(p, name, def, opts...)
}

// Float64List ended ========================================

// BoolList start ========================================

func (f *Flagx) BoolList(name string, def []bool, opts ...Option) *[]bool {
	v := new([]bool)
	f.BoolListVar(v, name, def, opts...)
	return v
}

func BoolList(name string, def []bool, opts ...Option) *[]bool {
	return CommandLine.BoolList(name, def, opts...)
}

func (f *Flagx) BoolListVar(p *[]bool, name string, def []bool, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*boolList)(p), name, opts...)
}

func BoolListVar(p *[]bool, name string, def []bool, opts ...Option) {
	CommandLine.BoolListVar(p, name, def, opts...)
}

// BoolList ended ========================================

// StringList start ========================================

func (f *Flagx) StringList(name string, def []string, opts ...Option) *[]string {
	v := new([]string)
	f.StringListVar(v, name, def, opts...)
	return v
}

func StringList(name string, def []string, opts ...Option) *[]string {
	return CommandLine.StringList(name, def, opts...)
}

func (f *Flagx) StringListVar(p *[]string, name string, def []string, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*stringList)(p), name, opts...)
}

func StringListVar(p *[]string, name string, def []string, opts ...Option) {
	CommandLine.StringListVar(p, name, def, opts...)
}

// StringList ended ========================================

// DurationList start ========================================

func (f *Flagx) DurationList(name string, def []time.Duration, opts ...Option) *[]time.Duration {
	v := new([]time.Duration)
	f.DurationListVar(v, name, def, opts...)
	return v
}

func DurationList(name string, def []time.Duration, opts ...Option) *[]time.Duration {
	return CommandLine.DurationList(name, def, opts...)
}

func (f *Flagx) DurationListVar(p *[]time.Duration, name string, def []time.Duration, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*durationList)(p), name, opts...)
}

func DurationListVar(p *[]time.Duration, name string, def []time.Duration, opts ...Option) {
	CommandLine.DurationListVar(p, name, def, opts...)
}

// DurationList ended ========================================

// FileList start ========================================

func (f *Flagx) FileList(name string, def []*os.File, opts ...Option) *[]*os.File {
	v := new([]*os.File)
	f.FileListVar(v, name, def, opts...)
	return v
}

func FileList(name string, def []*os.File, opts ...Option) *[]*os.File {
	return CommandLine.FileList(name, def, opts...)
}

func (f *Flagx) FileListVar(p *[]*os.File, name string, def []*os.File, opts ...Option) {
	if def != nil {
		*p = def
	}
	f.append((*fileList)(p), name, opts...)
}

func FileListVar(p *[]*os.File, name string, def []*os.File, opts ...Option) {
	CommandLine.FileListVar(p, name, def, opts...)
}

// FileList ended ========================================
