goefa
=====

package goefa implements a go (golang) client library to access data of public
transport companies which provide an EFA interface. You can search a stop, get
its next departures or request a trip.

![goefa](misc/goefa.png)


## Installation
run

```bash
go get -v
```
to get 3rdparty libraries.

then run
```bash
go run examples/departure_monitor.go
```
to try the departure monitor example

## Usage
Simple example on how to verify a stop and get the departures:
```go
// create a new EFAProvider
myprovider, err := goefa.ProviderFromJson("avv")

// Find a stop by name
idtfd, stops, err := myprovider.FindStop("Königsplatz")

// If stop was identified get the 5 next departures
deps, err := stops[0].Departures(time.Now(), 5)
```

## Adding new providers

Edit the providers.json file to add new EFA providers. The german [wikipedia article on EFA](https://de.wikipedia.org/wiki/Elektronische_Fahrplanauskunft_%28Software%29) contains some hints at providers etc.

## Links

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/michiwend/goefa)
