package main

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type Client struct {
	client      *http.Client
	rateLimiter *rate.Limiter
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	ctx := context.Background()
	err := c.rateLimiter.Wait(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func NewClient(rl *rate.Limiter) *Client {
	c := &Client{
		client:      http.DefaultClient,
		rateLimiter: rl,
	}
	return c
}

func main() {
	rl := rate.NewLimiter(rate.Every(10*time.Second), 50)
	client := NewClient(rl)
	// здесь, например, API биржевых котировок
	URL := "https://iss.moex.com/iss/statistics/engines/futures/markets/indicativerates/securities.xml"
	req, _ := http.NewRequest("GET", URL, nil)
	for {
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// логика обработки результата запроса
		defer resp.Body.Close()
		scanner := bufio.NewScanner(resp.Body)
		for i := 0; scanner.Scan() && i < 20; i++ {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
}
