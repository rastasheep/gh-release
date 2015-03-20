package github

import "testing"

func TestCreateRepository(t *testing.T) {
	repo, err := CreateRepository("rastasheep/todo")

	if err != nil {
		t.Errorf("repository should be crated")
	}

	if repo.Owner != "rastasheep" {
		t.Errorf("repository owner should be set")
	}

	if repo.Name != "todo" {
		t.Errorf("repository name should be set")
	}
}

func TestCreateRepositoryMisformed(t *testing.T) {
	_, err := CreateRepository("rastasheeptodo")

	if err == nil {
		t.Errorf("repository should not be crated")
	}

	if err.Error() != "Misformed repository name" {
		t.Errorf("error message should be properly set")
	}
}
