package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var (
		domain      string
		wordlist    string
		output      string
		threads     int
		passiveOnly bool
		bruteOnly   bool
		verbose     bool
	)

	var rootCmd = &cobra.Command{
		Use:   "subenum",
		Short: "Advanced Subdomain Enumeration Tool",
		Long: `An advanced subdomain enumeration tool that combines passive sources
and brute-force techniques with status code checking.`,
		Run: func(cmd *cobra.Command, args []string) {
			if domain == "" {
				log.Fatal("❌ Domain is required")
			}

			fmt.Printf("[*] Starting enumeration on: %s\n", domain)

			var allSubdomains []string

			// 1. Passive Enumeration
			if !bruteOnly {
				fmt.Println("[*] Running passive enumeration...")
				passiveSubs := passiveEnumeration(domain)
				allSubdomains = append(allSubdomains, passiveSubs...)
			}

			// 2. Brute-force (DNS)
			if !passiveOnly && wordlist != "" {
				fmt.Printf("[*] Running brute-force with %d threads...\n", threads)
				bruteSubs := bruteForceSubdomains(domain, wordlist, threads, verbose)
				allSubdomains = append(allSubdomains, bruteSubs...)
			}

			if len(allSubdomains) == 0 {
				fmt.Println("[!] No subdomains found")
				return
			}

			// 3. Process results
			uniqueSubs := removeDuplicates(allSubdomains)
			fmt.Printf("[+] Found %d unique subdomains\n", len(uniqueSubs))

			// 4. HTTP Status Check
			fmt.Println("[*] Checking HTTP status codes...")
			results := checkStatusCodes(uniqueSubs, verbose)

			// 5. Display final results
			printResults(results)

			// 6. Save to file if requested
			if output != "" {
				if err := writeResultsToFile(output, results); err != nil {
					log.Fatalf("Failed to save results: %v", err)
				}
				fmt.Printf("[+] Results saved to %s\n", output)
			}
		},
	}

	// Flags
	rootCmd.Flags().StringVarP(&domain, "domain", "d", "", "Target domain (required)")
	rootCmd.Flags().StringVarP(&wordlist, "wordlist", "w", "", "Path to wordlist file")
	rootCmd.Flags().StringVarP(&output, "output", "o", "", "File to save results")
	rootCmd.Flags().IntVarP(&threads, "threads", "t", 10, "Number of threads for brute-force")
	rootCmd.Flags().BoolVar(&passiveOnly, "passive", false, "Only perform passive enumeration")
	rootCmd.Flags().BoolVar(&bruteOnly, "brute", false, "Only perform brute-force enumeration")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}
