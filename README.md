[![Build Status](https://travis-ci.org/mamal72/gojib.svg?branch=master)](https://travis-ci.org/mamal72/gojib)
[![Go Report Card](https://goreportcard.com/badge/github.com/mamal72/gojib)](https://goreportcard.com/report/github.com/mamal72/gojib)
[![GoDoc](https://godoc.org/github.com/mamal72/gojib?status.svg)](https://godoc.org/github.com/mamal72/gojib)
[![license](https://img.shields.io/github/license/mamal72/gojib.svg)](https://github.com/mamal72/gojib/blob/master/LICENSE)


# gojib

A simple package to scrape and get price data from [Iranjib](https://www.iranjib.ir) website.


## Installation

```bash
go get github.com/mamal72/gojib
# or use dep, vgo, glide or anything else
```


## Usage

```go
package main

import (
	"log"

	"github.com/mamal72/gojib"
)

func main() {
	iranianCarsPrices, _ := gojib.GetIranianCarsPrices() // returns []CarPrice, error
	log.Println("Iranian Cars:")
	for _, car := range iranianCarsPrices {
		log.Printf("%+v", car)
	}

	foreignCarsPrices, _ := gojib.GetForeignCarsPrices() // returns []CarPrice, error
	log.Println("Foreign Cars:")
	for _, car := range foreignCarsPrices {
		log.Printf("%+v", car)
	}
}
```


## Ideas || Issues

Just create an issue and describe it. I'll check it ASAP!


## Contribution

You can fork the repository, improve or fix some part of it and then send the pull requests back (with some tests ofc) if you want to see them here. I really appreciate that. ❤️
