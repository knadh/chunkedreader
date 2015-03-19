# Go chunkedreader

Kailash Nadh, March 2015

MIT License

## What?
chunkedreader is a light weight wrapper for Go's `bufio` that enables reading of byte streams in fixed size chunks.

It makes use of `bufio.Scanner`'s custom split function capability, which otherwise is meant for splitting
based on delimiters, to split based on fixed lengths.

This could especially be useful for reading continuous fixed length messages from a TCP stream, emulating "packets", for instance.

## Installation (go 1.1+)
`go get github.com/knadh/chunkedreader`

## Example
```go
package main

import (
	"github.com/knadh/chunkedreader"
)

func main() {
	// Open file for reading.
	f, _ := os.Open("test.txt")

	// Initialize the reader to read chunks of 4 bytes.
	ch := chunkedreader.New(f, 4)

	for ch.Read() {
		// Length of b will always be 4, or less than 4 if there
		// are no more bytes available to read.
		b := ch.Bytes()
	}
}
```
