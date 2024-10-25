package repositories

import "net/http"

type Err string

func (e Err) Error() string {
	return string(e)
}

const (
	ErrInternal    Err = "internal error"
	ErrUnreachable Err = "unreachable"
)

func getGithubErr(response *http.Response) error {
	if response == nil {
		return ErrInternal
	}
	switch response.StatusCode {
	case 422:
		return ErrUnreachable
	default:
		return ErrInternal
	}
}
