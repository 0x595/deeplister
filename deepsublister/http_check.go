// http_check.go
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type SubdomainResult struct {
	Subdomain  string
	StatusCode int
	Error      error
}

func checkStatusCodes(subdomains []string, verbose bool) []SubdomainResult {
	var results []SubdomainResult
	var wg sync.WaitGroup
	var mu sync.Mutex

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	jobs := make(chan string, len(subdomains))
	resultsChan := make(chan SubdomainResult, len(subdomains))

	// Start workers
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for subdomain := range jobs {
				result := checkSubdomain(client, subdomain, verbose)
				resultsChan <- result
			}
		}()
	}

	// Feed jobs
	for _, sub := range subdomains {
		jobs <- sub
	}
	close(jobs)

	// Collect results
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	for result := range resultsChan {
		mu.Lock()
		results = append(results, result)
		mu.Unlock()
	}

	return results
}

func checkSubdomain(client *http.Client, subdomain string, verbose bool) SubdomainResult {
	url := "http://" + subdomain
	resp, err := client.Get(url)
	if err != nil {
		if verbose {
			fmt.Printf("[HTTP Error] %s: %v\n", subdomain, err)
		}
		return SubdomainResult{
			Subdomain: subdomain,
			Error:     err,
		}
	}
	defer resp.Body.Close()

	if verbose {
		if resp != nil {
			fmt.Printf("[HTTP] Checking %s: %d\n", url, resp.StatusCode)
		} else {
			fmt.Printf("[HTTP] Checking %s: Error - %v\n", url, err)
		}
	}

	return SubdomainResult{
		Subdomain:  subdomain,
		StatusCode: resp.StatusCode,
	}
}
