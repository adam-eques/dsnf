package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/acentior/dnsf/cache"
	"github.com/acentior/dnsf/handlers"
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

	c, err := cache.NewCache()
	if err != nil {
		panic(fmt.Errorf("Unable to intialize cache %v", err))
	}
	h := handlers.QueryHandler{
		Cache: c,
	}

	dns.HandleFunc("service.", h.HandleDNSClient)

	srv := &dns.Server{
		Addr: fmt.Sprintf(":%s", port),
		Net:  "udp",
	}

	done := make(chan struct{}, 1)
	log.Println("Currently running the DNS server")
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
		done <- struct{}{}
	}()

	<-done
}
