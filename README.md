# HostDB Collector for HP OneView

Queries the OneView REST API, to get all available data, and sends that data to HostDB.

## Getting Started

This section will describe the process of developing the collector.
Please see [Deployment](#deployment) for notes on how the collector is used when deployed.

### Prerequisites

This collector requires a few things to operate:

* An instance of OneView to query
* A HostDB instance to write to

For development, you'll also need:

* Docker
* Golang

### Installing

The collector is a golang binary, and after compilation, it can be run on any Linux x86 system. No installation necessary.

## Running tests

This should be as simple as `make test`. It will execute `go fmt`, `go vet`, `golint`, `errcheck` and `go test`.

## Deployment

The collector ran from a container on a regular schedule.

## Debugging

Set the environment variable `HOSTDB_COLLECTOR_ONEVIEW_COLLECTOR_DEBUG` to true, and the collector will output additional detail, *including secrets*.
In addition, the variable `HOSTDB_COLLECTOR_ONEVIEW_COLLECTOR_SAMPLE_DATA` to true, and the collector will output all collected data to files (configurable via `HOSTDB_COLLECTOR_ONEVIEW_COLLECTOR_SAMPLE_DATA_PATH`), instead of sending it to HostDB.

## Built With

Build system will run tests, compile the golang binary, create a container including the binary, and upload that container image to registry for use.

## Authors & Support

- Email: info@pdxfixit.com

## See Also

- [OneView golang module](https://github.com/HewlettPackard/oneview-golang)
