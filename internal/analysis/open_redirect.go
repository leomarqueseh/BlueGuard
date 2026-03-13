package analysis

func RunOpenRedirect(target string) []Finding {

	return []Finding{
		{
			Title:       "Possible Open Redirect",
			Description: "Parameter may allow redirection",
			Severity:    Medium,
			Confidence:  60,
			Target:      target,
		},
	}
}
