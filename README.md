# detect-libc for Go

⚙️ Detect glibc or musl

<table align=center><td>

<div><sub><i>On Linux</i></sub></div>

```go
libc, ok := detectlibc.Family()
if !ok {
    log.Fatal("could not detect libc")
}
log.Printf("libc=%s", libc)
// Possible output: libc=glibc
// Possible output: libc=musl
```

</table>

## Installation

```sh
go get github.com/jcbhmr/go-detectlibc
```

## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/jcbhmr/go-detectlibc"
)

func main() {
	libc, ok := detectlibc.Family()
	if !ok {
		log.Fatal("libc family could not be determined")
	}
	fmt.Printf("libc=%s", libc)
}
```

### What about detecting GNU or MSVC on Windows?

```go
//go:build windows
package mydetectlibc
func Family() (string, bool) {
    cmd := exec.Command("uname", "-s")
    out, err := cmd.Output()
    if err == nil {
        if bytes.HasPrefix(out, []byte("MINGW")) || bytes.HasPrefix(out, []byte("MSYS")) || bytes.HasPrefix(out, []byte("CYGWIN")) || bytes.Equal(out, []byte("Windows_NT\n")) {
            return "gnu", true
        }
    }
    
    // Assume MSVC
    return "msvc", true
}
```
