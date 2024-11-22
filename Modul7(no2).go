package main

import (
	"fmt"
)

func main() {
	var N, x int
	fmt.Scan(&N)

	var arr []int
	for i := 0; i < N; i++ {
		var bilangan int
		fmt.Scan(&bilangan)
		arr = append(arr, bilangan)
	}

	fmt.Println("a. Menampilkan keseluruhan isi dari array.")
	fmt.Println(arr)

	fmt.Println("b. Menampilkan elemen-elemen array dengan indeks ganjil saja.")
	for i := 0; i < N; i++ {
		if i%2 != 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	fmt.Println("c. Menampilkan elemen-elemen array dengan indeks genap saja.")
	for i := 0; i < N; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	fmt.Println("d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x.")
	fmt.Scan(&x)
	for i := 0; i < N; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	fmt.Println("e. Menghapus elemen array pada indeks tertentu.")
	var index int
	fmt.Scan(&index)
	if index < len(arr) {
		arr = append(arr[:index], arr[index+1:]...)
		fmt.Println("Array setelah penghapusan:")
		fmt.Println(arr)
	} else {
		fmt.Println("Indeks tidak valid.")
	}

	fmt.Println("f. Menampilkan rata-rata dari bilangan yang ada di dalam array.")
	var sum float64
	for i := 0; i < len(arr); i++ {
		sum += float64(arr[i])
	}
	fmt.Printf("%.2f\n", sum/float64(len(arr)))
	fmt.Println("g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array.")
	var mean float64
	for i := 0; i < len(arr); i++ {
		mean += float64(arr[i])
	}
	mean /= float64(len(arr))

	var sd float64
	for i := 0; i < len(arr); i++ {
		sd += (float64(arr[i]) - mean) * (float64(arr[i]) - mean)
	}
	sd = sqrt(sd / float64(len(arr)))
	fmt.Printf("%.2f\n", sd)

	fmt.Println("h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi.")
	var bilangan int
	fmt.Scan(&bilangan)
	var freq int
	for i := 0; i < len(arr); i++ {
		if arr[i] == bilangan {
			freq++
		}
	}
	fmt.Println(freq)
}

func sqrt(x float64) float64 {
	var z float64
	z = x / 2
	for i := 0; i < 1000; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}
