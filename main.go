package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID             json.Number `json:"id"`
	Name           string      `json:"name"`
	LastUpdateDate time.Time   `json:"last_update_date"`
}

type Error struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

var users = []User{
	{
		ID:             "1",
		Name:           "太郎",
		LastUpdateDate: time.Now(),
	}, {
		ID:             "2",
		Name:           "花子",
		LastUpdateDate: time.Now(),
	}, {
		ID:             "3",
		Name:           "次郎",
		LastUpdateDate: time.Now(),
	},
}

var error_400 = Error{
	ID:      "ERROR-001",
	Message: "リクエストパラメータが不足しています",
}

func handler(w http.ResponseWriter, r *http.Request) {

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)

	service := r.URL.Query().Get("service")
	fmt.Println("service =>", service)
	if service == "" {
		if err := enc.Encode(&error_400); err != nil {
			log.Fatal(err)
		}
		fmt.Println(buf.String())
		http.Error(w, buf.String(), 400)
		return
	}

	if err := enc.Encode(&users); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())

	_, err := fmt.Fprint(w, buf.String())
	if err != nil {
		return
	}

}

func main() {

	http.HandleFunc("/users", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))

}
