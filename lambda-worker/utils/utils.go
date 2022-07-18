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

func CheckMainDiagonal(dna [][]string, count int) int {

	switch {
	case dna[0][0] == dna[3][3]:
		if dna[1][1] == dna[2][2] && dna[1][1] == dna[0][0] {
			count++
			break
		}
		fallthrough
	case true:
		if dna[3][3] == dna[4][4] {
			if dna[2][2] == dna[5][5] && dna[2][2] == dna[3][3] {
				count++
				break
			} else if dna[2][2] == dna[1][1] && dna[2][2] == dna[3][3] {
				count++
				break
			}
		}
	}

	return count
}

func CheckAdjacentDiagonals(dna [][]string, count int) int {
	adjX1 := 0
	adjX2 := 3
	centerX := 2
	centerY := 3
	c_A := make(chan int)
	c_B := make(chan int)
	//TODO: Put those two functions in goroutines
	go CheckAdjacentDiagonal_A(dna, adjX1, adjX2, centerX, centerY, c_A)
	go CheckAdjacentDiagonal_B(dna, adjX1, adjX2, centerX, centerY, c_B)
	count += <-c_A
	count += <-c_B
	
	if count > 1 {
		return count
	}
	return count
}

func CheckAdjacentDiagonal_B(dna [][]string, adjX1, adjX2, centerX, centerY int, c chan int) {
	count := 0
	for i := 1; i <= 2; i++ {
		//Check left diagonals
		switch {
		case dna[i][adjX1] == dna[i+adjX2][adjX2]:
			if dna[i+1][adjX1+1] == dna[i+2][adjX2-1] && dna[i][adjX1] == dna[i+1][adjX1+1] {
				count++
				i = 3
				break
			}
			fallthrough
		case true:
			if dna[centerY][centerX] == dna[centerY+1][centerX+1] {
				if dna[centerY-1][centerX-1] == dna[centerY+2][centerX+2] && dna[centerY-1][centerX-1] == dna[centerY][centerX] {
					count++
					i = 3
				}
			}
		}
	}
	c <- count
}

func CheckAdjacentDiagonal_A(dna [][]string, adjX1, adjX2, centerX, centerY int, c chan int) {
	count := 0
	for i := 1; i <= 2; i++ {
		switch {
		case dna[adjX1][i] == dna[adjX2][i+adjX2]:
			if dna[adjX1+1][i+1] == dna[adjX2-1][i+2] && dna[adjX2-1][i+2] == dna[adjX1][i] {
				count++
				i = 3
				break
			}
			fallthrough
		case true:
			if dna[centerX][centerY] == dna[centerX+1][centerY+1] {
				if dna[centerX-1][centerY-1] == dna[centerX+2][centerY+2] && dna[centerX-1][centerY-1] == dna[centerX+1][centerY+1] {
					count++
					i = 3
				}
			}
		}
	}
	c <- count
}
