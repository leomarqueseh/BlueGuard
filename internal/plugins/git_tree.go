package plugins

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ReconstructGit(target string) error {

	basePath := fmt.Sprintf("outputs/git/%s/.git/objects", sanitize(target))
	outputPath := fmt.Sprintf("outputs/git/%s/reconstructed", sanitize(target))

	os.MkdirAll(outputPath, os.ModePerm)

	var commitData []byte

	filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {

		if info == nil || info.IsDir() {
			return nil
		}

		data := readObject(path)

		if bytes.HasPrefix(data, []byte("commit")) {
			commitData = data
		}

		return nil
	})

	if commitData == nil {
		return nil
	}

	lines := strings.Split(string(commitData), "\n")

	treeHash := ""

	for _, line := range lines {
		if strings.HasPrefix(line, "tree ") {
			treeHash = strings.TrimSpace(strings.Split(line, " ")[1])
			break
		}
	}

	if treeHash == "" {
		return nil
	}

	parseTree(basePath, treeHash, outputPath)

	return nil
}

func parseTree(basePath, hash, outputPath string) {

	path := filepath.Join(basePath, hash[:2], hash[2:])
	data := readObject(path)

	idx := bytes.IndexByte(data, 0)
	content := data[idx+1:]

	for len(content) > 0 {

		spaceIdx := bytes.IndexByte(content, ' ')
		nullIdx := bytes.IndexByte(content, 0)

		if spaceIdx < 0 || nullIdx < 0 {
			break
		}

		name := string(content[spaceIdx+1 : nullIdx])
		hashBytes := content[nullIdx+1 : nullIdx+21]

		hash := fmt.Sprintf("%x", hashBytes)

		objPath := filepath.Join(basePath, hash[:2], hash[2:])
		objData := readObject(objPath)

		if bytes.HasPrefix(objData, []byte("blob")) {

			i := bytes.IndexByte(objData, 0)
			fileContent := objData[i+1:]

			fullPath := filepath.Join(outputPath, name)
			os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
			os.WriteFile(fullPath, fileContent, 0644)
		}

		if bytes.HasPrefix(objData, []byte("tree")) {
			parseTree(basePath, hash, filepath.Join(outputPath, name))
		}

		content = content[nullIdx+21:]
	}
}

func readObject(path string) []byte {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	reader, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil
	}

	out, err := io.ReadAll(reader)
	reader.Close()

	if err != nil {
		return nil
	}

	return out
}
