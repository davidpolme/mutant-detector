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
		dna = append(dna, StringToSlice(slice[i], 1))
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

func CheckMainDiagonal(dna [][]string, count int) int {
	if count > 1 {
		return count
	}
	if dna[0][0] == dna[3][3] {
		if dna[1][1] == dna[2][2] {
			count++
		}
	} else if dna[3][3] == dna[4][4] {
		if dna[2][2] == dna[5][5] {
			count++
		}
	}
	return count
}

func CheckAdjacentDiagonal(dna [][]string, count int) int {
	if count > 1 {
		return count
	}

	adjX1 := 0
	adjX2 := 3
	centerX := 2
	centerY := 3

	//TODO: Separate in two functions
	//TODO: Put those two functions in goroutines
	for i := 1; i <= 2; i++ {
		//Check left diagonals
		if dna[adjX1][i] == dna[adjX2][i+adjX2] {
			if dna[adjX1+1][i+1] == dna[adjX2-1][i+2] {
				count++
			}
		} else if i == 1 {
			//Its necesary to check both of possibilities in the nearest diagonal from  the main diagonal. When i = 1
			if dna[centerX][centerY] == dna[centerX+1][centerY+1] {
				if dna[centerX-1][centerY-1] == dna[centerX+2][centerY+2] {
					count++
				}
			}
		}
	}

	if count > 1 {
		return count
	}

	for i := 1; i <= 2; i++ {
		//Check left diagonals
		if dna[i][adjX1] == dna[i+adjX2][adjX2] {
			if dna[i+1][adjX1+1] == dna[i+2][adjX2-1] {
				count++
			}
		} else if i == 1 {
			//Its necesary to check both of possibilities in the nearest diagonal from  the main diagonal. When i = 1
			if dna[centerY][centerX] == dna[centerY+1][centerX+1] {
				if dna[centerY-1][centerX-1] == dna[centerY+2][centerX+2] {
					count++
				}
			}
		}
	}
	return count
}
