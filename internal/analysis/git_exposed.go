package analysis

func RunGitExposed(target string) []Finding {

	return []Finding{
		{
			Title:       "Git Repository Exposed",
			Description: ".git/config accessible",
			Severity:    High,
			Confidence:  90,
			Target:      target,
		},
	}
}

