# bsdOpts [![GoDoc](https://godoc.org/github.com/cristianrz/bsdOpts?status.svg)](https://godoc.org/github.com/cristianrz/bsdOpts)

Package bsdOpts implements BSD-style command-line flag parsing in less than
50 lines of code.

## Installation

Install with 

```
go get github.com/cristianrz/bsdOpts
```

## Usage

Import with

```
import "github.com/cristianrz/bsdOpts"`
```

Create a set of flags using 

```
args := make(bsdOpts.Opts)
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

