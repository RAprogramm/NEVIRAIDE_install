// package operations consist of functions wich make installation process easy
package operations

import (
	"log"
	"os/exec"
	"strings"
)

// Detect_distro function return name of the package manager using by default in your distro
func Detect_manager() string {
	var distro_name string
	cmd, err := exec.Command("uname", "-a").Output()
	if err != nil {
		log.Fatal(err)
	}
	switch strings.ContainsAny(string(cmd), distro_name) {
	case distro_name == "ubuntu":
		return "apt"
	case distro_name == "manjaro":
		return "pamac"
	default:
		return "pacman"
	}
}
