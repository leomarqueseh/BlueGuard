package analysis

import (
	"bytes"
)

func BodyContains(respBody []byte, pattern string) bool {
	return bytes.Contains(respBody, []byte(pattern))
}
