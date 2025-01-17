package archiver

import "io"

func ExtractFile(r io.Reader, fileType string) (io.ReadCloser, error) {
	switch fileType {
	case ZIP:
		return ExtractZip(r)
	case TAR:
		return ExtractTar(r)
	default:
		return ExtractZip(r)
	}
}
