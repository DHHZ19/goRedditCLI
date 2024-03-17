package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"

	"github.com/pkg/browser"
)

type RedditPostData struct {
	Data struct {
		After    string `json:"after"`
		Children []struct {
			Kind string `json:"kind"`
			Data struct {
				Title     string `json:"title"`
				Permalink string `json:"permalink"`
			} `json:"data,omitempty"`
		} `json:"children"`
	} `json:"data"`
}

func main() {
	resp, err := http.Get("https://www.reddit.com/.json")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)

	// Convert response body to RedditPostData struct
	var data RedditPostData
	err = json.Unmarshal([]byte(bodyBytes), &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	randomIndex := rand.Intn(len(data.Data.Children))

	printPtr := flag.Bool("print", false, "a bool")
	flag.Parse()

	// if -print flag is passed log the reddit post title and URL
	if *printPtr {
		fmt.Printf("The Reddit Post Title is:  %v\n The permalink is https://reddit.com%v\n", data.Data.Children[randomIndex].Data.Title, data.Data.Children[randomIndex].Data.Permalink)
	} else {
		randomPostPermaLink := data.Data.Children[randomIndex].Data.Permalink
		postURL := fmt.Sprintf("%s%s", "https://reddit.com", randomPostPermaLink)
		browser.OpenURL(postURL)
	}

}
