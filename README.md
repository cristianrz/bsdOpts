# opts [![GoDoc](https://godoc.org/github.com/cristianrz/opts?status.svg)](https://godoc.org/github.com/cristianrz/opts)

Package opts implements BSD-style command-line flag parsing in less than
50 lines of code.

## Installation

Install with 

```
go get github.com/cristianrz/opts
```

## Usage

Import with

```
import "github.com/cristianrz/opts"`
```

Create a set of flags using 

```
args := make(opts.Opts)
```

Define each flags default value as 

```
args['b'] = false
args['i'] = 3
args['f'] = file.txt
```

Parse with

```
args.Parse()
```

