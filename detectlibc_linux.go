package detectlibc

import (
	"bytes"
	"os"
	"os/exec"
	"regexp"
	"sync"
)

var getconfigGnuLibcVersion = sync.OnceValues(func() ([]byte, error) {
	cmd := exec.Command("getconf", "GNU_LIBC_VERSION")
	return cmd.CombinedOutput()
})

var lddVersion = sync.OnceValues(func() ([]byte, error) {
	cmd := exec.Command("ldd", "--version")
	return cmd.CombinedOutput()
})

// Returns:
// - glibc or musl when the libc family can be determined
// - null when the libc family cannot be determined
var Family = sync.OnceValues(func() (string, bool) {
	ldd, err := os.ReadFile("/usr/bin/ldd")
	if err == nil {
		if bytes.Contains(ldd, []byte("musl")) {
			return "musl", true
		}
		if bytes.Contains(ldd, []byte("GNU C Library")) {
			return "glibc", true
		}
	}
	out, err := getconfigGnuLibcVersion()
	if err == nil {
		if bytes.Contains(bytes.TrimSpace(out), []byte("glibc")) {
			return "glibc", true
		}
	}
	out, err = lddVersion()
	if err == nil {
		if bytes.Contains(bytes.TrimSpace(out), []byte("musl")) {
			return "musl", true
		}
	}
	return "", false
})

// Returns:
// - The version when it can be determined
// - null when the libc family cannot be determined
var Version = sync.OnceValues(func() (string, bool) {
	ldd, err := os.ReadFile("/usr/bin/ldd")
	if err == nil {
		re := regexp.MustCompile(`(?i)LIBC[a-z0-9 \-).]*?(\d+\.\d+)`)
		match := re.FindSubmatch(ldd)
		if match != nil {
			return string(match[1]), true
		}
	}
	out, err := getconfigGnuLibcVersion()
	if err == nil {
		fields := bytes.Fields(bytes.TrimSpace(out))
		if len(fields) >= 2 {
			return string(fields[1]), true
		} else {
			return "", false
		}
	}
	out, err = lddVersion()
	if err == nil {
		lines := bytes.SplitN(out, []byte("\n"), 2)
		if len(lines) >= 2 {
			fields := bytes.Fields(lines[1])
			if len(fields) >= 2 {
				return string(fields[1]), true
			} else {
				return "", false
			}
		}
	}
	return "", false
})
