package archiver

import (
	"archive/tar"
	"bytes"
	"errors"
	"fmt"
	"io"
	"path/filepath"
)

func ExtractTar(tarStream io.Reader) (io.ReadCloser, error) {
	data, err := io.ReadAll(tarStream)
	if err != nil {
		return nil, fmt.Errorf("failed to read input stream: %w", err)
	}

	tarReader := tar.NewReader(bytes.NewReader(data))

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading tar archiver: %w", err)
		}

		if filepath.Ext(header.Name) != ".csv" || filepath.Base(header.Name)[0] == '.' {
			continue
		}

		return io.NopCloser(tarReader), nil
	}

	return nil, errors.New("no valid CSV file found in the tar archiver")
}
