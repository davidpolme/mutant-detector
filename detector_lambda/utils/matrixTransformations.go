package utils


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