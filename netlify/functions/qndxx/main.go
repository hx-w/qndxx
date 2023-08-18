package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const DEFAULT_URL = "https://b23.tv/BV1uT4y1P7CX"
const URL = "http://news.cyol.com/gb/channels/vrGlAKDl/index.html"
const PATTERN = "<li>.*<a href=\"(http.*?.html)\""

func Resp(StatusCode int, Body string) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode:      int(StatusCode),
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            Body,
		IsBase64Encoded: false,
	}, nil
}

func fetch(url string, pattern string) (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64)")
	resp, err := client.Do(req)
	if err != nil {
		return DEFAULT_URL, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return DEFAULT_URL, err
	}
	re := regexp.MustCompile(pattern)

	matched := re.FindAllStringSubmatch(string(body), -1)

	if len(matched) == 0 {
		fmt.Println("Target URL not found")
		return DEFAULT_URL, nil
	}
	target := matched[0][1]
	target = strings.Replace(target, "m.html", "images/end.jpg", 1)
	target = strings.Replace(target, "index.html", "images/end.jpg", 1)

	return target, nil
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	result, err := fetch(URL, PATTERN)
	if err != nil {
		return Resp(403, "{\"result\": \""+result+"\"}")
	}
	return Resp(200, "{\"result\": \""+result+"\"}")
}

func main() {
	lambda.Start(handler)
}
