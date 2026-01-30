package analysis

import "strings"

func DetectFingerprint(cname string) *Fingerprint {
	for _, fp := range LoadedFingerprints {
		for _, c := range fp.CNAME {
			if strings.Contains(cname, c) {
				return &fp
			}
		}
	}
	return nil
}
