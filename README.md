# A TinyPNG client written in Go

Compress a PNG file with the help of the [TinyPNG](https://tinypng.com/) service.

[![Build Status](https://travis-ci.org/peterhellberg/tinypng.png?branch=master)](https://travis-ci.org/peterhellberg/tinypng)
[![GoDoc](https://godoc.org/github.com/peterhellberg/tinypng?status.png)](https://godoc.org/github.com/peterhellberg/tinypng)

## Installation

Just go get the command:

    go get -u github.com/peterhellberg/tinypng/tinypng

## Usage

First you need to `export TINYPNG_API_KEY=yourTinyPNGApiKey`

Then you can run the command:

    tinypng <input.png> [output.png]

If only the input filename was specified, then the
output filename will be `tiny-<input.png>`

## License

Copyright (c) 2014 [Peter Hellberg](http://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
