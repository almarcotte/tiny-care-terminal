package git

import (
	"strings"
	"testing"
)

func TestToRepositories(t *testing.T) {
	paths := []string{
		"~/Projects/Top-Secret-Project",
		"~/Projects/Website",
		"/Go/src/github.com/gnumast/tiny-care-terminal",
	}

	repos := ToRepositories(strings.Join(paths, ";"))

	for i, repo := range repos {
		if repo.Path != paths[i] {
			t.Fatalf("Expected repository %d to have path %s. Got %s", i, paths[i], repo.Path)
		}
	}
}
