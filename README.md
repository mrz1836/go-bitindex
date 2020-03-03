# go-bitindex
**go-bitindex** is the unofficial golang implementation for the [BitIndex API](https://www.bitindex.network/developers/api-documentation-v3.html)

[![Go](https://img.shields.io/github/go-mod/go-version/mrz1836/go-bitindex)](https://golang.org/)
[![Build Status](https://travis-ci.com/mrz1836/go-bitindex.svg?branch=master&v=1)](https://travis-ci.com/mrz1836/go-bitindex)
[![Report](https://goreportcard.com/badge/github.com/mrz1836/go-bitindex?style=flat&v=1)](https://goreportcard.com/report/github.com/mrz1836/go-bitindex)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/dde6d46426bd4c12be65916da8cf04d2)](https://www.codacy.com/app/mrz1818/go-bitindex?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=mrz1836/go-bitindex&amp;utm_campaign=Badge_Grade)
[![Release](https://img.shields.io/github/release-pre/mrz1836/go-bitindex.svg?style=flat&v=2)](https://github.com/mrz1836/go-bitindex/releases)
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

**go-bitindex** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```bash
$ go get -u github.com/mrz1836/go-bitindex
```

## Documentation
You can view the generated [documentation here](https://godoc.org/github.com/mrz1836/go-bitindex).

You can also view the [BitIndex api](https://www.bitindex.network/developers/api-documentation-v3.html) documentation.

### Features
- Supports >= V3 API requests
- [Client](client.go) is completely configurable
- Customize the network per request (`main`, `test` or `stn`)
- Using [heimdall http client](https://github.com/gojek/heimdall) with exponential backoff & more
- Current (V3) coverage for the [BitIndex](https://developers.bitindex.com/) API
    - [x] Address
    - [x] Block
    - [x] Chain Info
    - [x] Transaction
    - [x] Webhooks
    - [x] Xpub

## Examples & Tests
All unit tests and [examples](bitindex_test.go) run via [Travis CI](https://travis-ci.org/mrz1836/go-bitindex) and uses [Go version 1.14.x](https://golang.org/doc/go1.14). View the [deployment configuration file](.travis.yml).

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
- View the [bitindex examples](#examples--tests) above

Basic implementation:
```golang
package main

import (
	"log"

	"github.com/mrz1836/go-bitindex"
)

func main() {

	// Create a new client
	client, _ := bitindex.NewClient("your-secret-api-key", bitindex.NetworkMain, nil)

	// Get balance for an address
	info, _ := client.AddressInfo("16ZqP5Tb22KJuvSAbjNkoiZs13mmRmexZA")

	// What's the balance?
	log.Println("address balance:", info.Balance)
}
```

## Maintainers

| [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) |
|:---:|
| [MrZ](https://github.com/mrz1836) |

## Contributing

View the [contributing guidelines](CONTRIBUTING.md) and follow the [code of conduct](CODE_OF_CONDUCT.md).

Support the development of this project 🙏

[![Donate](https://img.shields.io/badge/donate-bitcoin-brightgreen.svg)](https://mrz1818.com/?tab=tips&af=go-bitindex)

#### Credits

[@Attila](https://github.com/attilaaf) & [BitIndex](https://www.bitindex.network/) for their hard work on the [BitIndex API](https://www.bitindex.network/developers/api-documentation-v3.html)

Looking for a Javascript version? Check out the [BitIndex JS SDK](https://github.com/BitIndex/bitindex-sdk)

Looking for Matter Cloud? Checkout the [go-mattercloud](https://github.com/mrz1836/go-mattercloud) package.

## License

![License](https://img.shields.io/github/license/mrz1836/go-bitindex.svg?style=flat&v=1)