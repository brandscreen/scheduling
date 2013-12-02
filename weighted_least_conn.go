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
WeightedLeastConnectionsScheduler is a Scheduler that uses a weighted least
connections algorithm to schedule backends.

Supposing there is a server set S = {S0, S1, ..., Sn-1},
W(Si) is the weight of server Si;
C(Si) is the current connection number of server Si;
CSUM = ΣC(Si) (i=0, 1, .. , n-1) is the sum of current connection numbers;

The new connection is assigned to the server j, in which
  (C(Sm) / CSUM)/ W(Sm) = min { (C(Si) / CSUM) / W(Si)}  (i=0, 1, . , n-1),
  where W(Si) isn't zero
Since the CSUM is a constant in this lookup, there is no need to divide by CSUM,
the condition can be optimized as
  C(Sm) / W(Sm) = min { C(Si) / W(Si)}  (i=0, 1, . , n-1), where W(Si) isn't zero
*/
type WeightedLeastConnectionsScheduler struct {
	backends []*Backend
}

// NewWeightedLeastConnectionsScheduler returns a new WeightedLeastConnectionsScheduler
func NewWeightedLeastConnectionsScheduler(backends []*Backend) *WeightedLeastConnectionsScheduler {
	return &WeightedLeastConnectionsScheduler{
		backends: backends,
	}
}

// Schedule returns the next backend according to the algorithm.
func (s *WeightedLeastConnectionsScheduler) Schedule() *Backend {
	n := uint32(len(s.backends))

	for m := uint32(0); m < n; m++ {
		if s.backends[m].Weight > 0 {
			for i := m + 1; i < n; i++ {
				if s.backends[m].Connections*s.backends[i].Weight > s.backends[i].Connections*s.backends[m].Weight {
					m = i
				}
			}
			return s.backends[m].open()
		}
	}

	return nil
}

// Returns the backends in use
func (s *WeightedLeastConnectionsScheduler) GetBackends() []*Backend {
	return s.backends
}
