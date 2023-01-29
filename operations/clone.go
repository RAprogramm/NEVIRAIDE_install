package operations

import (
	"log"
	"os/exec"
)

// clone_repo downloads NEVIRAIDE from git repository
func Clone_repo() {
	repo := "https://github.com/RAprogramm/NEVIRAIDE.git"
	place := "~/.config/nvim"
	println("Cloning...")
	cmd, err := exec.Command("git", "clone", repo, place).Output()
	if err != nil {
		log.Fatal(err)
	}
	println(string(cmd))
}
