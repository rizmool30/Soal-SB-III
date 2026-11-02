package main

import (
	"encoding/json/v2"
	"fmt"
	"os"
)

type Review struct {
	Username       string `json:"username"`
	Score          int    `json:"score"`
	RestaurantName string `json:"restaurant_name"`
	Comment        string `json:"comment"`
}

func cleanReview(review Review) Review {
	if review.Username == "" {
		review.Username = "N/A"
	}

	if review.RestaurantName == "" {
		review.RestaurantName = "N/A"
	}

	if review.Comment == "" {
		review.Comment = "N/A"
	}

	return review
}

func main() {
	data, err := os.ReadFile("reviews.json")
	if err != nil {
		fmt.Println("Gagal membaca file:", err)
		return
	}

	var reviews []Review

	err = json.Unmarshal(data, &reviews)
	if err != nil {
		fmt.Println("Gagal membaca JSON:", err)
		return
	}
	for i, review := range reviews {
		reviews[i] = cleanReview(review)
	}
	hasil, err := json.Marshal(reviews)
	if err != nil {
		fmt.Println("Gagal mengubah data ke JSON:", err)
		return
	}
	err = os.WriteFile("reviews_cleaned.json", hasil, 0644)
	if err != nil {
		fmt.Println("Gagal menyimpan file:", err)
		return
	}

	fmt.Println("Data cleaning berhasil!")

	fmt.Println("Jumlah review:", len(reviews))
	fmt.Println("Data review pertama:", reviews[0])
}
