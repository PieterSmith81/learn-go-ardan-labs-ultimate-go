package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println(SHA1Sig("http.log.gz"))
	fmt.Println(SHA1Sig("sha1.go")) // This file.
}

// SHA1Sig returns the SHA1 signature of an uncompressed gzip file.
// So, similar to the output returned by this terminal command: cat http.log.gz | gunzip | sha1sum
func SHA1Sig(fileName string) (string, error) {
	// Open the compressed file (which implicitly implements the io.Reader interface).
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Uncompress and read the file (if it's a gzip file).
	// Note the use of the io.Reader interface here again, this time as a parameter for the gzip.NewReader() function.
	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {
		gz, err := gzip.NewReader(file)
		if err != nil {
			return "", fmt.Errorf("%q - gzip: %w", fileName, err)
		}
		defer gz.Close()
		r = gz
	}

	// Get the SHA1 checksum value for the uncompressed file.
	// This time around, we are using the io.Writer interface.
	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", fmt.Errorf("%q - copy: %w", fileName, err)
	}

	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}
