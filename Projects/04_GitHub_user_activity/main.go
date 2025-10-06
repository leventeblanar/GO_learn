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


func newGitHubClient() *GitHubClient { 			// pointer a GitHubClient structra - nem másolatot adok át, hanem hivatkozást az eredetire (pointer a structra) 
	token := os.Getenv("TOKEN")					// Ez azért fontos mert másolat esetén a módosítás nem hat az eredetire
	headers := http.Header{}					// így viszont mindenhol ugyan azt az objektumot érjukel

	if token != "" {
		headers.Set("Authorization", "token "+token)
	}
	return &GitHubClient{
		BaseURL: baseURL,
		Token: token,
		Headers: headers,
	}
}

func (c *GitHubClient) GetUserEvents(username string) {
	url := fmt.Sprintf("%s/users/%s/events", c.BaseURL, username)
	fmt.Println("GET", url, "with headers", c.Headers)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln("new request:ű", err)
	}

	req.Header.Set("User-Agent", "gh-activity/1.0")
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	token := strings.TrimSpace(c.Token)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}


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

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("bad status: %s", resp.Status)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("read body:", err)
	}
	fmt.Println(string(b))
}

func main() {
	_ = godotenv.Load()
	client := newGitHubClient()
	client.GetUserEvents("leventeblanar")
}