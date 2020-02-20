# Fuzzctl

CLI to manage workflows with Fuzzball.

## Quick Start

Configure your go environment to pull private go modules, by forcing `go get` to use `git+ssh` instead of `https`. This lets the go compiler pull private dependencies using your machine's ssh keys.

```sh
git config --global url."ssh://git@github.com/sylabs".insteadOf "https://github.com/sylabs"
```

If using Go 1.13, the `go` command defaults to downloading modules from the public Go module mirror, and validating downloaded modules against the public Go checksum database. Since private Sylabs projects are not availble in the public mirror nor the public checksum database, we must tell Go about this. One way to do this is to set `GOPRIVATE` in the Go environment:

```sh
go env -w GOPRIVATE=github.com/sylabs
```

In order for go to execute this binary the path `${GOPATH}/bin` needs to be included in your `PATH`.

To run the CLI, you'll need to either `go get` the tool from github:
```sh
go get -u github.com/sylabs/compute-cli
```

Developers can either install to `${GOPATH}/bin` with:

```sh
go install ./...
```

or build the binary to a temporary location with:

```sh
go build -o <path> ./...
```


Finally, run a command:

```sh
$ Fuzzctl
fuzzctl enables control of workflows for Fuzzball.

Usage:
  fuzzctl [command]

Available Commands:
  create      create enables you to submit a workflow to the Fuzzball queue.
  delete      delete allows you to remove a workflow from the Fuzzball queue.
  help        Help about any command
  info        info allows you to see a workflow's state within Fuzzball.
  list        list allows you to list workflows within Fuzzball.

Flags:
  -d, --debug              Enable debug output
  -h, --help               help for fuzzctl
      --http_addr string   Address to reach compute server (default "http://localhost:8080")

Use "fuzzctl [command] --help" for more information about a command.
```