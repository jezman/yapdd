package utils

import "strings"

// IsAccount checks
func IsAccount(flag string) bool {
	if strings.Contains(flag, "@") {
		return true
	}
	return false
}
