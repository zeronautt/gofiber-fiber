package main

import (
	"net"
	"strings"
)

type Config struct {
	ProxyHeader string
}

type App struct {
	config Config
}

func New() *App {
	return &App{
		config: Config{},
	}
}

func (a *App) Config() *Config {
	return &a.config
}

type Ctx struct {
	app      *App
	headers  map[string]string
	remoteIP string
}

const HeaderXForwardedFor = "X-Forwarded-For"

func (c *Ctx) App() *App {
	return c.app
}

// IP returns the client IP address.
func (c *Ctx) IP() string {
	proxyHeader := c.app.Config().ProxyHeader
	if proxyHeader != "" {
		headerVal := c.headers[proxyHeader]
		if headerVal != "" {
			ips := strings.Split(headerVal, ",")
			for _, ip := range ips {
				ip = strings.TrimSpace(ip)
				if ip == "" {
					continue
				}
				// Verify if it's a valid IP address
				if net.ParseIP(ip) != nil {
					return ip
				}
			}
		}
	}
	return c.remoteIP
}
