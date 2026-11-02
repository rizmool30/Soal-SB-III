package main

import "fmt"

func checkAnagram(word1, word2 string) bool {

	if len(word1) != len(word2) {
		return false
	}

	jumlah := make(map[rune]int)

	for _, huruf := range word1 {
		jumlah[huruf]++
	}

	for _, huruf := range word2 {
		jumlah[huruf]--
	}

	for _, huruf := range jumlah {
		if huruf != 0 {
			return false
		}
	}
	return true

}
func main() {
	fmt.Println(checkAnagram("lemon", "melon"))
	fmt.Println(checkAnagram("listen", "silent"))
	fmt.Println(checkAnagram("triangle", "integral"))
	fmt.Println(checkAnagram("hello", "world"))
	fmt.Println(checkAnagram("", "   "))

}
