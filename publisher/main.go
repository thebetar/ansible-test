package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func publishMessage(base_url string, subject string) {
	url := base_url + "/queue/push?id=" + subject

	data := map[string]string{"message": "Hello, World!"}
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Message published successfully!")

		fmt.Println("Response content:")
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	} else {
		fmt.Printf("Failed to publish message. Status code: %d\n", resp.StatusCode)
	}
}

func main() {
	base_url := os.Getenv("BASE_URL")

	if base_url == "" {
		base_url = "http://localhost:10526"
	}

	subject := os.Getenv("SUBJECT")

	if subject == "" {
		subject = "test-subject"
	}

	once := len(os.Args) > 1 && os.Args[1] == "--once"

	for {
		publishMessage(base_url, subject)
		if once {
			return
		}
		time.Sleep(30 * time.Second)
	}
}