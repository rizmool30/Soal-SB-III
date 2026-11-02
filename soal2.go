package main

import (
	"fmt"
	"strings"
)

func getTopKeyword(kata string) string {
	jumlah := make(map[string]int)

	katakata := strings.Fields(kata)

	for _, perkata := range katakata {
		jumlah[perkata]++

	}

	terbanyak := ""
	makskata := 0

	for perkata, hitung := range jumlah {
		if hitung > makskata {
			makskata = hitung
			terbanyak = perkata
		}
	}
	return terbanyak

}

func main() {
	fmt.Println(getTopKeyword("vibe coding is a coding while you are vibing along with coding"))
	fmt.Println(getTopKeyword("go makes programming fun fun"))
	fmt.Println(getTopKeyword("go go go"))
	fmt.Println(getTopKeyword("cats chase dogs and dogs chase dogs"))
	fmt.Println(getTopKeyword("the quick brown fox"))
	fmt.Println(getTopKeyword(""))

}
