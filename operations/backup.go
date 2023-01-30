// package operations consist of functions wich make installation process easy
package operations

import (
	"fmt"
	"log"
	"os/exec"
)

// Backup_old function save old neovim config folder
func Backup_old() {
	before := "~/.config/nvim"
	after := "~/.config/nvim.old"
	cmd, err := exec.Command("mv", before, after).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(cmd))
}
