// package requirements define all dependecies for NEVIRAIDE
package requirements

import "github.com/RAprogramm/NEVIRAIDE_install/operations"

// Requirements function
func Requirements() map[string]operations.Application {
	var fd = "fd"

	if operations.Detect_manager() == "apt" {
		fd = "fdfind"
	}

	var requirements = map[string]operations.Application{
		"neovim":   {Name: "nvim", Check_with: "which"},
		"kitty":    {Name: "kitty", Check_with: "which"},
		fd:         {Name: fd, Check_with: "which"},
		"ripgrep":  {Name: "rg", Check_with: "which"},
		"npm":      {Name: "npm", Check_with: "which"},
		"git":      {Name: "git", Check_with: "which"},
		"lazygit":  {Name: "lazygit", Check_with: "which"},
		"unzip":    {Name: "unzip", Check_with: "which"},
		"delta":    {Name: "delta", Check_with: "which"},
		"nonicons": {Name: "nonicons", Check_with: "fc-list"},
		"neo":      {Name: "neo", Check_with: "which"},
		"neofetch": {Name: "neofetch", Check_with: "which"},
	}

	return requirements
}
