package git

import (
	"fmt"
	"os/exec"
	"strings"
)

// CloneBare clones given repository from its url to optionally defined path.
func CloneBare(url string, path string) error {
	cmd := exec.Command("git", "clone", "--bare", url, path)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git clone bare: %w, with output: %s", err, out)
	}
	return nil
}

// Pull calls `git pull --ff-only`
func Pull(path string) error {
	cmd := exec.Command("git", "pull", "--ff-only")
	cmd.Dir = path
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git pull: %w, with output: %s", err, out)
	}
	return nil
}

// Fetch calls `git fetch`
func Fetch(path string) error {
	cmd := exec.Command("git", "fetch")
	cmd.Dir = path
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git fetch: %w, with output: %s", err, out)
	}
	return nil
}

// BranchCreate checks out new branch from given base in given repository.
func BranchCreate(path, name, base string) error {
	cmd := exec.Command("git", "checkout", "-b", name, base)
	cmd.Dir = path
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git checkout -b: %w, with output: %s", err, out)
	}
	return nil
}

// WTAdd checks out new worktree of an existing branch at given path.
func WTAdd(repoPath, wtPath, branch string) error {
	cmd := exec.Command("git", "worktree", "add", wtPath, branch)
	cmd.Dir = repoPath
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git wt add: %w, with output: %s", err, out)
	}
	return nil
}

// WTRemove deletes worktree within repository. Branch is not deleted.
func WTRemove(repoPath, wtPath string) error {
	cmd := exec.Command("git", "worktree", "remove", wtPath)
	cmd.Dir = repoPath
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git wt remove: %w, with output: %s", err, out)
	}
	return nil
}

// WTList returns all worktrees of given repository
func WTList(path string) ([]string, error) {
	cmd := exec.Command("git", "worktree", "list")
	cmd.Dir = path
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("git wt list: %w, with output: %s", err, out)
	}
	lines := strings.Split(string(out), "\n")
	wts := make([]string, 0, len(lines))
	for _, l := range lines[:len(lines)-1] {
		wts = append(wts, strings.SplitN(l, " ", 2)[0])
	}
	return wts, nil
}
