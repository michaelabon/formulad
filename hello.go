package main

import "fmt"

type Corner struct {
	StopsRequired int
	MaxPathLength int
	MinPaxLength  int
}

func makeCorners() []Corner {
	corners := make([]Corner, 0, 16)
	corners = append(corners, Corner{1, 4, 1})
	corners = append(corners, Corner{2, 10, 4})
	corners = append(corners, Corner{2, 10, 5})
	corners = append(corners, Corner{2, 11, 5})
	corners = append(corners, Corner{1, 5, 1})
	corners = append(corners, Corner{1, 8, 3})
	corners = append(corners, Corner{2, 11, 6})
	return corners
}

type Straight int

func makeStraights() []Straight {
	straights := []Straight{16, 34, 19, 10, 7, 5, 15}
	return straights
}

func rollGear(gear int) []int {
	switch gear {
	case 1:
		return []int{1, 2}
	case 2:
		return []int{2, 3, 3, 4, 4, 4}
	case 3:
		return []int{4, 5, 6, 6, 7, 7, 8, 8}
	case 4:
		return []int{7, 8, 9, 10, 11, 12}
	case 5:
		return []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	case 6:
		return []int{21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	}
	return nil
}

func main() {
	corners := makeCorners()
	fmt.Println(corners)
	straights := makeStraights()
	fmt.Println(straights)
	gears := [6]int{1, 2, 3, 4, 5, 6}
	for g := range gears {
		fmt.Println(rollGear(g))
	}
}
