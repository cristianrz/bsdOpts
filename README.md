# bsdOpts

[![GoDoc](https://godoc.org/github.com/cristianrz/bsdOpts?status.svg)](https://godoc.org/github.com/cristianrz/bsdOpts)

`import "github.com/cristianrz/bsdOpts"`

 Package bsdOpts implements BSD-style command-line flag parsing. 

## Usage

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

