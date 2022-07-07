package utils

import (
	"fmt"
	"strings"
)

//Verify the Adn sequence is valid
//Inputs: ss []string ,s string  - string DNA sequence and id string - id of the sequence
//Outputs error, bool - error if any, true if valid
func CheckADNStruct(ss []string, s string) (error) {
	//verify s is a matrix
	err := checkMatrix(ss)
	if err != nil {
		return err
	}

	err = validateChars(s)
	if err != nil {
		return err
	}
	return nil
}

//Verify the sequence only contains atcg characters
//Inputs: s string - string DNA sequence
//Outputs error, bool - error if any, true if valid
func validateChars(s string) (error) {
	const alpha = "atcg"
	for _, char := range s {
		if !strings.Contains(alpha, strings.ToLower(string(char))) {
			return fmt.Errorf("DNA sequence must only contain A,T,C or G characters")
		}
	}
	return nil
}

//Verify the sequence is a matrix
//Inputs: ss []string - matrix of DNA sequences
//Outputs error, bool - error if any, true if valid
func checkMatrix(ss []string) (error) {
	if len(ss) == 0 {
		return fmt.Errorf("length of the matrix must be greater than zero")
	}
	for _, row := range ss {
		if len(row) != 6 {
			return fmt.Errorf("length of the rows of the matrix must be equal to 6")
		}
	}
	return nil
}
