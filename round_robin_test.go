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

func TestRoundRobin(t *testing.T) {
	s := scheduling.NewRoundRobinScheduler([]*scheduling.Backend{
		&scheduling.Backend{Weight: 100, Connections: 5},
		&scheduling.Backend{Weight: 100, Connections: 2},
		&scheduling.Backend{Weight: 100, Connections: 6},
		&scheduling.Backend{Weight: 50, Connections: 4},
		&scheduling.Backend{Weight: 100, Connections: 5},
	})
	runTests(t, s, [][]uint32{
		{5, 3, 6, 4, 5},
		{5, 3, 7, 4, 5},
		{5, 3, 7, 5, 5},
		{5, 3, 7, 5, 6},
		{6, 3, 7, 5, 6},
		{6, 4, 7, 5, 6},
		{6, 4, 8, 5, 6},
		{6, 4, 8, 6, 6},
		{6, 4, 8, 6, 7},
	})
	expectBackend(t, s, true)
}

func TestRoundRobinWithSingleBackend(t *testing.T) {
	s := scheduling.NewRoundRobinScheduler([]*scheduling.Backend{
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

func TestRoundRobinWithSingleInactiveBackend(t *testing.T) {
	s := scheduling.NewRoundRobinScheduler([]*scheduling.Backend{
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

func TestRoundRobinWithTwoInactiveBackends(t *testing.T) {
	s := scheduling.NewRoundRobinScheduler([]*scheduling.Backend{
		&scheduling.Backend{Weight: 0},
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

func TestRoundRobinWithNoBackends(t *testing.T) {
	s := scheduling.NewRoundRobinScheduler(make([]*scheduling.Backend, 0))
	expectBackend(t, s, false)
}
