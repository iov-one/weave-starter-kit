# Weave Starter Kit

This is a simple app to show how to build an application
using the weave framework. Feel free to copy this repo
to start your own application.

Before starting with this code (or while you do so), you may well want to check
out the documentation at [iov-docs](https://docs.iov.one), which presents many
of the high-level concepts you need to understand in order to code.

You will probably want to clone http://github.com/iov-one/weave and
keep this in another folder in your IDE for reference. Or, maybe you
find [godocs](http://godoc.org/github.com/iov-one/weave) more helpful.

## Installing dependencies

### Requirements

- [golang 1.11.4+](https://golang.org/doc/install)
- [tendermint 0.31.5](https://github.com/tendermint/tendermint/tree/v0.31.5)
  - [Installation](https://github.com/tendermint/tendermint/blob/master/docs/introduction/install.md)
- [weave](https://github.com/iov-one/weave)
  - `go get github.com/iov-one/weave`
- [docker](https://docs.docker.com/install/)

**Important**: At IOV we use [go modules](https://github.com/golang/go/wiki/Modules). You can append `export GO111MODULE=on` to `.rc` file of your favorite shell(`zsh, bash, fish, etc.`)

## Running the demo app

```sh
cd <your-project-location>
make            # install binary
make inittm     # initilize tendermint config folder ~/.custom
customd init    # initilize genesis
make runtm      # run tendermint
customd start   # run customd application
```

## Interacting with demo app

After `make`, `customcli` will be built and placed in go bin. You can play with `customcli`. Be sure that `customd` is running.

```sh
customcli send-tokens \
        -src "seq:test/custom/1" \
        -dst "seq:test/custom/2" \
        -amount "4 CSTM" \
        -memo "customcli test" | customcli view
```

## Using for your own app

You have two options: fork or copy-replace.

### Forking

I recommend forking the repo if you want to test something out. Also merging the latest changes from `weave` will be easier.

1. Fork the project
2. Append `replace github.com/iov-one/weave-starter-kit => github.com/<username>/<project-name> v0.0.1` to `go.mod`. Example: [github.com/orkunkl/starter-kit-test](https://github.com/orkunkl/starter-kit-test/blob/master/go.mod#L15)

### Copy then replace

You can download the project and make the required replacement. You will want to adjust it to proper project path and then
do a search and replace for `iov-one/weave-starter-kit` with your github project name.
In particular:

- prototool.yaml (import path)
- all import paths
- all files in `cmd`

## Building custom modules

`x` directory contains modules. `x/custom` directory is a basic module placeholder without any meaningful functionality. You can write your own module via copying examples from `custom` module. After you implement your module, write required code up to `cmd` folder to reflect the module on blockchain app.

## Protobuf tips

After you edit any `.proto` file, do not forget to run `make protoc` in projects root. Otherwise you will not see the changes in go code
