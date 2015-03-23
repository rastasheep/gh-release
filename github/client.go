package github

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/octokit/go-octokit/octokit"
)

const (
	userAgent    = "gh-release"
	gitHubAPIURL = "https://api.github.com"
)

type Client struct {
	Host        string
	AccessToken string
}

func NewClient(token string) *Client {
	return &Client{
		Host:        gitHubAPIURL,
		AccessToken: token,
	}
}

func (client *Client) CreateRelease(release *Release) (*Release, error) {
	api := client.api()
	url := relesesUrl(release.Repository)

	rel, result := api.Releases(url).Create(params(release))
	release.Release = rel

	if result.HasError() {
		return nil, errors.New(fmt.Sprintf("Error creating relese: '%v'.", result.Err))
	}

	return release, nil
}

func (client *Client) api() *octokit.Client {
	auth := octokit.TokenAuth{AccessToken: client.AccessToken}

	return octokit.NewClientWith(client.Host, userAgent, auth, nil)
}

func params(r *Release) *octokit.ReleaseParams {
	return &octokit.ReleaseParams{
		TagName:         r.Version,
		TargetCommitish: "master",
		// TODO: enchant release with additional info
		//Name:            title,
		//Body:            body,
		//Draft:           flagReleaseDraft,
		//Prerelease:      flagReleasePrerelease,
	}
}

func relesesUrl(r *Repository) *url.URL {
	url, _ := octokit.ReleasesURL.Expand(octokit.M{"owner": r.Owner, "repo": r.Name})
	return url
}
