package finding

type Severity string

const (
	Info     Severity = "info"
	Low      Severity = "low"
	Medium   Severity = "medium"
	High     Severity = "high"
	Critical Severity = "critical"
)

type Finding struct {
	ID          string
	Title       string
	Category    string
	Severity    Severity
	Confidence  int // 0–100
	Target      string
	Host        string
	Path        string
	Evidence    string
	Recommendation string
}
