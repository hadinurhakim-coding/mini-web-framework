package framework

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON mengirimkan response dalam format JSON.
func (c *Context) JSON(code int, obj any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// String mengirimkan response teks biasa.
func (c *Context) String(code int, format string, values ...any) {
	c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.Writer.WriteHeader(code)
	if len(values) > 0 {
		fmt.Fprintf(c.Writer, format, values...)
		return
	}
	c.Writer.Write([]byte(format))
}

// AbortWithStatus menghentikan proses dan kirim status code.
func (c *Context) AbortWithStatus(code int) {
	c.Status(code)
	c.Abort()
}

// Status menetapkan HTTP status code tanpa menulis body.
func (c *Context) Status(code int) {
	c.Writer.WriteHeader(code)
}

// AbortWithError mempermudah pengiriman error seragam.
func (c *Context) AbortWithError(err *AppError) {
	c.JSON(err.Code, err)
	c.Abort()
}