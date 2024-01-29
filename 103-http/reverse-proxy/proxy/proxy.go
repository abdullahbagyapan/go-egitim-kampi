package proxy

import (
	"github.com/gofiber/fiber/v2"
)

type ProxyInterface interface {
	Proxy() string
}

type Proxy struct {
	Destination string
}

type HandlerInterface interface {
	ProxyHandler(c *fiber.Ctx) error
}

type Handler struct {
}

var i = 0 // counter

var proxies = []ProxyInterface{
	Proxy{Destination: "127.0.0.2"},
	Proxy{Destination: "127.0.0.3"},
	Proxy{Destination: "127.0.0.4"},
	Proxy{Destination: "127.0.0.5"},
	Proxy{Destination: "127.0.0.6"},
}

func NewHandler() *Handler {
	return &Handler{}
}

func (p Proxy) Proxy() string {
	return p.Destination
}

func (h Handler) ProxyHandler(c *fiber.Ctx) error {

	proxy := randomProxy()

	return c.Redirect(proxy.Proxy())
}

func randomProxy() ProxyInterface {

	len := len(proxies)
	i++
	return proxies[i%len]
}
