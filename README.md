# Fuzzctl

[![Built with Mage](https://magefile.org/badge.svg)](https://magefile.org)
[![CI Workflow](https://github.com/sylabs/fuzzctl/workflows/ci/badge.svg)](https://github.com/sylabs/fuzzctl/actions)
[![Dependabot](https://api.dependabot.com/badges/status?host=github&repo=sylabs/fuzzctl&identifier=233642046)](https://app.dependabot.com/accounts/sylabs/repos/233642046)

CLI to manage workflows with [Fuzzball](https://github.com/sylabs/fuzzball-service).

## Quick Start

Ensure that you have one of the two most recent minor versions of Go installed as per the [installation instructions](https://golang.org/doc/install).

Install [Mage](https://magefile.org) as per the [installation instructions](https://magefile.org/#installation).

To install `fuzzctl` in `$(go env GOPATH)/bin`:

```sh
mage install
```

Ensure `$(go env GOPATH)/bin` is included in your `$PATH`, at which point you can execute `fuzzctl` commands:

```sh
fuzzctl help
```

## Testing

Unit tests can be run like so:

```sh
mage test
```

## Installation and Packaging

To install `fuzzctl` in `$GOBIN`:

```sh
mage install
```

To build a `.deb` and/or `.rpm`:

```sh
mage deb
mage rpm
```

## License

This project is licensed under a 3-clause BSD license found in the [license file](LICENSE.md).
