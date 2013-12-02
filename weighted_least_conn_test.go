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
	"code.brandscreen.net/go/scheduling"
	"testing"
)

func TestWeightedLeastConnections(t *testing.T) {
	s := scheduling.NewWeightedLeastConnectionsScheduler([]*scheduling.Backend{
		&scheduling.Backend{Weight: 100},
		&scheduling.Backend{Weight: 100},
		&scheduling.Backend{Weight: 100},
		&scheduling.Backend{Weight: 50},
		&scheduling.Backend{Weight: 100},
	})
	runTests(t, s, [][]uint32{
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1},
		{2, 1, 1, 1, 1},
		{2, 2, 1, 1, 1},
		{2, 2, 2, 1, 1},
		{2, 2, 2, 1, 2},
		{3, 2, 2, 1, 2},
		{3, 3, 2, 1, 2},
		{3, 3, 3, 1, 2},
		{3, 3, 3, 2, 2},
		{3, 3, 3, 2, 3},
		{4, 3, 3, 2, 3},
		{4, 4, 3, 2, 3},
		{4, 4, 4, 2, 3},
		{4, 4, 4, 2, 4},
	})
	expectBackend(t, s, true)
}

func TestWeightedLeastConnectionsWithSingleBackend(t *testing.T) {
	s := scheduling.NewWeightedLeastConnectionsScheduler([]*scheduling.Backend{
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

func TestWeightedLeastConnectionsWithSingleInactiveBackend(t *testing.T) {
	s := scheduling.NewWeightedLeastConnectionsScheduler([]*scheduling.Backend{
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

func TestWeightedLeastConnectionsWithNoBackends(t *testing.T) {
	s := scheduling.NewWeightedLeastConnectionsScheduler(make([]*scheduling.Backend, 0))
	expectBackend(t, s, false)
}
