package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const baseURL = "https://api.github.com"


type GitHubClient struct {
	BaseURL			string
	Token			string
	Headers			http.Header
}


func newGitHubClient() *GitHubClient {
	token := os.Getenv("TOKEN")
	headers := http.Header{}

	if strings.TrimSpace(token) != "" {
		headers.Set("Authorization", "token "+strings.TrimSpace(token))
	}
	return &GitHubClient{
		BaseURL:  baseURL,
		Token:    token,
		Headers:  headers,
	}
}

func (c *GitHubClient) GetUserEvents(username string) {
	url := fmt.Sprintf("%s/users/%s/events", c.BaseURL, username)
	fmt.Println("GET", url, "with headers", c.Headers)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln("new request:", err)
	}

	req.Header.Set("User-Agent", "gh-activity/1.0")
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	for k, vals := range c.Headers {
		for _, v := range vals {
			req.Header.Add(k, v)
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("do request:", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("bad status: %s\n%s", resp.Status, string(body))
	}

	fmt.Println(string(body))
}

func main() {
	_ = godotenv.Load()
	client := newGitHubClient()
	client.GetUserEvents("leventeblanar")
}