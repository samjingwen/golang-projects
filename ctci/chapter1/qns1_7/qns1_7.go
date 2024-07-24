package qns1_7

type Image [][]int32

// 1, 2, 3     7, 4, 1
// 4, 5, 6 ->  8, 5, 2
// 7, 8, 9     9, 6, 3
// 0,0 <- 2,0 | 0,1 <- 1,0
// 2,0 <- 2,2 | 1,0 <- 2,1
// 2,2 <- 0,2 | 2,1 <- 1,2
// 0,2 <- 0,0 | 1,2 <- 0,1

func RotateMatrix(image Image) Image {
	n := len(image)
	var x, y int
	if n%2 == 0 {
		x = n / 2
		y = n / 2
	} else {
		x = n/2 + 1
		y = n / 2
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			temp := image[i][j]
			image[i][j] = image[n-1-j][i]
			image[n-1-j][i] = image[n-1-i][n-1-j]
			image[n-1-i][n-1-j] = image[j][n-1-i]
			image[j][n-1-i] = temp
		}
	}
	return image
}

func (image Image) equals(other Image) bool {
	for i, x := range image {
		for j, y := range x {
			if other[i][j] != y {
				return false
			}
		}
	}
	return true
}
