package analysis

import "strings"

// DetectFingerprintByCNAME tenta identificar takeover via CNAME
func DetectFingerprintByCNAME(cname string) *Fingerprint {
	cname = strings.ToLower(cname)

	for _, fp := range LoadedFingerprints {
		for _, indicator := range fp.CNAME {
			if strings.Contains(cname, strings.ToLower(indicator)) {
				return &fp
			}
		}
	}
	return nil
}
