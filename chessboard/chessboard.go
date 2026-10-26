package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	howMany := 0
	for _, content := range cb[file] {
		if content == true {
			howMany++
		}
	}
	return howMany
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	howMany := 0
	where := rank - 1
	for _, content := range cb {
		if where < len(content) && content[where] {
			howMany++
		}
	}
	return howMany
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	howMany := 0
	for _, content := range cb {
		howMany += len(content)
	}
	return howMany
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	howMany := 0
	for _, content := range cb {
		for _, index := range content {
			if index == true {
				howMany++
			}
		}
	}
	return howMany
}
