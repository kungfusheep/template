package template

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Engine is the
type Engine struct {
	instructions []func(unsafe.Pointer, *Buffer) // instruction set
	f            reflect.StructField             // current field
	tpl          []byte                          // current template bytes
	t            interface{}                     // type
	i            int                             // iter
	cb           Buffer                          // side buffer for static data
	cpos         int                             // side buffer position
}

// Execute runs the set of instructions for this template engine using the input
// data provided.
func (e *Engine) Execute(v interface{}, w *Buffer) {

	p := unsafe.Pointer(reflect.ValueOf(v).Pointer())
	for i, l := 0, len(e.instructions); i < l; i++ {
		e.instructions[i](p, w)
	}
}

// New generates a new template.Engine instruction set
func New(template []byte, t interface{}) *Engine {
	e := &Engine{}
	e.t = t
	e.tpl = template
	// tt := reflect.TypeOf(t)

	// e.chunk(string(template))

	for e.i = 0; e.i < len(e.tpl); e.i++ {

		switch b := e.tpl[e.i]; {
		case e.peekAdv(`{` + `{`):
			e.flunk()
			// value instruction

			val, n := e.readVal()
			e.i += n

			println(fmt.Sprint(val, " ", n))

			r, _ := reflect.TypeOf(e.t).FieldByName(val)

			switch r.Type.Kind() {

			case reflect.String:
				e.instructions = append(e.instructions, func(p unsafe.Pointer, w *Buffer) {
					ptrStringToBuf(unsafe.Pointer(uintptr(p)+r.Offset), w)
				})
			}

			println(fmt.Sprint(r))

		default:
			// add to static chunk
			// println(fmt.Sprint(string(b)))
			e.chunkb(b)
		}
	}

	e.flunk()
	return e
}

func (e *Engine) readVal() (o string, l int) {

	// this is garbage

	var reading bool
	for i := e.i; i < len(e.tpl); i++ {
		l++

		switch b := e.tpl[i]; {

		case b == '}' && e.tpl[i+1] == '}':
			// l++
			println("end")
			return

		case reading && b == ' ':
			reading = false

		case !reading && b == '.':
			reading = true

		case reading:
			o += string(e.tpl[i])

		}
	}

	return "", 0
}

func (e *Engine) peekAdv(p string) bool {

	if l := len(p); e.i+l < len(e.tpl) && string(e.tpl[e.i:e.i+l]) == p {
		e.i += l
		return true
	}

	return false
}

// chunk writes a chunk of body data to the chunk buffer. only for writing static
//  structure and not dynamic values.
func (e *Engine) chunk(b []byte) {
	e.cb.Write(b)
}

func (e *Engine) chunkb(b byte) {
	e.cb.WriteByte(b)
}

// flunk flushes whatever chunk data we've got buffered into a single instruction
func (e *Engine) flunk() {

	b := e.cb.Bytes
	bs := b[e.cpos:]
	e.cpos = len(b)

	if len(bs) == 0 {
		return
	}

	e.instructions = append(e.instructions, func(_ unsafe.Pointer, w *Buffer) {
		w.Write(bs)
	})
}
