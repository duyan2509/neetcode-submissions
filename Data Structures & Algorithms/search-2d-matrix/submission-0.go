func searchMatrix(matrix [][]int, target int) bool {
	m,n:= len(matrix),len(matrix[0])
	left := 0
	right := m*n-1

	for left <= right {
		mid := left + (right -left)/ 2
		i:= mid/n
		j:= mid%n
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			left=mid+1
		} else {
			right=mid-1
		}
	}

	return false
}
