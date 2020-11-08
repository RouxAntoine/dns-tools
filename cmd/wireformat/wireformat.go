package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"

	"github.com/miekg/dns"
)

func main() {
	file := flag.String("f", "", "file path to read hexa decimal wireformat data")
	flag.Parse()

	var body []byte
	var err error
	if *file != "" {
		if *file == "-" {
			*file = "/dev/stdin"
		}
		body, err = ioutil.ReadFile(*file)
		if err != nil {
			log.Fatalln(err)
		}

		msg := &dns.Msg{}
		err = msg.Unpack(body)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(msg)
	} else {
		log.Println(errors.New("Too few argument pass"))
		flag.Usage()
	}
}
