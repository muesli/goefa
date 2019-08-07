goEFA
=====

package goefa implements a Go client library to access data of public transport
services, which provide an EFA interface. You can search for a stop, query for
upcoming departures and request a route / trip itinerary.

![goefa](misc/goefa.png)

## Installation

Make sure you have a working Go environment (Go 1.3 or higher is required).
See the [install instructions](http://golang.org/doc/install.html).

To install goEFA, simply run:

    go get github.com/muesli/goefa

To compile it from source:

    cd $GOPATH/src/github.com/muesli/goefa
    go get -u -v
    go build && go test -v

## Usage
Simple example on how to verify a stop and get the departures:
```go
// create a new EFAProvider
provider := goefa.NewProvider("https://efa.mvv-muenchen.de/mvv/", true)

// Find a stop by name
stops, err := provider.FindStop("Königsplatz")

// Get the 5 next departures for a stop
deps, err := stops[0].Departures(time.Now(), 5)

// Plan a trip between two stops
routes, err := provider.Route(originStop.ID, destinationStop.ID, time.Now())
...
```

## Available Providers

| City          | Provider | Base URL                                |
| ------------- | -------- | --------------------------------------- |
| Augsburg      | AVV      | https://efa.avv-augsburg.de/avv/        |
| Munich        | MVV      | https://efa.mvv-muenchen.de/mvv/        |

The german [wikipedia article on EFA](https://de.wikipedia.org/wiki/Elektronische_Fahrplanauskunft_%28Software%29)
contains more information about EFA and available providers.

## Links

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/muesli/goefa)
[![Build Status](https://travis-ci.org/muesli/goefa.svg?branch=master)](https://travis-ci.org/muesli/goefa)
[![Coverage Status](https://coveralls.io/repos/github/muesli/goefa/badge.svg?branch=master)](https://coveralls.io/github/muesli/goefa?branch=master)
[![Go ReportCard](http://goreportcard.com/badge/muesli/goefa)](http://goreportcard.com/report/muesli/goefa)
