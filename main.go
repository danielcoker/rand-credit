package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

const creditScoreMin = 500
const creditScoreMax = 900

type creditRating struct {
	CreditRating int `json:"creditRating"`
}

func getCreditScore(w http.ResponseWriter, r *http.Request) {
	creditRating := creditRating{
		CreditRating: rand.Intn(creditScoreMax-creditScoreMin) + creditScoreMin,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(creditRating)
}

func handleRequests() {
	http.HandleFunc("/creditScore", getCreditScore)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
