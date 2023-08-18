package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

const default_url = "https://b23.tv/BV1uT4y1P7CX"

func fetch(url string, pattern string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64)")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return default_url
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return default_url
	}
	re := regexp.MustCompile(pattern)

	matched := re.FindAllStringSubmatch(string(body), -1)

	if len(matched) == 0 {
		fmt.Println("Target URL not found")
		return default_url
	}
	return matched[0][1] // first match, first group
}

func main() {
	url := "http://news.cyol.com/gb/channels/vrGlAKDl/index.html"
	res := fetch(url, "<li>.*<a href=\"(http.*?.html)\"")
	fmt.Println(res)
}
