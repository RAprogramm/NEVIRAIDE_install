// This is auto install script of NEVIRAIDE
package main

import (
	"github.com/RAprogramm/NEVIRAIDE_install/operations"
	"github.com/RAprogramm/NEVIRAIDE_install/requirements"
)

func main() {
	operations.Check_installed(requirements.Requirements())

}
