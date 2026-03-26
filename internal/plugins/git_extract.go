package plugins

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// 🔥 Extrai blobs dos objetos git
func ExtractGit(target string) error {

	basePath := fmt.Sprintf("outputs/git/%s/.git/objects", sanitize(target))
	outputPath := fmt.Sprintf("outputs/git/%s/extracted", sanitize(target))

	os.MkdirAll(outputPath, os.ModePerm)

	filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {

		if err != nil || info == nil || info.IsDir() {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return nil
		}

		reader, err := zlib.NewReader(bytes.NewReader(data))
		if err != nil {
			return nil
		}

		decompressed, err := io.ReadAll(reader)
		reader.Close()

		if err != nil {
			return nil
		}

		// 🔥 Detecta blob
		if bytes.HasPrefix(decompressed, []byte("blob")) {

			idx := bytes.IndexByte(decompressed, 0)
			if idx < 0 {
				return nil
			}

			content := decompressed[idx+1:]

			fileName := filepath.Join(outputPath, filepath.Base(path))
			os.WriteFile(fileName, content, 0644)
		}

		return nil
	})

	return nil
}
