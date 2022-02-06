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
	Hdr   dns.RR_Header
	V6Net au.IPv6Network `dns:"net"`
}

func (rr *Net) Header() *dns.RR_Header {
	return &rr.Hdr
}

func (rr *Net) String() string {
	return rr.Hdr.String()
}

func (rr *Net) Next() {
}

func (rr *Net) Len() {
}

func (rr *Net) Copy() {
}

func (rr *Net) Pack() {

}
func (rr *Net) Unpack() {
}

func (rr *Net) Parse() {
}
func (rr *Net) isDuplicate() {
}
