package file

import (
	"strings"
	"io"
	"os"
	"mime/multipart"
)

func Upload(path string, fl *multipart.FileHeader) (bool, error) {
	// Source
	src, err := fl.Open()

	defer src.Close()

	path = strings.Join([]string{path, "pdf"}, ".")

	// Destination
	dst, err := os.Create(path)

	defer dst.Close()

	// Copy
	 _, err = io.Copy(dst, src)

	return true, err
}