# OSDU-Types-Go

## Description

OSDU-Types-go is a Go library for containing all current OSDU (Open Subsurface Data Universe) well known schemas. There's an action set up to update this repo at a set interval if changes occur.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)

## Installation

To install OSDU-Types-Go, you need to have Go installed on your machine. You can then use `go get` to fetch the library.

```bash
go get github.com/Frelsaren/osdu-client-go
```

## Usage

```go
package main

import (
	"github.com/Frelsaren/osdu-types-go/masterdata"
)

func main() {

	field, err := masterdata.UnmarshalField(fieldJsonAsBytes)
	// or newField := masterdata.Field{ ... }

	doStuffWithFields(&field)
}
```
