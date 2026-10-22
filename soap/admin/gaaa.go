package admin

const adminAccountsQuery = "(|(zimbraIsAdminAccount=TRUE)(zimbraIsDelegatedAdminAccount=TRUE)(zimbraIsDomainAdminAccount=TRUE))"

func GetAllAdminAccounts(conf map[string]string, verbose bool, attributes []string) error {
	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	return searchAccounts(zcs, SearchOptions{
		Query:         adminAccountsQuery,
		SortBy:        "name",
		SortAscending: true,
		Verbose:       verbose || len(attributes) > 0,
		Attributes:    attributes,
	})
}
