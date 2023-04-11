package nti

import (
	"context"
	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/miekg/dns"
	"net"
)

const name = "nti"

var log = clog.NewWithPlugin("nti")

type SDKPlugin struct {
	blackDNSAddress string
	whiteDNSAddress string
}

func (h SDKPlugin) Name() string {
	return name
}

func (h SDKPlugin) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	m := new(dns.Msg)
	m.SetReply(r)
	log.Infof("handle the domain: %s, length: %d", r.Question[0].Name, len(r.Question[0].Name))
	// 检查查询是否在黑名单中
	if len(r.Question[0].Name) > 13 {
		log.Info("black")
		// 将查询转发到黑名单 DNS
		blackServer := net.JoinHostPort(h.blackDNSAddress, "53")
		c := new(dns.Client)
		resp, _, err := c.Exchange(r, blackServer)
		if err != nil {
			log.Info(err.Error())
			return dns.RcodeServerFailure, err
		}
		m.Answer = resp.Answer
		m.Ns = resp.Ns
		m.Extra = resp.Extra
	} else {
		// 将查询转发到白名单 DNS
		log.Info("white")
		whiteServer := net.JoinHostPort(h.whiteDNSAddress, "53")
		c := new(dns.Client)
		resp, _, err := c.Exchange(r, whiteServer)
		if err != nil {
			return dns.RcodeServerFailure, err
		}
		m.Answer = resp.Answer
		m.Ns = resp.Ns
		m.Extra = resp.Extra
	}

	err := w.WriteMsg(m)
	if err != nil {
		return dns.RcodeServerFailure, err
	}

	return dns.RcodeSuccess, nil
}
