# Go - Iterator

[![Go Report Card](https://goreportcard.com/badge/github.com/danielgatis/go-iterator?style=flat-square)](https://goreportcard.com/report/github.com/danielgatis/go-iterator)
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/danielgatis/go-iterator/master/LICENSE)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/danielgatis/go-iterator)

This package provides an iterator interface and a slice iterator implementation.

## Install

```bash
go get -u github.com/danielgatis/go-iterator
```

And then import the package in your code:

```go
import "github.com/danielgatis/go-iterator"
```

### Example

An example described below is one of the use cases.

```go
package main

import (
	"fmt"

	"github.com/danielgatis/go-iterator"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	iter := iterator.NewIterator(numbers)

	for iter.HasNext() {
		value := iter.GetNextOrDefault(0)
		fmt.Println(value)
	}
}

```

```
❯ go run main.go
```

### License

Copyright (c) 2023-present [Daniel Gatis](https://github.com/danielgatis)

Licensed under [MIT License](./LICENSE)

### Buy me a coffee

Liked some of my work? Buy me a coffee (or more likely a beer)

<a href="https://www.buymeacoffee.com/danielgatis" target="_blank"><img src="https://bmc-cdn.nyc3.digitaloceanspaces.com/BMC-button-images/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;"></a>
