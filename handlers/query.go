package handlers

import (
	"fmt"
	"log"
	"sync"

	"github.com/miekg/dns"

	"github.com/acentior/dnsf/cache"
)

var ecords = map[string]string{
	"test.service.": "192.169.0.2",
}

type QueryHandler struct {
	lock  sync.Mutex
	Cache *cache.Cache
}

func (h *QueryHandler) HandleDNSClient(w dns.ResponseWriter, r *dns.Msg) {
	log.Printf("%d request here", r.Id)
	n := new(dns.Msg)

	n.SetReply(r)
	n.Compress = false

	switch r.Opcode {
	case dns.OpcodeQuery:
		h.parseQuery(n)
	}

	w.WriteMsg(n)
}

func (h *QueryHandler) parseQuery(m *dns.Msg) {
	if m == nil {
		return
	}

	for _, v := range m.Question {
		switch v.Qtype {
		case dns.TypeA:
			log.Println("Query for ", v.Name)
			ip, err := h.Cache.Get(v.Name)
			if err == nil {
				rr, err := dns.NewRR(fmt.Sprintf("%s A %s", v.Name, ip))
				if err != nil {
					m.Answer = append(m.Answer, rr)
				}
			}
		case dns.TypeAAAA:
			log.Println("Query for and IPv6 address ", v.Name)
			ip, err := h.Cache.Get(v.Name)
			if err == nil {
				rr, err := dns.NewRR(fmt.Sprintf("%s A %", v.Name, ip))
				if err != nil {
					m.Answer = append(m.Answer, rr)
				}
			}
		default:
			log.Println("Requested query not found")
		}
	}
}
