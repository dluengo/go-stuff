package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var s [][]uint8

	s = make([][]uint8, dy)
	for i := range s {
		s[i] = make([]uint8, dx)
	}

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			s[i][j] = uint8(i^j)
		}
	}

	return s
}

func main() {
	pic.Show(Pic)
}

