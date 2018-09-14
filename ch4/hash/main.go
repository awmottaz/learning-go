// Hash prints the SHA256 hash of its input, with flags to support SHA384 or
// SHA512 algorithms.

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
	"strings"
)

var algorithm string

func init() {
	flag.StringVar(&algorithm, "a", "sha256", "which algorithm to use, one of sha256, sha384, or sha512.")
	flag.Parse()
}

func main() {
	in := []byte(strings.Join(flag.Args(), ""))

	switch algorithm {
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256(in))
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384(in))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512(in))
	default:
		flag.Usage()
		os.Exit(1)
	}

}
