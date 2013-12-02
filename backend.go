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

// Backend represents some abstract backend that has connections and
// optionsllay a weight.  The Value gives the actual backend value used
// by the application.  Backend implements the io.Closer interface so
// callers should use defer backend.Close()
type Backend struct {
	Connections uint32
	Weight      uint32
	Value       interface{}
}

// NewBackend returns a new backend using the given value and 100% weighting.
func NewBackend(value interface{}) *Backend {
	return &Backend{
		Value:  value,
		Weight: 100,
	}
}

// NewWeightedBackend returns a new backend using the given value and weithing.
func NewWeightedBackend(value interface{}, weight uint32) *Backend {
	return &Backend{
		Value:  value,
		Weight: weight,
	}
}

func (b *Backend) open() *Backend {
	b.Connections = b.Connections + 1
	return b
}

// Close closes the backend, releasing its connections.
func (b *Backend) Close() error {
	b.Connections = b.Connections - 1
	return nil
}
