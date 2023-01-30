// package operations consist of functions wich make installation process easy
package operations

import (
	"fmt"
	"log"
	"os/exec"
)

// clone_repo downloads NEVIRAIDE from git repository
func Clone_repo() {
	repo := "https://github.com/RAprogramm/NEVIRAIDE.git"
	place := "~/.config/nvim"
	fmt.Println("Cloning...")
	cmd, err := exec.Command("git", "clone", repo, place).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(cmd))
}
