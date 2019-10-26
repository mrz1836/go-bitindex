package bitindex

import (
	"fmt"
	"os"
	"testing"
)

var (
	testAPIKey = os.Getenv("BITINDEX_API_KEY")
)

// TestNewClient test new client
func TestNewClient(t *testing.T) {
	// Not testing against bitindex
	if testing.Short() && len(testAPIKey) == 0 {
		testAPIKey = "dummy-key-not-testing"
	}
	client, err := NewClient(testAPIKey)
	if err != nil {
		t.Fatal(err)
	}

	if client.Parameters.Network != NetworkMain {
		t.Fatal("expected value to be main")
	}
}

// ExampleNewClient example using NewClient()
func ExampleNewClient() {
	client, _ := NewClient(testAPIKey)
	fmt.Println(client.Parameters.Network)
	// Output:main
}

// BenchmarkNewClient benchmarks the NewClient method
func BenchmarkNewClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewClient(testAPIKey)
	}
}
