package detectlibc

import "runtime"

// A string constant containing the value "glibc".
const Glibc = "glibc"
// A string constant containing the value "musl".
const Musl = "musl"

// Returns:
//  - false when the libc family is glibc
//  - true when the libc family is not glibc
//  - false when run on a non-Linux platform
func IsNonGlibcLinux() bool {
	if runtime.GOOS != "linux" {
		return false
	}
	libc, ok := Family()
	if !ok {
		return false
	}
	return libc != "glibc"
}
