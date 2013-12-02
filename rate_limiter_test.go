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
	"time"
)

func TestRateLimiter(t *testing.T) {
	tests := []struct {
		remaining  int64
		sleep      time.Duration
		successful bool
	}{
		{4, 0, true},
		{3, 0, true},
		{2, 0, true},
		{1, 0, true},
		{0, 0, true},
		{0, 0, false},
		{0, 2, false},
		{4, 0, true},
		{3, 0, true},
		{2, 1, true},
		{4, 0, true},
		{3, 0, true},
		{2, 0, true},
		{1, 0, true},
		{0, 0, true},
		{0, 0, false},
	}

	r := scheduling.NewRateLimiter(5, time.Millisecond)
	for _, test := range tests {
		if test.successful != r.Schedule() {
			t.Error("Expected successful", test.successful)
		}

		if r.Remaining != test.remaining {
			t.Error("Expected remaining to be ", test.remaining, "but was", r.Remaining)
		}

		time.Sleep(test.sleep * time.Millisecond)
	}
}
