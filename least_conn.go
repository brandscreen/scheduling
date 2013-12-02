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

/*
LeastConnectionsScheduler is a scheduler that uses connections to determine the next
backend to use.  It does not take into account any weighting applied to the backends.
*/
type LeastConnectionsScheduler struct {
	backends []*Backend
}

// NewLeastConnectionsScheduler returns a new LeastConnectionsScheduler with the given
// backends.
func NewLeastConnectionsScheduler(backends []*Backend) Scheduler {
	return &LeastConnectionsScheduler{
		backends: backends,
	}
}

// Schedule returns the next backend according to the algorithm.
func (s *LeastConnectionsScheduler) Schedule() *Backend {
	const maxConnections = ^uint32(0)

	minConnections := maxConnections
	bestBackend := uint32(0)

	for m, backend := range s.backends {
		if backend.Weight > 0 && backend.Connections < minConnections {
			minConnections = backend.Connections
			bestBackend = uint32(m)
		}
	}

	if minConnections != maxConnections {
		return s.backends[bestBackend].open()
	}

	return nil
}

// Returns the backends in use
func (s *LeastConnectionsScheduler) GetBackends() []*Backend {
	return s.backends
}
