package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func consumeMessage(base_url string, subject string) {
	url := base_url + "/queue/pop?id=" + subject

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Message consumed successfully!")
		fmt.Println("Message content:")

		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	} else {
		fmt.Printf("Failed to consume message. Status code: %d\n", resp.StatusCode)
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
		consumeMessage(base_url, subject)
		if once {
			return
		}
		time.Sleep(30 * time.Second)
	}
}