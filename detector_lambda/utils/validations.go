package utils

func IsMutant(dnaRequest []string) bool {
	dnaMatrix := SliceToMatrix(dnaRequest)
	dnaResponse := searchPatterns(dnaMatrix)
	return dnaResponse
}

func searchPatterns(dna [][]string) bool {
	count := 0
	count = checkHorizontal(dna, count)
	//log.Println("[count horizontal]", count)
	if count > 1 {
		return true
	}
	count = checkVertical(dna, count)
	//log.Println("[count vertical]", count)

	if count > 1 {
		return true
	}
	count = checkDiagonalNegative(dna, count)

	if count > 1 {
		return true
	}

	count = checkDiagonalPositive(dna, count)
	return count > 1
}

func checkHorizontal(dna [][]string, count int) int {

	for i := 0; i < len(dna); i++ {
		//Si en esta secuencia count es mayor que 1 se retorna el valor de count
		if count > 1 {
			break
		}
		//search hints for possible patterns
		if dna[i][0] == dna[i][3] {
			if dna[i][0] == dna[i][1] && dna[i][1] == dna[i][2] {
				count++
				continue
			}
		}
		if dna[i][3] == dna[i][4] {
			if dna[i][2] == dna[i][3] && dna[i][1] == dna[i][3] {
				count++
				continue
			}
		}
	}
	return count
}

func checkVertical(dna [][]string, count int) int {
	dnaMatrix := TransposeMatrix(dna)
	count = checkHorizontal(dnaMatrix, count)
	//log.Println("[count Vertical]", count)
	return count
}

func checkDiagonalNegative(dna [][]string, count int) int {
	//1 step: Check main diagonal
	count = CheckMainDiagonal(dna, count)

	if count > 1 {
		return count
	}

	//2 step, check adjacent diagonals
	count = CheckAdjacentDiagonals(dna, count)
	return count
}

func checkDiagonalPositive(dna [][]string, count int) int {
	dnaMatrix := ReverseMatrix(dna)
	count = checkDiagonalNegative(dnaMatrix, count)
	//log.Println("[count d-positive]: ", count)
	return count
}
