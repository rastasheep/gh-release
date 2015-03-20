package github

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Release struct {
	Repository *Repository
	Path       string
	Files      []string
	Version    string
}

func CreateRelease(slug string, version string, path string) (*Release, error) {
	repo, err := CreateRepository(slug)
	if err != nil {
		return nil, err
	}

	files, err := artifacts(path)
	if err != nil {
		return nil, err
	}

	return &Release{
		Repository: repo,
		Files:      files,
		Path:       path,
		Version:    version,
	}, nil
}

func (r *Release) Deploy() error {
	return nil
}

func artifacts(path string) ([]string, error) {
	result := make([]string, 0)

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			result = append(result, path)
		}

		return nil
	})

	if err == nil && len(result) == 0 {
		err = errors.New(fmt.Sprintf("Release dir: '%v' empty.", path))
	}

	return result, err
}
