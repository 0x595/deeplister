// output.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func removeDuplicates(subdomains []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range subdomains {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func printResults(results []SubdomainResult) {
	if len(results) == 0 {
		fmt.Println("[!] No subdomains found.")
		return
	}

	// Sort by status code
	sort.Slice(results, func(i, j int) bool {
		return results[i].StatusCode < results[j].StatusCode
	})

	fmt.Println("\n[+] Results:")
	fmt.Println("======================================")
	for _, result := range results {
		if result.Error != nil {
			fmt.Printf("  - %s (Error: %v)\n", result.Subdomain, result.Error)
		} else {
			fmt.Printf("  - %s (HTTP %d)\n", result.Subdomain, result.StatusCode)
		}
	}
	fmt.Println("======================================")
	fmt.Printf("[+] Found %d unique subdomains\n", len(results))
}

func writeResultsToFile(filename string, results []SubdomainResult) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, result := range results {
		if result.Error != nil {
			file.WriteString(fmt.Sprintf("%s,error,%v\n", result.Subdomain, result.Error))
		} else {
			file.WriteString(fmt.Sprintf("%s,%d\n", result.Subdomain, result.StatusCode))
		}
	}

	return nil
}
