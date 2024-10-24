package query

type RepositoryDTO struct {
	FullName       string         `json:"full_name"`
	Owner          string         `json:"owner"`
	RepositoryName string         `json:"repository_name"`
	Language       map[string]int `json:"languages"`
}
