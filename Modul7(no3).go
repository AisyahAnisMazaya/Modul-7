package main

import "fmt"

func main() {
	var club1, club2 string
	var skor1, skor2 int
	var clubMenang []string
	var hasil []string

	fmt.Print("klub A: ")
	fmt.Scan(&club1)
	fmt.Print("klub B: ")
	fmt.Scan(&club2)

	var pertandingan int = 1
	for {
		fmt.Printf("Pertandingan %d: ", pertandingan)
		fmt.Scan(&skor1)
		fmt.Scan(&skor2)

		if skor1 < 0 || skor2 < 0 {
			break
		}

		if skor1 > skor2 {
			clubMenang = append(clubMenang, club1)
			hasil = append(hasil, club1)
		} else if skor2 > skor1 {
			clubMenang = append(clubMenang, club2)
			hasil = append(hasil, club2)
		} else {
			hasil = append(hasil, "Draw")
		}

		pertandingan++
	}

	for i, result := range hasil {
		fmt.Printf("Hasil %d: %s\n", i+1, result)
	}
	fmt.Println("Pertandingan selesai")

}
