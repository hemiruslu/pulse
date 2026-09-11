package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	URL      string
	Status   int
	Duration time.Duration
	Err      error
}

func checkSite(url string, ch chan<- Result) {
	start := time.Now()
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	duration := time.Since(start)

	if err != nil {
		ch <- Result{URL: url, Err: err, Duration: duration}
		return
	}
	defer resp.Body.Close()

	ch <- Result{URL: url, Status: resp.StatusCode, Duration: duration}
}

func main() {
	sites := []string{
		"https://google.com",
		"https://github.com",
		"https://mercari.com",
		"https://yahoo.co.jp",
		"https://amazon.co.jp",
		"https://invalid-host-test-domain.org",
	}

	resultsChannel := make(chan Result)
	fmt.Println("🚀 pulse: Running concurrent health checks...\n")
	startTime := time.Now()

	for _, site := range sites {
		go checkSite(site, resultsChannel)
	}

	for i := 0; i < len(sites); i++ {
		res := <-resultsChannel
		if res.Err != nil {
			fmt.Printf("❌ [DOWN]   %-38s -> Unreachable (%v)\n", res.URL, res.Duration.Round(time.Millisecond))
		} else {
			fmt.Printf("✅ [ONLINE] %-38s -> HTTP %d | Latency: %v\n", res.URL, res.Status, res.Duration.Round(time.Millisecond))
		}
	}

	fmt.Printf("\n✨ Completed %d checks in %v\n", len(sites), time.Since(startTime).Round(time.Millisecond))
}
