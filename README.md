# eyhash

[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?logo=go)](go.mod)

> Simple, fast hashing library and CLI for computing MD5, SHA1, SHA256, and SHA512 checksums of files and directories.

## Features

- **Four hash algorithms** — MD5, SHA1, SHA256, SHA512
- **CLI & library** — use it as a command-line tool or import as a Go package
- **Directory support** — hash every file within a directory
- **Cross-platform** — Linux, Windows, macOS (amd64 & arm64)
- **Zero dependencies** — built entirely on Go's standard library

## Installation

### Build from source

```sh
make build
```

The binary will be placed at `./bin/eyhash`.

### Install globally

```sh
make install
```

### Pre-built binaries

Releases are published via [GoReleaser](.goreleaser.yaml) on every tagged push (`v*`). Download the appropriate archive from the [GitHub Releases](../../releases) page.

## Usage

### Command-Line Interface

```
eyhash <filename/directory> -f/-d
```

| Flag | Description             |
|------|-------------------------|
| `-f` | Hash a single file      |
| `-d` | Hash all files in a dir |
| `-h` | Show help               |

**Examples:**

```sh
# Hash a single file
eyhash test.txt -f

# Hash all files in a directory
eyhash ./my-folder -d

# Show usage
eyhash -h
```

**Sample output:**

```
FileName: test.txt, Size: 22 bytes, ModTime: 2024-08-05 00:28:36.2926994 +0700 +07
MD5:    ecc8c106ae46e1f89be3f12fa1352952
SHA1:   ce86972a7a726bad768121c36e50e832e98343b7
SHA256: 951ed50ee98965f335a087777e5ffa6bc171f1ebde1b3b8dd22608409d04f5a2
SHA512: c141a290be00e73a46c3f20f67e4d10ae68a74fc895acbc82749001a347d75dcdf35b53a14752698a1517f8333bf82591ffe817bbb817d183768385cc59dbb9e
```

### As a Go Library

```go
package main

import (
    "fmt"

    "github.com/garudaproject/eyhash"
)

func main() {
    // Get file metadata
    info, err := eyhash.FileInfo("test.txt")
    if err != nil {
        panic(err)
    }
    fmt.Printf("Name: %s, Size: %d, ModTime: %s\n", info.Name, info.Size, info.ModTime)

    // Compute hashes
    md5, err := eyhash.MD5File("test.txt")
    if err != nil { panic(err) }

    sha1, err := eyhash.SHA1File("test.txt")
    if err != nil { panic(err) }

    sha256, err := eyhash.SHA256File("test.txt")
    if err != nil { panic(err) }

    sha512, err := eyhash.SHA512File("test.txt")
    if err != nil { panic(err) }

    fmt.Println("MD5:   ", md5)
    fmt.Println("SHA1:  ", sha1)
    fmt.Println("SHA256:", sha256)
    fmt.Println("SHA512:", sha512)
}
```

## API Reference

### `eyhash.FileInfo(path string) (*Info, error)`

Returns file metadata including name, size, and modification time.

### `eyhash.MD5File(path string) (string, error)`

Computes the MD5 checksum of the file at `path`. Returns a 32-character hexadecimal string.

### `eyhash.SHA1File(path string) (string, error)`

Computes the SHA1 checksum of the file at `path`. Returns a 40-character hexadecimal string.

### `eyhash.SHA256File(path string) (string, error)`

Computes the SHA256 checksum of the file at `path`. Returns a 64-character hexadecimal string.

### `eyhash.SHA512File(path string) (string, error)`

Computes the SHA512 checksum of the file at `path`. Returns a 128-character hexadecimal string.

## Development

### Makefile Targets

| Target        | Description                           |
|---------------|---------------------------------------|
| `make build`  | Build the binary to `./bin/eyhash`    |
| `make install`| Install binary to `$GOBIN`            |
| `make run`    | Build and run the CLI                 |
| `make test`   | Run all tests                         |
| `make test-race` | Run tests with race detector       |
| `make clean`  | Remove the `./bin` directory          |

### Running Tests

```sh
make test

# Or directly:
go test -v ./...

# With race detector:
make test-race
```

## License

Copyright (c) 2024-2026 GarudaProject. Licensed under the [Apache License, Version 2.0](LICENSE).
