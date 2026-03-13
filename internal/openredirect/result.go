package openredirect

type Finding struct {
	URL       string
	Mode      string // passive | standard
	Payload   string
	Evidence  string
	Severity  string
}
