package framework

import (
	"math"
	"net"
	"net/http"
	"strings"
)

const abortIndex int8 = math.MaxInt8 / 2

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Path    string
	Method  string
	Params  Params
	// handlers menyimpan daftar middleware + handler utama
	handlers HandlersChain
	// index melacak middleware mana yang sedang berjalan
	index int8

	// Keys digunakan untuk menyimpan data antar middleware (misal: user_id)
	Keys map[string]any
}

// Next akan menjalankan handler berikutnya dalam chain.
// Digunakan di dalam middleware.
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

// Abort menghentikan eksekusi handler berikutnya.
func (c *Context) Abort() {
	c.index = abortIndex
}

// Set menyimpan data baru ke dalam context.
func (c *Context) Set(key string, value any) {
	if c.Keys == nil {
		c.Keys = make(map[string]any)
	}
	c.Keys[key] = value
}

// Get mengambil data dari context.
func (c *Context) Get(key string) (value any, exists bool) {
	value, exists = c.Keys[key]
	return
}

// Param mengambil parameter path (misal: :id)
func (c *Context) Param(key string) string {
	return c.Params.ByName(key)
}

// ClientIP mencoba mengambil IP asli pengguna.
func (c *Context) ClientIP() string {
	if ip := c.Request.Header.Get("X-Forwarded-For"); ip != "" {
		return strings.Split(ip, ",")[0]
	}
	if ip := c.Request.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}
	host, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
	return host
}