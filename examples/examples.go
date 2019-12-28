/*
Package main is an example package that uses go-bitindex
*/
package main

import (
	"log"
	"os"

	"github.com/mrz1836/go-bitindex"
)

// main is an example of using this go-bitindex package
func main() {

	// Set your api key
	yourAPIKey := os.Getenv("BITINDEX_API_KEY")

	// Create a new client
	client, err := bitindex.NewClient(yourAPIKey, bitindex.NetworkMain, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Set test net?
	//client.Parameters.Network = bitindex.NetworkTest

	// Get balance for an address
	var info *bitindex.AddressInfo
	info, err = client.AddressInfo("16ZqP5Tb22KJuvSAbjNkoiZs13mmRmexZA")
	if err != nil {
		log.Fatal(err.Error())
	}

	// What's the balance?
	log.Println("address balance:", info.Balance)
}
