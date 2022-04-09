# Log2u
[![GoDoc](https://godoc.org/github.com/victormagalhaess/log2u?status.svg)](http://godoc.org/github.com/victormagalhaess/log2u)
<p align="center">
  <img src="https://github.com/victormagalhaess/log2u/blob/main/public/log2u.png?raw=true" width="200" alt="Log2u Logo">
</p>

Log2u is a generalization of a log implementation that I used in a few projects before.
It aims to allow you to have a rich and structured log with flexiblity, allowing you to use different colors, date patterns
and even stack trace for messages.

<p align="center">
  <img src="https://github.com/victormagalhaess/log2u/blob/main/public/example.png?raw=true" alt="Log2u logs on terminal">
</p>

# Usage

You can add log2u to your project simply running:
```sh
go get github.com/victormagalhaess/log2u
```

You can see and example of usage [here](https://github.com/victormagalhaess/log2u/blob/main/example/example.go)

# Build

log2u is built in go, and can be built for your machine, however, it was intended to be used as a package for go 1.13 or superior.

# Credits

The log2u logo was made using Flaticon free icons:
<a href="https://www.flaticon.com/br/icones-gratis/arquivos-e-pastas" title="arquivos e pastas ícones">Files and folders icons by manshagraphics - Flaticon</a>

I would like to really thanks to [@apsdehal](github.com/apsdehal) who heavily inspired this package with his [go-logger](https://github.com/apsdehal/go-logge) project.

