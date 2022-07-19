package utils

//transformIntoMatrix takes a slice of strings and returns a slice of slices of strings
//Inputs: slice of strings
//Outputs: slice of slices of strings
func SliceToMatrix(slice []string) [][]string {
	var dna [][]string
	for i := 0; i < len(slice); i++ {
		dna = append(dna, StringToSlice(slice[i], 1))
	}
	return dna
}

//stringToGrid takes a string and returns a slice of strings
//Inputs:
//     s : string to convert
//	   col : number of columns in the strings
// Outputs: slice of strings
//Outputs: slice of strings
func StringToSlice(s string, col int) []string {
	var newS []string
	for i := 0; i < len(s); i += col {
		newS = append(newS, s[i:i+col])
	}
	return newS
}
