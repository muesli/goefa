goefa
=====

someday this is going to be a golang client for EFA (Elektronische Fahrplan Auskunft)

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
