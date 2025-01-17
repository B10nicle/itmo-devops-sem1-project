package archiver

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"path/filepath"
)

func ExtractFromZip(zipStream io.Reader) (io.ReadCloser, error) {
	data, err := io.ReadAll(zipStream)
	if err != nil {
		return nil, fmt.Errorf("failed to read zip stream: %w", err)
	}

	zipReader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, fmt.Errorf("failed to open zip archiver: %w", err)
	}

	for _, file := range zipReader.File {
		if filepath.Ext(file.Name) != ".csv" || filepath.Base(file.Name)[0] == '.' {
			continue
		}
		rc, err := file.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file %s in zip archiver: %w", file.Name, err)
		}
		return rc, nil
	}

	return nil, errors.New("no valid CSV file found in the zip archiver")
}

func ZipFile(content *bytes.Buffer, output io.Writer, fileName string) error {
	zipWriter := zip.NewWriter(output)
	defer func(zipWriter *zip.Writer) {
		err := zipWriter.Close()
		if err != nil {

		}
	}(zipWriter)

	file, err := zipWriter.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file %s in zip archiver: %w", fileName, err)
	}

	_, err = file.Write(content.Bytes())
	if err != nil {
		return fmt.Errorf("failed to write content to file %s in zip archiver: %w", fileName, err)
	}

	return nil
}
