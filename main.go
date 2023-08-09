package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/crunchy-labs/crunchyroll-go/v3"
)

func main() {
	username := os.Getenv("crunchyUsername")
	password := os.Getenv("crunchyUsername")
	client := http.DefaultClient

	_, err := crunchyroll.LoginWithCredentials(username, password, crunchyroll.US, client)
	if err != nil {
		fmt.Println("Unsuccessfully logged in", err.Error())
	}
}
