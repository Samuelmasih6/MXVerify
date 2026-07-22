# MXVerify

A lightweight command-line email domain verification tool written in Go.

MXVerify validates whether a domain is properly configured to send and receive emails by checking essential DNS records used in modern email infrastructure.

## Features

* Verify MX (Mail Exchange) records
* Verify SPF (Sender Policy Framework) records
* Verify DMARC (Domain-based Message Authentication, Reporting and Conformance) records
* Process single or multiple domains from standard input
* Export results in CSV format
* Fast and dependency-free implementation using Go's standard library

## Why MXVerify?

Misconfigured email domains often lead to:

* Emails being rejected
* Messages landing in spam folders
* Email spoofing vulnerabilities
* Reduced deliverability

MXVerify helps identify common email configuration issues by inspecting DNS records associated with a domain.

## How It Works

For each domain provided, MXVerify performs:

### MX Record Lookup

Checks whether the domain can receive emails.

Example:

```text
gmail.com -> MX records found
```

### SPF Record Lookup

Checks for a TXT record beginning with:

```text
v=spf1
```

SPF specifies which mail servers are authorized to send emails on behalf of the domain.

### DMARC Record Lookup

Checks for a TXT record under:

```text
_dmarc.<domain>
```

A valid DMARC record begins with:

```text
v=DMARC1
```

DMARC helps protect domains from spoofing and phishing attacks.

## Installation

### Prerequisites

* Go 1.20+

### Clone the Repository

```bash
git clone https://github.com/<your-username>/mxverify.git
cd mxverify
```

### Run

```bash
go run main.go
```

## Usage

### Interactive Mode

Start the application:

```bash
go run main.go
```

Enter domains one per line:

```text
gmail.com
google.com
github.com
```

Terminate input:

```bash
Ctrl + D
```

### File Input

Create a file:

```text
gmail.com
google.com
openai.com
github.com
```

Run:

```bash
go run main.go < domains.txt
```

### Export Results

```bash
go run main.go < domains.txt > results.csv
```

## Example Output

```csv
domain,hasMX,hasSPF,SPFrecord,hasDMARC,DMARCrecord
gmail.com,true,true,"v=spf1 redirect=_spf.google.com",true,"v=DMARC1; p=none"
github.com,true,true,"v=spf1 include:_spf.github.com ~all",true,"v=DMARC1; p=reject"
```

## Sample Checks

| Check      | Description                                   |
| ---------- | --------------------------------------------- |
| MX         | Determines whether a domain can receive email |
| SPF        | Validates sender authorization policy         |
| DMARC      | Validates anti-spoofing policy                |
| TXT Lookup | Retrieves domain TXT records                  |

## Project Structure

```text
.
├── main.go
├── domains.txt
├── results.csv
└── README.md
```

## Current Limitations

MXVerify currently validates domain-level email configuration only.

Supported:

* MX validation
* SPF validation
* DMARC validation

Not yet supported:

* Email syntax validation
* SMTP connectivity testing
* Catch-all detection
* Disposable email detection
* Individual mailbox verification
* DKIM validation

## Future Improvements

* DKIM record verification
* SMTP handshake testing
* Catch-all domain detection
* Concurrent DNS lookups
* JSON output support
* REST API version
* Bulk verification support
* Web dashboard

## Example Domains

```text
gmail.com
google.com
github.com
openai.com
amazon.com
microsoft.com
```

## Tech Stack

* Go
* DNS Lookups (`net` package)
* Standard Library

## Learning Outcomes

This project demonstrates:

* DNS record resolution in Go
* Network programming fundamentals
* CSV generation
* Command-line application development
* Email infrastructure concepts
* Error handling and input processing
