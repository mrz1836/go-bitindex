# go-bitindex
**go-bitindex** is the unofficial golang implementation for the [bitindex API](https://www.bitindex.network/developers/api-documentation-v3.html)

[![Build Status](https://travis-ci.com/mrz1836/go-bitindex.svg?branch=master&v=2)](https://travis-ci.com/mrz1836/go-bitindex)
[![Report](https://goreportcard.com/badge/github.com/mrz1836/go-bitindex?style=flat&v=2)](https://goreportcard.com/report/github.com/mrz1836/go-bitindex)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/f9815e59758743b9adca25c11558ab1c)](https://www.codacy.com/app/mrz1818/go-bitindex?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=mrz1836/go-bitindex&amp;utm_campaign=Badge_Grade)
[![Release](https://img.shields.io/github/release-pre/mrz1836/go-bitindex.svg?style=flat&v=1)](https://github.com/mrz1836/go-bitindex/releases)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat)](https://github.com/RichardLitt/standard-readme)
[![GoDoc](https://godoc.org/github.com/mrz1836/go-bitindex?status.svg&style=flat)](https://godoc.org/github.com/mrz1836/go-bitindex)

## Table of Contents
- [Installation](#installation)
- [Documentation](#documentation)
- [Examples & Tests](#examples--tests)
- [Benchmarks](#benchmarks)
- [Code Standards](#code-standards)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

## Installation

**go-bitindex** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy) and [dep](https://github.com/golang/dep).
```bash
$ go get -u github.com/mrz1836/go-bitindex
```

Updating dependencies in **go-bitindex**:
```bash
$ cd ../go-bitindex
$ dep ensure -update -v
```

## Documentation
You can view the generated [documentation here](https://godoc.org/github.com/mrz1836/go-bitindex).

### Features
- Client is completely configurable
- Customize User Agent per request
- Customize the network per request (`main`, `test` or `stn`)
- Using [heimdall http client](https://github.com/gojek/heimdall) with exponential backoff & more
- Current coverage for the [bitindex.com](https://developers.bitindex.com/) API
    - [x] Address
    - [x] Block
    - [x] Chain Info
    - [x] Transaction
    - [x] Webhooks
    - [x] Xpub

## Examples & Tests
All unit tests and [examples](bitindex_test.go) run via [Travis CI](https://travis-ci.org/mrz1836/go-bitindex) and uses [Go version 1.13.x](https://golang.org/doc/go1.13). View the [deployment configuration file](.travis.yml).

Examples & Tests by API section:
- [address](address_test.go)
- [block](block_test.go)
- [chain](chain_test.go)
- [transaction](transaction_test.go)
- [webhook](webhook_test.go)
- [xpub](xpub_test.go)

Run all tests (including integration tests)
```bash
$ cd ../go-bitindex
$ go test ./... -v
```

Run tests (excluding integration tests)
```bash
$ cd ../go-bitindex
$ go test ./... -v -test.short
```

## Benchmarks
Run the Go [benchmarks](bitindex_test.go):
```bash
$ cd ../go-bitindex
$ go test -bench . -benchmem
```

## Code Standards
Read more about this Go project's [code standards](CODE_STANDARDS.md).

## Usage
- View the [bitindex examples](bitindex_test.go)

Basic implementation:
```golang
package main

import (
	"log"

	"github.com/mrz1836/go-bitindex"
)

func main() {

	// Create a new client
	client, _ := bitindex.NewClient("your-secret-api-key")

	// Get balance for an address
	info, _ := client.AddressInfo("16ZqP5Tb22KJuvSAbjNkoiZs13mmRmexZA")

	// What's the balance?
	log.Println("address balance:", info.Balance)
}
```

## Maintainers

[@MrZ](https://github.com/mrz1836)

## Contributing

View the [contributing guidelines](CONTRIBUTING.md) and follow the [code of conduct](CODE_OF_CONDUCT.md).

Support the development of this project 🙏

[![Donate](https://img.shields.io/badge/donate-bitcoin-brightgreen.svg)](https://mrz1818.com/?tab=tips&af=go-bitindex)

#### Credits

[@Attila](https://github.com/attilaaf) & [BitIndex](https://www.bitindex.network/) for their hard work on the [BitIndex API](https://www.bitindex.network/developers/api-documentation-v3.html)

## License

![License](https://img.shields.io/github/license/mrz1836/go-bitindex.svg?style=flat)