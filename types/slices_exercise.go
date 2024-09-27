package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		slice[y] = make([]uint8, dx)

		for x := 0; x < dx; x++ {
			if float64(((x- 128)^2 + (y-128)^2)) < float64(128^2) {
				slice[y][x] = uint8(0)
			} else {
				slice[y][x] = uint8(255)
			}
			// slice[i][j] = uint8(i ^ j)
		}
	}
	return slice
}

func sliceExcercise() {
	pic.Show(Pic)
}
