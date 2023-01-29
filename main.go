package main

import (
	"github.com/RAprogramm/NEVIRAIDE_install/operations"
)

var requirements = [11]string{}

func main() {
	println(operations.Detect_distro())
	operations.Backup_old()
	// package_manager := [2]string{"apt", "pacman"}
	// for _, manager := range package_manager {
	// 	cmd, err := exec.Command("which", manager).Output()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	// return string(cmd)
	// 	println(string(cmd))
	// }

}
