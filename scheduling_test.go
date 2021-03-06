package scheduling

/*
Copyright (c) 2013 Brandscreen Pty Ltd

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

import (
	"fmt"
	"testing"
)

func connectionsMatch(backend []*Backend, connections []uint32) bool {
	for i := range connections {
		if backend[i].Connections != connections[i] {
			return false
		}
	}
	return true
}

func runTests(t *testing.T, scheduler Scheduler, tests [][]uint32) {
	for _, test := range tests {
		scheduler.Schedule()
		if !connectionsMatch(scheduler.GetBackends(), test) {
			fmt.Println(test)
			fmt.Println(scheduler.GetBackends())
			t.Fatal("Expected connections to match")
		}
	}
}

func expectBackend(t *testing.T, scheduler Scheduler, expected bool) {
	be := scheduler.Schedule()
	if expected {
		if be == nil {
			t.Fatal("Expected to receive a backend")
		}
	} else {
		if be != nil {
			t.Fatal("Expected to receive no backends")
		}
	}
}
