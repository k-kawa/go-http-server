go-http-server
===

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][LICENSE]

[license]: https://github.com/k-kawa/go-http-server/blob/master/LICENSE

The easiest way to start a web server in your Windows/Mac/Linux laptop.
Downloand it, Copy it and Execute it. That's all!

## Usage
You just execute `go-http-server start` in your console.

```bash
$ go-http-server start
```

You can specify the port number the server should listen to like this.

```bash
$ go-http-server start --port 8080
```

## Cross compile
Currently the following OSs and Architectures are supported.

- Linux/386
- Linux/amd64
- Darwin/386
- Darwin/amd64
- Windows/386
- Windows/amd64

To make a binary for each environment, Gox is required.
Setting Gox up is easy by the following commands

```bash
# Install gox
$ go get github.com/mitchellh/gox

# Setup tools
$ gox -build-toolchain
```

Then you can cross-compile it by this command.

```bash
# Compile them
$ ./build.sh
```

You'll find the binaries in `pkg` directory.

## Author
[Kohei Kawasaki](https://github.com/k-kawa)
