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
	"testing"
)

func TestNewBackend(t *testing.T) {
	value := 5
	be := NewBackend(value)
	if be.Connections != 0 {
		t.Error("Expected no connections")
	}
	if be.Weight != 100 {
		t.Error("Expected weight of 100")
	}
	if be.Value != value {
		t.Error("Expected value to be set")
	}

	be.open()
	if be.Connections != 1 {
		t.Error("Expected 1 connections")
	}

	be.Close()
	if be.Connections != 0 {
		t.Error("Expected no connections")
	}
}

func TestNewWeightedBackend(t *testing.T) {
	value := 5
	be := NewWeightedBackend(value, 75)
	if be.Connections != 0 {
		t.Error("Expected no connections")
	}
	if be.Weight != 75 {
		t.Error("Expected weight of 75")
	}
	if be.Value != value {
		t.Error("Expected value to be set")
	}
}
