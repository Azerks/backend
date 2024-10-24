package query

type RepositoryDTO struct {
	FullName       string         `json:"full_name"`
	Owner          string         `json:"owner"`
	RepositoryName string         `json:"repository_name"`
	License        string         `json:"license"`
	Language       map[string]int `json:"languages"`
}

type RepositoriesFilters struct {
	Limit    int
	Language string
	License  string
}
