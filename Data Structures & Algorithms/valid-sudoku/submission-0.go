func isValidSudoku(board [][]byte) bool {

	for row:=0;row<9;row++{
		seen := make(map[byte]bool)
		for j:=0;j<9;j++{
			if board[row][j]=='.' {
				continue
			}
			if seen[board[row][j]] {
				return false
			} else {
				seen[board[row][j]] = true
			}
		}
	}	

	for col:=0;col<9;col++{
		seen := make(map[byte]bool)
		for j:=0;j<9;j++{
			if board[j][col]=='.' {
				continue
			}
			if seen[board[j][col]] {
				return false
			} else {
				seen[board[j][col]] = true
			}
		}
	}	

	col := []int{0, -1, -1, 0, 1, 1, 1, 0, -1}
	row := []int{0, 0, -1, -1, -1, 0, 1, 1, 1}
	for i:=0;i<9;i++{
		for j:=0;j<9;j++{
			if is3x3Center(i,j) {
				seen := make(map[byte]bool)

				for t:=0;t<9;t++{
					if seen[board[i+row[t]][j+col[t]]]{
						return false
					} else if board[i+row[t]][j+col[t]]!='.'{
						seen[board[i+row[t]][j+col[t]]]=true
					}
				}
			}
		}
	}

	return true;
}

func is3x3Center(i, j int) bool{
	iValid := false;
	jValid := false;
	if i==1 || i==4 || i==7 {
		iValid = true;
	}
	if j==1 || j==4 || j==7 {
		jValid = true;
	}
	return iValid && jValid
}
