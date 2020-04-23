# BSDOpts [![GoDoc](https://godoc.org/github.com/cristianrz/BSDOpts?status.svg)](https://godoc.org/github.com/cristianrz/BSDOpts)

Package BSDOpts implements BSD-style command-line flag parsing in less than
50 lines of code.

## Installation

Install with 

```
go get github.com/cristianrz/BSDOpts
```

## Usage

Import with

```
import "github.com/cristianrz/BSDOpts"`
```

Create a set of flags using 

```
args := make(BSDOpts.Opts)
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

