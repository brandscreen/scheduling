package scheduling_test

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
	"github.com/brandscreen/scheduling"
	"testing"
)

func TestLeastConnections(t *testing.T) {
	s := scheduling.NewLeastConnectionsScheduler(
		[]*scheduling.Backend{
			&scheduling.Backend{Weight: 100, Connections: 5},
			&scheduling.Backend{Weight: 100, Connections: 2},
			&scheduling.Backend{Weight: 100, Connections: 6},
			&scheduling.Backend{Weight: 50, Connections: 4},
			&scheduling.Backend{Weight: 100, Connections: 5},
		})
	runTests(t, s, [][]uint32{
		{5, 3, 6, 4, 5},
		{5, 4, 6, 4, 5},
		{5, 5, 6, 4, 5},
		{5, 5, 6, 5, 5},
		{6, 5, 6, 5, 5},
		{6, 6, 6, 5, 5},
		{6, 6, 6, 6, 5},
		{6, 6, 6, 6, 6},
	})
	expectBackend(t, s, true)
}

func TestLeastConnectionsWithNoBackends(t *testing.T) {
	s := scheduling.NewLeastConnectionsScheduler(make([]*scheduling.Backend, 0))
	expectBackend(t, s, false)
}

func TestLeastConnectionsWithSingleBackend(t *testing.T) {
	s := scheduling.NewLeastConnectionsScheduler([]*scheduling.Backend{
		&scheduling.Backend{Weight: 100},
	})
	runTests(t, s, [][]uint32{
		{1},
		{2},
		{3},
		{4},
	})
	expectBackend(t, s, true)
}

func TestLeastConnectionsWithSingleInactiveBackend(t *testing.T) {
	s := scheduling.NewLeastConnectionsScheduler([]*scheduling.Backend{
		&scheduling.Backend{Weight: 0},
	})
	runTests(t, s, [][]uint32{
		{0},
		{0},
		{0},
		{0},
	})
	expectBackend(t, s, false)
}
