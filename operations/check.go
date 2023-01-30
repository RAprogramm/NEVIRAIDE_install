// package operations consist of functions wich make installation process easy
package operations

import (
	"fmt"
	"os/exec"
)

type Application struct {
	Name       string
	Check_with string
}

// Check_installed function checking all dependecies on your computer
func Check_installed(list_of_apps map[string]Application) {
	for idx, app := range list_of_apps {
		_, err := exec.Command(app.Check_with, app.Name).Output()
		if err != nil {
			fmt.Println(idx, "not installed")
			fmt.Println("Now will be installed", idx)
			_, err1 := exec.Command("sudo", "apt", "install", "-y", app.Name).Output()
			if err1 != nil {
				fmt.Println("Problem with installation", idx)
			}
		}
		fmt.Println(idx, "installed")
	}
}
