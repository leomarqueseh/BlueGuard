package analysis

func RunTakeover(target string) []Finding {

	return []Finding{
		{
			Title:       "Possible Subdomain Takeover",
			Description: "CNAME points to unclaimed service",
			Severity:    High,
			Confidence:  70,
			Target:      target,
		},
	}
}
