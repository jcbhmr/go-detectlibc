//go:build !linux

package detectlibc

// Returns: "", false
func Family() (string, bool) {
	return "", false
}

// Returns: "", false
func Version() (string, bool) {
	return "", false
}
