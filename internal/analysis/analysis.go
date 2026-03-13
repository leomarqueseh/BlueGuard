package analysis

func RunAll(target string) []Finding {

	var results []Finding

	results = append(results, RunGitExposed(target)...)
	results = append(results, RunTakeover(target)...)
	results = append(results, RunOpenRedirect(target)...)

	return results
}
