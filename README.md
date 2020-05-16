# go-bitindex
> The unofficial Go implementation for the [BitIndex API](https://www.bitindex.network/developers/api-documentation-v3.html)

<br/>

[![Release](https://img.shields.io/github/release-pre/mrz1836/go-bitindex.svg?logo=github&style=flat&v=2)](https://github.com/mrz1836/go-bitindex/releases)
[![Build Status](https://travis-ci.com/mrz1836/go-bitindex.svg?branch=master&v=1)](https://travis-ci.com/mrz1836/go-bitindex)
[![Report](https://goreportcard.com/badge/github.com/mrz1836/go-bitindex?style=flat&v=1)](https://goreportcard.com/report/github.com/mrz1836/go-bitindex)
[![Go](https://img.shields.io/github/go-mod/go-version/mrz1836/go-bitindex)](https://golang.org/)
[![Sponsor](https://img.shields.io/badge/sponsor-MrZ-181717.svg?logo=github&style=flat&v=3)](https://github.com/sponsors/mrz1836)
[![Donate](https://img.shields.io/badge/donate-bitcoin-ff9900.svg?logo=bitcoin&style=flat)](https://mrz1818.com/?tab=tips&af=go-bitindex)

<br/>

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

<br/>

## Installation

**go-bitindex** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```shell script
go get -u github.com/mrz1836/go-bitindex
```

<br/>

## Documentation
View the generated [documentation](https://pkg.go.dev/github.com/mrz1836/go-bitindex)

[![GoDoc](https://godoc.org/github.com/mrz1836/go-bitindex?status.svg&style=flat)](https://pkg.go.dev/github.com/mrz1836/go-bitindex)

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

<details>
<summary><strong><code>Library Deployment</code></strong></summary>
<br/>

[goreleaser](https://github.com/goreleaser/goreleaser) for easy binary or library deployment to Github and can be installed via: `brew install goreleaser`.

The [.goreleaser.yml](.goreleaser.yml) file is used to configure [goreleaser](https://github.com/goreleaser/goreleaser).

Use `make release-snap` to create a snapshot version of the release, and finally `make release` to ship to production.
</details>

<details>
<summary><strong><code>Makefile Commands</code></strong></summary>
<br/>

View all `makefile` commands
```shell script
make help
```

List of all current commands:
```text
clean                  Remove previous builds and any test cache data
clean-mods             Remove all the Go mod cache
coverage               Shows the test coverage
godocs                 Sync the latest tag with GoDocs
help                   Show this help message
install                Install the application
install-go             Install the application (Using Native Go)
lint                   Run the Go lint application
release                Full production release (creates release in Github)
release                Runs common.release then runs godocs
release-snap           Test the full release (build binaries)
release-test           Full production test release (everything except deploy)
replace-version        Replaces the version in HTML/JS (pre-deploy)
run-examples           Runs all the examples
tag                    Generate a new tag and push (tag version=0.0.0)
tag-remove             Remove a tag if found (tag-remove version=0.0.0)
tag-update             Update an existing tag to current commit (tag-update version=0.0.0)
test                   Runs vet, lint and ALL tests
test-short             Runs vet, lint and tests (excludes integration tests)
test-travis            Runs all tests via Travis (also exports coverage)
test-travis-short      Runs unit tests via Travis (also exports coverage)
uninstall              Uninstall the application (and remove files)
vet                    Run the Go vet application
```
</details>

<br/>

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
```shell script
make test
```

Run tests (excluding integration tests)
```shell script
make test-short
```

<br/>

## Benchmarks
Run the Go [benchmarks](bitindex_test.go):
```shell script
make bench
```

<br/>

## Code Standards
Read more about this Go project's [code standards](CODE_STANDARDS.md).

<br/>

## Usage
View the [bitindex examples](#examples--tests) above

Basic implementation:
```go
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

<br/>

## Maintainers
| [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) |
|:---:|
| [MrZ](https://github.com/mrz1836) |

<br/>

## Contributing
View the [contributing guidelines](CONTRIBUTING.md) and follow the [code of conduct](CODE_OF_CONDUCT.md).

### How can I help?
All kinds of contributions are welcome :raised_hands:! 
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:. 
You can also support this project by [becoming a sponsor on GitHub](https://github.com/sponsors/mrz1836) :clap: 
or by making a [**bitcoin donation**](https://mrz1818.com/?tab=tips&af=go-bitindex) to ensure this journey continues indefinitely! :rocket:


#### Credits
[@Attila](https://github.com/attilaaf) & [BitIndex](https://www.bitindex.network/) for their hard work on the [BitIndex API](https://www.bitindex.network/developers/api-documentation-v3.html)

Looking for a Javascript version? Check out the [BitIndex JS SDK](https://github.com/BitIndex/bitindex-sdk)

Looking for MatterCloud? Checkout the [go-mattercloud](https://github.com/mrz1836/go-mattercloud) package.

<br/>

## License

![License](https://img.shields.io/github/license/mrz1836/go-bitindex.svg?style=flat&v=1)