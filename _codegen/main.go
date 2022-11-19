package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

const (
	_PACKAGE = `package flagx

import (
	"os"
	"time"
)
`
	_TEMPLATE = `
// {{ .Upper }} start ========================================

func (f *Flagx) {{ .Upper }}(short rune, long string, def {{ .Def }}, opts ...Option) *{{ .Type }} {
	v := new({{ .Type }})
	f.{{ .Upper }}Var(v, short, long, def, opts...)
	return v
}

func {{ .Upper }}(short rune, long string, def {{ .Def }}, opts ...Option) *{{ .Type }} {
	return CommandLine.{{ .Upper }}(short, long, def, opts...)
}

func (f *Flagx) {{ .Upper }}Var(p *{{ .Type }}, short rune, long string, def {{ .Def }}, opts ...Option) {
	{{- if .Check }}
	if def != nil {
		*p = {{ .DefSt }}def
	}
	{{- else }}
	*p = {{ .DefSt }}def
	{{- end }}
	f.append((*{{ .Lower }})(p), short, long, def, opts...)
}

func {{ .Upper }}Var(p *{{ .Type }}, short rune, long string, def {{ .Def }}, opts ...Option) {
	CommandLine.{{ .Upper }}Var(p, short, long, def, opts...)
}

// {{ .Upper }} ended ========================================
`
)

func main() {
	keys := []string{
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64",
		"bool", "string", "duration", "file",
		"intList", "int8List", "int16List", "int32List", "int64List",
		"uintList", "uint8List", "uint16List", "uint32List", "uint64List",
		"float32List", "float64List",
		"boolList", "stringList", "durationList", "fileList"}

	types := []string{
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64", "bool", "string", "time.Duration", "os.File",
		"[]int", "[]int8", "[]int16", "[]int32", "[]int64",
		"[]uint", "[]uint8", "[]uint16", "[]uint32", "[]uint64",
		"[]float32", "[]float64", "[]bool", "[]string", "[]time.Duration", "[]*os.File"}

	parser, err := template.New("generate").Parse(_TEMPLATE)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("flag_gen.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(_PACKAGE)
	if err != nil {
		return
	}

	for i, k := range keys {
		typ := types[i]
		defStar := ""
		def := typ
		check := k == "file" || strings.HasSuffix(k, "List")
		if k == "file" {
			defStar = "*"
			def = "*" + def
		}
		upper := strings.ToUpper(string(k[0])) + k[1:]
		if !strings.HasSuffix(k, "List") {
			k += "Value"
		}

		err := parser.Execute(file, map[string]any{
			"Upper": upper,
			"Lower": k,
			"Type":  typ,
			"Def":   def,
			"DefSt": defStar,
			"Check": check,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}
