// Copyright 2023 Seppo Takalo
// SPDX-License-Identifier: MIT

package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/go-ocf/go-coap"
	"github.com/pion/dtls/v2"
)

var method = flag.String("x", "GET", "Request method: GET PUT POST")
var unsec = flag.Bool("i", false, "Insecure: skip most security checks (CA chain, hostname)")

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] <URL>:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(2)
	}
	u, err := url.Parse(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	// Add default port to host
	if u.Port() == "" {
		switch u.Scheme {
		case "coap":
			u.Host += ":5683"
		case "coaps":
			u.Host += ":5684"
		}
	}

	dtlsconf := dtls.Config{
		InsecureSkipVerify: *unsec,
		InsecureHashes:     *unsec,
	}
	var co *coap.ClientConn
	switch u.Scheme {
	case "coap":
		co, err = coap.Dial("udp", u.Host)
	case "coaps":
		co, err = coap.DialDTLS("udp-dtls", u.Host, &dtlsconf)
	}
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}

	switch *method {
	case "GET":
		resp, err := co.Get(u.Path)
		if err != nil {
			log.Fatalf("Error sending request: %v", err)
		}
		fmt.Println(string(resp.Payload()))
	default:
		log.Fatal("Unknown method:", *method)
	}
}
