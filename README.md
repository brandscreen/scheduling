scheduling
==========

[![Build Status](https://travis-ci.org/brandscreen/scheduling.png)](https://travis-ci.org/brandscreen/scheduling) [![Coverage Status](https://coveralls.io/repos/brandscreen/scheduling/badge.png?branch=HEAD)](https://coveralls.io/r/brandscreen/scheduling?branch=HEAD)

[Scheduling algorithms](http://www.loadbalancerblog.com/blog/2013/06/load-balancing-scheduling-methods-explained) in [Go](http://golang.org) language.  The following algorithms are implemented:

* __Round Robin__ - Schedules each request sequentially around the pool of backends.  Using this algorithm, all the backends are treated as equals.
* __Least Connections__ - Schedules more requests to backends with fewer active connections.
* __Weighted Least Connections__ - Schedules each request sequentially around the pool of backends but gives more requests to backends with greater weights (capacity).
* __Rate Limiting__ - Schedules a request to a backend based on a set limit of requests per duration (e.g., requests/second), using the [Token bucket](http://en.wikipedia.org/wiki/Token_bucket) algorithm.

# Requirements

* Go 1.1 or higher

# Installation

```bash
$ go get github.com/brandscreen/scheduling
```

# Usage

## Import Scheduling

Ensure you have imported the scheduling package at the top of your source file.

```golang
import "github.com/brandscreen/scheduling"
```

## Scheduling

### Setup the Backends

Setup a slice of Backends.  Each `Backend` has a `Value` which provides context to the application for selecting the real server.

```golang
var server1, server2, server3 interface{} // These are your actual servers...

backends := make([]*scheduling.Backend, 3)
backends[0] = scheduling.NewBackend(server1)
backends[1] = scheduling.NewWeightedBackend(server2, 50)
backends[2] = scheduling.NewBackend(server3)
```

### Create a Scheduler

Next, create an instance of the scheduling algorithm.

#### Round Robin

```golang
scheduler := scheduling.NewRoundRobinScheduler(backends)
```

#### Least Connections

```golang
scheduler := scheduling.NewLeastConnectionsScheduler(backends)
```

#### Weighted Least Connections

```golang
scheduler := scheduling.NewWeightedLeastConnectionsScheduler(backends)
```

### Schedule an activity using the Scheduler

Call the `Schedule` function on the scheduler to obtain the selected `Backend`.  Note you should call `Close()` in order to release the backend otherwise the connection count will continue to increase.

```golang
backend := scheduler.Schedule()
defer backend.Close()
```

## Rate Limiting

Not strictly a scheduling algorithm, you can rate limit requests to a backend using the `RateLimiter` object.

### Create a RateLimiter

This example creates a `RateLimiter` that is limited to 500 requests per second.  Ensure to import `time` package.

```golang
limiter := scheduling.NewRateLimiter(500, time.Second)
```

### Check if the request fits within the limit

Call the `Schedule()` function to determine if the request can be sent to the backend.

```golang
if limiter.Schedule() {
    // Send the request to the backend
} else {
    // Drop the request
}
```

# Contributing

Please see [CONTRIBUTING.md](https://github.com/brandscreen/scheduling/blob/master/CONTRIBUTING.md).  If you have a bugfix or new feature that you would like to contribute, please find or open an issue about it first.

# License

Licensed under the MIT License.
