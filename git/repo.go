package git

import (
	"os/exec"
)

// Repo basic struct of a git repository
type Repo struct {
	Repo     string
	Dir      string
	Name     string
	FullName string
}

// Clone clones the git project
func (r *Repo) Clone() (int, error) {
	cmd := exec.Command("git", "clone", r.Repo, r.Dir)
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode(), err
		}
		return -1, err
	}
	return 0, nil
}

// Pull pulls the git project
func (r *Repo) Pull() (int, error) {
	cmd := exec.Command("git", "pull")
	cmd.Dir = r.Dir
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode(), err
		}
		return -1, err
	}
	return 0, nil
}
