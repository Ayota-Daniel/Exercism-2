package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	temp, exist := cb[file]
	if !exist {
		return 0
	}
	var count int
	for _, sq := range temp {
		if sq {
			count++
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if !(rank >= 1 && rank <= 8) {
		return 0
	}
	var count int
	for _, temp := range cb {
		if temp[rank-1] {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var count int
	for _, temp := range cb {
		for range temp {
			count++
		}
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var count int
	for _, temp := range cb {
		for _, sq := range temp {
			if sq {
				count++
			}
		}
	}
	return count
}
