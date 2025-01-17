//go:build linux

package detectlibc_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jcbhmr/go-detectlibc"
)

func TestMain(m *testing.M) {
	if libc, _ := detectlibc.Family(); libc != "glibc" {
		os.Args = append(os.Args, "-test.skip=ExampleFamily")
	}
	os.Exit(m.Run())
}

func ExampleFamily() {
	libc, ok := detectlibc.Family()
	if !ok {
		log.Fatal("libc family could not be determined")
	}
	fmt.Printf("libc=%s", libc)
	// Output: libc=glibc
}
