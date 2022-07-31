package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/miekg/dns"
)

var (
	help    string
	version string
	port    string
)

func main() {
	flag.StringVar(&help, "help", "", "This describes the command and flags of dnsf")
	flag.StringVar(&version, "version", "", "The version of the program")
	flag.StringVar(&port, "port", "5053", "The port where the dns server will be served")

	flag.Parse()
	dns.HandleFunc("service.", handleDNSClient)

	srv := &dns.Server{
		Addr: fmt.Sprintf(":%s", port),
		Net:  "udp",
	}

	done := make(chan struct } , 1)
	log.Println("Currently running the DNS server")
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
		done <- true
	}()

	<-done
}
