package operations

import (
	"log"
	"os/exec"
	"strings"
)

func Detect_distro() string {
	cmd, err := exec.Command("uname", "-a").Output()
	if err != nil {
		log.Fatal(err)
	}
	if strings.ContainsAny(string(cmd), "ubuntu") {
		package_manager := "apt"
		return package_manager
	}
	return "pacman"
}
