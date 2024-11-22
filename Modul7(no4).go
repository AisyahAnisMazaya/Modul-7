package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	fmt.Print("Teks: ")
	*n = 0
	var input rune

	for {
		fmt.Scanf("%c", &input)

		if input == '.' || *n >= NMAX {
			break
		}

		if input != ' ' {
			t[*n] = input
			*n++
		}
	}

	if *n == NMAX {
		fmt.Println("Maksimal 127 karakter telah tercapai.")
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]), " ")
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	var temp tabel
	copy(temp[:], t[:])
	balikanArray(&temp, n)
	for i := 0; i < n; i++ {
		if t[i] != temp[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int
	isiArray(&tab, &m)
	balikanArray(&tab, m)
	fmt.Print("Reverse Teks: ")
	cetakArray(tab, m)
	if palindrom(tab, m) {
		fmt.Println("Palindrom : true")
	} else {
		fmt.Println("Palindrom : flase")
	}
}
