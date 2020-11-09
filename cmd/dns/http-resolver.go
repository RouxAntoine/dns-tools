package main

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	check "antoine-roux.ml/projects/go/dns-tools/internal"
)

const (
	DnsResolverUrl = "https://cloudflare-dns.com/dns-query"
	DnsMethod      = "POST"
)

func main() {

	// Remove timestamp prefix from log out
	log.SetFlags(0)

	// Read wireformat from stdin
	data, err := ioutil.ReadFile(os.Stdin.Name())
	check.Check(err)

	// Create dns-message post request
	client := &http.Client{}
	req, err := http.NewRequest(DnsMethod, DnsResolverUrl, bytes.NewReader(data))
	req.Header.Add("Content-type", "application/dns-message")
	check.Check(err)

	// Do post request
	resp, err := client.Do(req)
	check.Check(err)
	defer resp.Body.Close()

	// Read http response and Print wireformat to stdout
	body, err := ioutil.ReadAll(resp.Body)
	check.Check(err)

	binary.Write(os.Stdout, binary.LittleEndian, body)
}
