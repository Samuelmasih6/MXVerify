package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain,hasMX,hasSPF,SPFrecord,hasDMARC,DMARCrecord")

	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		if domain == "" {
			continue
		}
		checkdomain(domain)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input: %v", err)
	}
}

func checkdomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var SPFrecord, DMARCrecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			SPFrecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			DMARCrecord = record
			break
		}
	}
	fmt.Printf("%v,%v,%v,%q,%v,%q\n", domain, hasMX, hasSPF, SPFrecord, hasDMARC, DMARCrecord)

}
