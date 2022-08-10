package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord")

	fmt.Println("Enter domain:")
	for scanner.Scan(){
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: Could not read from input, %v", err)
	}
}

func checkDomain(domain string){
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	MXRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Can't access MX records, %v", err)
		return
	}
	if len(MXRecords) > 0 {
		hasMX = true
	}
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Can't access SPF Records, %v", err)
		return
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Can't access dmarc Records, %v", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("MX: %v\nSPF: %v\nSPF Record: %v\nDMARC: %v\nDMARC Record: %v\n", hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}