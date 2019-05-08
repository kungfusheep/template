package template

import (
	"reflect"
	"unsafe"
)

// Engine is the
type Engine struct {
	instructions []func(unsafe.Pointer, *Buffer) // instruction set
	f            reflect.StructField             // current field
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
	// tt := reflect.TypeOf(t)

	// e.chunk(string(template))

	for e.i = 0; e.i < len(template); e.i++ {

		switch b := template[e.i]; {
		case b == '{' && template[e.i+1] == '{':
			e.i++
			e.flunk()
			// value instruction

			// val := e.readVal()

			// println(val)

		default:
			// add to static chunk
			e.chunkb(b)
		}
	}

	e.flunk()
	return e
}

func (e *Engine) readVal() string {

	return ""
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
