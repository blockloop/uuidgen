uuidgen
---

uuidgen is intended to mimic [GNU uuidgen](http://gnu.wiki/man1/uuidgen.1.php) which 
provides both random (v4) and time (v1) based UUIDs.


## Installation

```bash
go install github.com/blockloop/uuidgen
```

## Usage

```bash
$GOPATH/bin/uuidgen    #creates a random UUID
$GOPATH/bin/uuidgen -r #creates a random UUID
$GOPATH/bin/uuidgen -t #creates a time-based UUID
```

## Why

This package does not do anything special whatsoever. It wraps an already great package
([satori/go.uuid](github.com/satori/go.uuid)) and provides an installable binary that I
can install because macOS lacks the same `uuidgen` as GNU linux.
