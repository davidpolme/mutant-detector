package utils

import (
	"strings"
)

//IsMatrix checks if the matrix is valid
func IsMatrix(ss []string) bool {
	// check if the matrix is size of 6x6
	if len(ss) != 6 {
		return false
	}
	for _, s := range ss {
		if len(s) != 6 {
			return false
		}
	}
	return true
}

func ContainsValidChars(ss []string) bool {
	s := strings.Join(ss, "")
	const alpha = "atcg"
	for _, char := range s {
		if !strings.Contains(alpha, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}
