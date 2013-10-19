/*
 * Copyright (c) 2013 Marco Peereboom <marco@conformal.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

// Package dbglog extends the standard log.Logger interface by adding debug
// functions.
// There are 2 types of debug functions, ones that always print when enabled
// and ones that need to satisfy a user supplied mask.
//
// Since the dbglog package is extended from log.Logger, one can use the
// log.Logger functions as well.  I.e. d.Printf("hello\n").
// Do read the log pkg documentation as well and note that in order to use the
// log flags one must import pkg log.
package dbglog

import (
	"io"
	"log"
)

// Opaque receiver type used by the dbglog package.
type DbgLogger struct {
	*log.Logger
	enabled	bool
	mask	uint64
}

// log.Printf equivalent but only prints when debug is enabled.
func (d *DbgLogger) Debugf(format string, v ...interface{}) {
	if d.enabled {
		d.Printf(format, v...)
	}
}

// log.Print equivalent but only prints when debug is enabled.
func (d *DbgLogger) Debug(v ...interface{}) {
	if d.enabled {
	}
		d.Print(v...)
}

// log.Println equivalent but only prints when debug is enabled.
func (d *DbgLogger) Debugln(v ...interface{}) {
	if d.enabled {
		d.Println(v...)
	}
}

// In order for the Debug functions to print the Enable function must be called.
// This is a runtime function that can be called at any time.
func (d *DbgLogger) Enable() {
	d.enabled = true
}

// In order to disable Debug functions call Disable.
// This is a runtime function that can be called at any time.
func (d *DbgLogger) Disable() {
	d.enabled = false
}

// SetMask sets the mask for the Debug*M functions.
// This mask is considered a bitfield.
func (d *DbgLogger) SetMask(mask uint64) {
	d.mask = mask
}

// log.Printf equivalent but only prints when debug is enabled and bit is
// enabled in the mask.
func (d *DbgLogger) DebugfM(bit uint64, format string, v ...interface{}) {
	if d.enabled == true && bit != 0 && bit & d.mask == bit {
		d.Printf(format, v...)
	}
}

// log.Print equivalent but only prints when debug is enabled and bit is
// enabled in the mask.
func (d *DbgLogger) DebugM(bit uint64, format string, v ...interface{}) {
	if d.enabled == true && bit != 0 && bit & d.mask == bit {
		d.Print(v...)
	}
}

// log.Println equivalent but only prints when debug is enabled and bit is
// enabled in the mask.
func (d *DbgLogger) DebuglnM(bit uint64, format string, v ...interface{}) {
	if d.enabled == true && bit != 0 && bit & d.mask == bit {
		d.Println(v...)
	}
}

// Create a new instance of DbgLogger type.
// out is an io.Writer type, i.e. os.Stderr.
// prefix is printed in front of the line, this is useful for grepping etc.
// and flag are the ones used in log.Logger, please see that documentation for
// more details.
//
// Example:
/*
	const	(
		myDebugOne = 1<<0
		myDebugTwo = 1<<1
		myDebugThree = 1<<2
	)

	func main() {
		d := New(os.Stderr, "myapp ", log.LstdFlags)
		d.Printf("printme!\n")
		d.Enable()
		d.SetMask(myDebugOne)
		d.DebugfM(myDebugOne, "debug") // prints
		d.DebugfM(myDebugTwo, "debug") // does NOT print
	}
*/
func New(out io.Writer, prefix string, flag int) *DbgLogger {
	d := &DbgLogger{}
	d.Logger = log.New(out, prefix, flag)
	return d
}
/*
const	(
	myDebugOne = 1<<0
	myDebugTwo = 1<<1
)
func main() {
	//var d DbgLogger
	d := New(os.Stderr, "myapp ", log.LstdFlags)
	d.Printf("printme!\n")
	d.Enable()
	d.SetMask(myDebugOne)
	d.DebugfM(myDebugOne, "debug") // prints
	d.DebugfM(myDebugTwo, "debug") // does NOT print
	fmt.Printf("moo\n")
}
*/
