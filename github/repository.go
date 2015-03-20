package github

import (
	"errors"
	"strings"
)

type Repository struct {
	Owner string
	Name  string
}

func CreateRepository(fullName string) (*Repository, error) {
	s := strings.Split(fullName, "/")
	if len(s) < 2 {
		return nil, errors.New("Misformed repository name")
	}

	return &Repository{
		Owner: s[0],
		Name:  s[1],
	}, nil
}
