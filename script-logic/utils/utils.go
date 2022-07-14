package utils

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

//transformIntoMatrix takes a slice of strings and returns a slice of slices of strings
//Inputs: slice of strings
//Outputs: slice of slices of strings
func SliceToMatrix(slice []string) [][]string {
	var dna [][]string
	for i := 0; i < len(slice); i++ {
		dna = append(dna, StringToSlice(slice[i],1))
	}
	return dna
}

//transformIntoMatrix Transpose matrix
//Inputs: slice of strings - Original matrix
//Outputs: slice of strings - Transposed matrix
func TransposeMatrix(slice [][]string) [][]string {
	var newSlice [][]string
	for i := 0; i < len(slice[0]); i++ {
		var newRow []string
		for j := 0; j < len(slice); j++ {
			newRow = append(newRow, slice[j][i])
		}
		newSlice = append(newSlice, newRow)
	}
	return newSlice
}

//reverseMatrix reverses the matrix
//Inputs: slice of strings - Original matrix
//Outputs: slice of strings - Reversed matrix
func ReverseMatrix(slice [][]string) [][]string {
	var newSlice [][]string
	for i := range slice {
		newSlice = append(newSlice, slice[len(slice)-1-i])
	}
	return newSlice
}



/*
TODO: Delete this function
func CheckIf4ConsecutiveChar(dna [][]string) bool {
	count := 0
	for i := 0; i < len(dna); i++ {
		for j := 0; j < len(dna[i]); j++ {
			if dna[i][j] == dna[i][j+1] {
				count++
			}
		}
		if count == 3 {
			return true
		}
		count = 0
	}
	return false
}
*/