package github

import (
	"os"
	"testing"
)

var pwd, _ = os.Getwd()
var tmpDir = pwd + "/tmp"

func TestCreateRelease(t *testing.T) {
	release, err := CreateRelease("rastasheep/todo", "1.0.0", pwd)

	if release.Path != pwd {
		t.Errorf("release should have a correct path, got %v", release.Path)
	}

	if release.Version != "1.0.0" {
		t.Errorf("release should have a correct version, got %v", release.Version)
	}

	if err != nil {
		t.Errorf("release should be returned, got %v", err)
	}
}

func TestCreateReleaseWithWrongSlug(t *testing.T) {
	_, err := CreateRelease("rastasheep", "1.0.0", "tmp")

	if err == nil {
		t.Errorf("error should be returned")
	}

	if err.Error() != "Misformed repository name" {
		t.Errorf("error message should be properly set, got '%v'", err.Error())
	}
}

func TestCreateReleaseWithEmptyReleaseDir(t *testing.T) {
	os.Mkdir(tmpDir, 0755)
	_, err := CreateRelease("rastasheep/todo", "1.0.0", tmpDir)

	if err == nil {
		t.Errorf("error should be returned")
	}

	if err.Error() != "Release dir: '"+tmpDir+"' empty." {
		t.Errorf("error message should be properly set")
	}
}

func TestArtifacts(t *testing.T) {
	artifacts, err := artifacts(pwd)

	if len(artifacts) == 0 {
		t.Errorf("artifacts should be presetn")
	}

	if err != nil {
		t.Errorf("artifacts should be returned, got %v", err)
	}
}

func TestDeploy(t *testing.T) {
	release, _ := CreateRelease("rastasheep/todo", "1.0.0", pwd)
	deploy := release.Deploy()
	if deploy != nil {
		t.Errorf("there should not be an errors")
	}
}

func TestArtifactsEmptyDir(t *testing.T) {
	os.Mkdir(tmpDir, 0755)

	_, err := artifacts(tmpDir)

	if err == nil {
		t.Errorf("error should be returned")
	}

	if err.Error() != "Release dir: '"+tmpDir+"' empty." {
		t.Errorf("error message should be properly set")
	}
}

func TestArtifactsNonexistentDir(t *testing.T) {
	_, err := artifacts("test")

	if err == nil {
		t.Errorf("error should be returned")
	}

	if err.Error() != "lstat test: no such file or directory" {
		t.Errorf("error message should be properly set")
	}
}
