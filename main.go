package main

import (
	"fmt"
	"log"

	"github.com/miekg/dns"
)

var records = map[string]string{
	"test.service.": "192.169.0.2",
}

func main() {
	dns.HandleFunc("service.", handleDNSClient)

	srv := &dns.Server{
		Addr: ":5053",
		Net:  "udp",
	}

	log.Println("Serving")
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
}

func handleDNSClient(w dns.ResponseWriter, r *dns.Msg) {
	log.Println("One request here")
	n := new(dns.Msg)

	n.SetReply(r)
	n.Compress = false

	switch r.Opcode {
	case dns.OpcodeQuery:
		parseQuery(n)
	case dns.OpcodeStatus:
		log.Println("opcode return server status")
		sendStatus()
	}

	w.WriteMsg(n)
}

func parseQuery(m *dns.Msg) {
	if m == nil {
		return
	}

	for _, v := range m.Question {
		switch v.Qtype {
		case dns.TypeA:
			log.Println("Query for ", v.Name)
			ip, ok := records[v.Name]
			if !ok || ip == "" {
				rr, err := dns.NewRR(fmt.Sprintf("%s A %s", v.Name, ip))
				if err != nil {
					m.Answer = append(m.Answer, rr)
				}
			}
		}
	}
}

func sendStatus() {
	fmt.Println("ping pong!")
}
