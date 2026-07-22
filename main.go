package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain,hasMX,hasSPF,SPFrecord,hasDMARC,DMARCrecord")

	for scanner.Scan() {
		checkdomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input: %v \n", err)
	}
}

func checkdomain(domain string) {

}
