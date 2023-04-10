package nti

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() { plugin.Register("nti", setup) }

func setup(c *caddy.Controller) error {
	c.Next() // 'whoami'
	if c.NextArg() {
		return plugin.Error("nti", c.ArgErr())
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Nti{}
	})

	return nil
}
