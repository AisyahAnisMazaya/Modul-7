package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	tengah titik
	radius int
}

func main() {
	var c1, c2 lingkaran
	var p titik

	fmt.Scan(&c1.tengah.x, &c1.tengah.y, &c1.radius)
	fmt.Scan(&c2.tengah.x, &c2.tengah.y, &c2.radius)
	fmt.Scan(&p.x, &p.y)

	if dalamLingkaran(c1, c2, p) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dL12(c1, p) {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dL12(c2, p) {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

func dalamLingkaran(c1, c2 lingkaran, p titik) bool {
	return dL12(c1, p) && dL12(c2, p)
}

func dL12(c lingkaran, p titik) bool {
	return math.Pow(float64(p.x-c.tengah.x), 2)+math.Pow(float64(p.y-c.tengah.y), 2) <= math.Pow(float64(c.radius), 2)
}
