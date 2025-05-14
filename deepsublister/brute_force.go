// brute_force.go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

func bruteForceSubdomains(domain, wordlist string, threads int, verbose bool) []string {
	var subdomains []string
	var mu sync.Mutex

	file, err := os.Open(wordlist)
	if err != nil {
		fmt.Println("❌ Failed to open wordlist:", err)
		return subdomains
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	jobs := make(chan string, threads)
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < threads; i++ {
		wg.Add(1)
		// In brute_force.go, modify the worker goroutine:
		go func() {
			defer wg.Done()
			for sub := range jobs {
				host := sub + "." + domain
				if _, err := net.LookupHost(host); err == nil {
					mu.Lock()
					subdomains = append(subdomains, host)
					mu.Unlock()
					if verbose {
						fmt.Printf("[DNS] Found: %s\n", host)
					}
				}
			}
		}()
	}

	// Feed jobs to workers
	for scanner.Scan() {
		word := scanner.Text()
		if word != "" {
			jobs <- word
		}
	}
	close(jobs)
	wg.Wait()

	if err := scanner.Err(); err != nil {
		fmt.Println("❌ Error reading wordlist:", err)
	}

	return subdomains
}
