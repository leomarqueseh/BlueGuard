package analysis

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var LoadedFingerprints []Fingerprint

func LoadFingerprints(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) != ".yaml" {
			continue
		}

		data, err := os.ReadFile(filepath.Join(dir, f.Name()))
		if err != nil {
			continue
		}

		var fp Fingerprint
		if err := yaml.Unmarshal(data, &fp); err != nil {
			continue
		}

		LoadedFingerprints = append(LoadedFingerprints, fp)
	}

	return nil
}
