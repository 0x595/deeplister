# SubEnum - Advanced Subdomain Enumeration Tool

![Go](https://img.shields.io/badge/Go-1.18+-blue.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

SubEnum is a powerful subdomain discovery tool that combines passive sources and brute-force techniques with HTTP status code verification.

## Features

- Passive subdomain enumeration (crt.sh)
- Brute-force subdomain discovery
- Concurrent DNS lookups
- HTTP status code checking
- Verbose output mode
- Results export to CSV
- Configurable worker threads

## Installation

### From Source
1. Ensure you have Go installed (v1.18+)
2. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/subenum.git
   cd subenum

Usage
Basic Command
./subenum -d example.com -w wordlist.txt
Full Command Options
./subenum -d example.com -w wordlist.txt \
    -o results.csv \
    -t 20 \
    -v \
    --passive \
    --brute
Options

Flag	Description	Example
-d, --domain	Target domain (required)	-d example.com
-w, --wordlist	Wordlist file for brute-force	-w wordlist.txt
-o, --output	Save results to file (CSV)	-o results.csv
-t, --threads	Number of threads (default: 10)	-t 20
-v, --verbose	Verbose output	-v
--passive	Passive enumeration only	--passive
--brute	Brute-force only	--brute


Examples

    Passive enumeration only:

./subenum -d example.com --passive -v

    Brute-force with custom threads:
./subenum -d example.com -w wordlist.txt -t 30

    Full scan with output file:
./subenum -d example.com -w wordlist.txt -o results.csv

    Verbose mode (show real-time results):
./subenum -d example.com -w wordlist.txt -v

Output Format

The tool displays results in the terminal and can export to CSV:

    [+] Results:
======================================
  - admin.example.com (HTTP 200)
  - beta.example.com (HTTP 302)
  - test.example.com (HTTP 404)
  - dev.example.com (Error: connection timed out)
======================================
[+] Found 4 unique subdomains

CSV Output Format
subdomain,status_code,error
admin.example.com,200,
beta.example.com,302,
test.example.com,404,
dev.example.com,0,connection timed out

Wordlist Recommendations

Use quality wordlists for better results:

    SecLists Subdomains

    Assetnote Wordlists
Troubleshooting

    Build errors: Run go mod tidy to fix dependencies

    Timeout issues: Increase timeout in http_check.go

    Missing subdomains: Try different wordlists or passive sources

Note: Use this tool only on domains you own or have permission to scan.


## Quick Command Reference

| Command | Description |
|---------|-------------|
| `./subenum -d DOMAIN -w WORDLIST` | Basic scan |
| `./subenum -d DOMAIN --passive` | Passive only |
| `./subenum -d DOMAIN -w WORDLIST --brute -t 20` | Brute-force with 20 threads |
| `./subenum -d DOMAIN -w WORDLIST -o out.csv` | Save results to CSV |
| `./subenum -d DOMAIN -w WORDLIST -v` | Verbose output |
