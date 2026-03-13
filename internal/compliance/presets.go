package compliance

var HackerOne = CompliancePolicy{
	Name:      "hackerone",
	UserAgent: "amazonvrpresearcher_yourh1username",
	Allowed:   true,
}

var GenericBugBounty = CompliancePolicy{
	Name:      "generic-bugbounty",
	UserAgent: "bugbounty-researcher",
	Allowed:   true,
}
