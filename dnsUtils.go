//additional fields to include net blocks in geodns
package dnsutils

import (
	au "github.com/gigaryte/addrUtils"
	"github.com/miekg/dns"
)

const (
	TypeNet uint16 = 1024
)

// Net RR (this is a fake construct to return unique IPv6 address each time)
type Net struct {
	Hdr    dns.RR_Header
	V6addr *au.IPv6Addr
	Mask   uint8
}

func (rr *Net) String() string {
	return rr.Hdr.String()
}

func (rr *Net) Next() {
}
