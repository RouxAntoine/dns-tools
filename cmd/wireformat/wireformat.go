package main

import (
	"encoding/binary"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"

	check "antoine-roux.ml/projects/go/dns-tools/internal"

	"github.com/miekg/dns"
)

func main() {
	file := flag.String("f", "", "file path to read hexa decimal wireformat data")
	domainToResolve := flag.String("s", "-", "Server adress to resolve")
	flag.Parse()

	// remove timestamp prefix from log out
	log.SetFlags(0)

	var body []byte
	var err error
	if *file != "" {
		if *file == "-" {
			*file = os.Stdin.Name()
		}
		body, err = ioutil.ReadFile(*file)
		check.Check(err)

		msg := &dns.Msg{}
		err = msg.Unpack(body)
		check.Check(err)

		log.Println(msg)
	} else if *domainToResolve != "" {
		if *domainToResolve == "-" {
			readDomainToResolve, err := ioutil.ReadFile(os.Stdin.Name())
			check.Check(err)
			*domainToResolve = strings.TrimSuffix(string(readDomainToResolve), "\n")
		}
		msg := &dns.Msg{
			Question: []dns.Question{
				dns.Question{
					dns.Fqdn(*domainToResolve),
					dns.TypeA,
					dns.ClassINET,
				},
			},
		}
		b, err := msg.Pack()
		check.Check(err)
		binary.Write(os.Stdout, binary.LittleEndian, b)
	} else {
		log.Println(errors.New("Too few argument pass"))
		flag.Usage()
	}
}
