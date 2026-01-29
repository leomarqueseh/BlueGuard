package analysis

type Fingerprint struct {
	ID       string   `yaml:"id"`
	Provider string   `yaml:"provider"`
	CNAME    []string `yaml:"cname"`

	HTTP struct {
		Matchers  []Matcher `yaml:"matchers"`
		Negative  []Matcher `yaml:"negative"`
	} `yaml:"http"`
}

type Matcher struct {
	Type    string   `yaml:"type"`   // body, header, status
	Words   []string `yaml:"words"`
	Status  []int    `yaml:"status"`
}
