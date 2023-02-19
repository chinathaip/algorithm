package main

func matrixReshape(mat [][]int, r int, c int) [][]int {

	if !isPossible(mat, r, c) {
		return mat
	}

	result := [][]int{}
	//if possible

	m := len(mat)    //rows in mat
	n := len(mat[0]) //columns in mat

	row := []int{}
	counter := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//populate individual row's size to match with the reshaped columns(c)
			row = append(row, mat[i][j])
			counter++

			if counter == c {
				//once the number of columns match with the given c --> the row is ready, --> add to the resulting matrix
				result = append(result, row)
				counter = 0
				row = []int{}
			}
		}
	}

	return result
}

func isPossible(mat [][]int, r, c int) bool {
	require := r * c

	have := 0
	for _, arr := range mat {
		have += len(arr)
	}

	return have == require
}
