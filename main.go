// Copyright 2023 Seppo Takalo
// SPDX-License-Identifier: MIT

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-ocf/go-coap"
)

var method = flag.String("x", "GET", "Request method: GET PUT POST")

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
	host, path := parseUri(flag.Arg(0))

	co, err := coap.Dial("udp", host)
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}

	switch *method {
	case "GET":
		resp, err := co.Get(path)
		if err != nil {
			log.Fatalf("Error sending request: %v", err)
		}
		fmt.Println(string(resp.Payload()))
	default:
		log.Fatal("Unknown method:", *method)
	}
}

func parseUri(uri string) (host string, path string) {
	u := strings.Split(uri, "/")
	host = u[0]
	path = u[1]
	if path == "" {
		path = "/"
	}
	if !strings.Contains(host, ":") {
		host = host + ":5683"
	}
	return
}
