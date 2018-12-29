package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for y := range image {
		image[y] = make([]uint8, dx)
		for x := range image[y] {
			image[y][x] = uint8(x * y)
		}
	}

	return image
}

func main() {
	pic.Show(Pic)
}
