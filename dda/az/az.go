package az

// Language: go
// Path: dda\az\az.go
// Wrapper for interacting with the az cli

import (
	"encoding/json"
	"os/exec"
)

// Ensure that the az cli is installed
func IsInstalled() bool {
	cmd := exec.Command("az", "--version")
	_, err := cmd.Output()

	return err == nil
}

type User struct {
	EnvironmentName string `json:"environmentName"`
	HomeTenantId    string `json:"homeTenantId"`
	Id              string `json:"id"`
	IsDefault       bool   `json:"isDefault"`
	Name            string `json:"name"`
	State           string `json:"state"`
	TenantId        string `json:"tenantId"`
	User            struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"user"`
}

func IsLoggedIn() bool {
	cmd := exec.Command("az", "account", "show")
	stdout, err := cmd.Output()

	if err != nil {
		return false
	}

	// Read stdout as json
	var user User
	err = json.Unmarshal(stdout, &user)
	if err != nil {
		return false
	}

	return user.State != ""
}
