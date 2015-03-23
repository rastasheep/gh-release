package github

import "testing"

func TestClient_newOctokitClient(t *testing.T) {
	c := NewClient("https://api.github.com")
	cc := c.api()

	if cc.Endpoint.String() != "https://api.github.com" {
		t.Errorf("endpoint should be properly set, got '%v'", cc.Endpoint.String())
	}

	c = NewClient("https://github.corporate.com")
	cc = c.api()

	if cc.Endpoint.String() != "https://github.corporate.com" {
		t.Errorf("endpoint should be properly set, got '%v'", cc.Endpoint.String())
	}

}
