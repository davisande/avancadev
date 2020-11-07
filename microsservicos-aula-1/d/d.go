package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Result struct {
	Status string
}

func main() {
	http.HandleFunc("/", process)
	http.ListenAndServe(":9093", nil)
}

func process(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")

	sendEmail(coupon)

	result := Result{Status: "success"}

	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error processing json")
	}

	fmt.Fprintf(w, string(jsonData))
}

func sendEmail(coupon string) {
	fmt.Println("Sendind message")
	fmt.Printf("Message text: The coupon %v was used \n", coupon)
	fmt.Println("Message sent successfully")
}
