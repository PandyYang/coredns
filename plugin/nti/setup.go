package nti

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"

	//clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/coredns/coredns/plugin"
)

//var log = clog.NewWithPlugin("nti")

func init() {
	plugin.Register("nti", func(c *caddy.Controller) error {
		h := new(SDKPlugin)
		err := h.setup(c)
		if err != nil {
			return err
		}
		return nil
	})
}

func (h SDKPlugin) setup(c *caddy.Controller) error {

	for c.Next() {
		log.Infof("load %s Corefile", c.Val())
		if c.Val() == "nti" {
			for c.NextBlock() {
				switch c.Val() {
				case "dns_black_address":
					arg := c.RemainingArgs()[0]
					log.Infof("black address %s", arg)
					h.blackDNSAddress = arg
				case "dns_white_address":
					arg := c.RemainingArgs()[0]
					log.Infof("white address %s", arg)
					h.whiteDNSAddress = arg
				}
			}
		}
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return h
	})
	log.Infof("add %s to plugin success", c.Val())
	return nil
}
