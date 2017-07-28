package git

import (
	"fmt"
	"os/exec"
	"strings"
)

type Repository struct {
	Path    string
	Commits []string
	Limit   int
	Format  string
}

func ToRepositories(paths string) []*Repository {
	if paths == "" {
		return make([]*Repository, 0)
	}

	repos := make([]*Repository, 0)

	for _, path := range strings.Split(paths, ";") {
		repos = append(repos, NewRepository(path, 0, ""))
	}

	return repos
}

func NewRepository(path string, limit int, format string) *Repository {
	repo := &Repository{Path: path}

	if format == "" {
		format = "[%h](fg-red) - %s [(%ad)](fg-green)"
	}

	repo.Format = format

	if limit == 0 {
		limit = 5
	}

	repo.Limit = limit

	repo.UpdateCommits()

	return repo
}

func (repo *Repository) UpdateCommits() {
	repo.Commits = getLastCommitsForDir(repo.Path, repo.Limit)
}

func getLastCommitsForDir(dir string, limit int) []string {
	out, err := exec.Command(
		"git",
		"-C",
		dir,
		"log",
		fmt.Sprintf("-%d", limit),
		"--date=relative",
		"--pretty=format:[%h](fg-red) - %s [(%ad)](fg-green)",
	).CombinedOutput()

	if err != nil {
		return []string{fmt.Sprintf("Couldn't fetch git log from [%s](fg-red). 🙀", dir)}
	}

	return strings.Split(string(out), "\n")
}
