package leetcode

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	helper(image, sr, sc, image[sr][sc], newColor)
	return image
}

func helper(image [][]int, sr int, sc int, oldColor int, newColor int) {
	image[sr][sc] = newColor
	if sr-1 >= 0 && image[sr-1][sc] == oldColor {
		helper(image, sr-1, sc, oldColor, newColor)
	}
	if sr+1 < len(image) && image[sr+1][sc] == oldColor {
		helper(image, sr+1, sc, oldColor, newColor)
	}
	if sc-1 >= 0 && image[sr][sc-1] == oldColor {
		helper(image, sr, sc-1, oldColor, newColor)
	}
	if sc+1 < len(image[0]) && image[sr][sc+1] == oldColor {
		helper(image, sr, sc+1, oldColor, newColor)
	}
}
