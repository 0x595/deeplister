// passive_enum.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func passiveEnumeration(domain string) []string {
	var subdomains []string

	// Get subdomains from crt.sh
	subdomains = append(subdomains, getSubdomainsFromCRT(domain)...)

	// Add more passive sources here if needed
	// subdomains = append(subdomains, getSubdomainsFromVirusTotal(domain)...)
	// subdomains = append(subdomains, getSubdomainsFromSecurityTrails(domain)...)
	// subdomains = append(subdomains, getSubdomainsFromAnubis(domain)...)

	return subdomains
}

func getSubdomainsFromCRT(domain string) []string {
	url := fmt.Sprintf("https://crt.sh/?q=%%25.%s&output=json", domain)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Warning: Error querying CRT.sh:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Warning: Error reading CRT.sh response:", err)
		return nil
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Warning: Error decoding CRT.sh JSON:", err)
		return nil
	}

	var subdomains []string
	for _, record := range result {
		if nameValue, exists := record["name_value"]; exists {
			// Split by newlines in case multiple domains are returned
			names := strings.Split(nameValue.(string), "\n")
			for _, name := range names {
				// Clean up the domain name
				cleanName := strings.TrimSpace(name)
				cleanName = strings.TrimPrefix(cleanName, "*.")
				if cleanName != "" && !strings.HasPrefix(cleanName, " ") {
					// Remove wildcard subdomains and invalid entries
					if !strings.Contains(cleanName, "*") {
						subdomains = append(subdomains, cleanName)
					}
				}
			}
		}
	}

	return subdomains
}

// You can add more passive enumeration functions here
/*
func getSubdomainsFromVirusTotal(domain string) []string {
	// Implementation for VirusTotal API
}

func getSubdomainsFromSecurityTrails(domain string) []string {
	// Implementation for SecurityTrails API
}

func getSubdomainsFromAnubis(domain string) []string {
	// Implementation for Anubis API
}
*/
