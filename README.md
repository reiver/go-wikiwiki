# go-wikiwiki

Package **wikiwiki** provides tools for working with wiki text, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-wikiwiki

[![GoDoc](https://godoc.org/github.com/reiver/go-wikiwiki?status.svg)](https://godoc.org/github.com/reiver/go-wikiwiki)

## Wiki Text Syntax

This is the syntax wiki text markup:

```
''superscript'
((no-op))
**bold**
,,subscript,,
//italics//
::smart-code::
[[internal-link]]
__underline__
``mononspaced``
{{name}}
||highlight||
‾‾overline‾‾
```

## Import

To import package **wikiwiki** use `import` code like the follownig:
```
import "github.com/reiver/go-wikiwiki"
```

## Installation

To install package **wikiwiki** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-wikiwiki
```

## Author

Package **wikiwiki** was written by [Charles Iliya Krempeaux](http://reiver.link)
