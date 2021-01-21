package net

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewHeaderProxy(dst string) (*HeaderProxy, error) {
	dstURL, err := url.Parse(dst)
	if err != nil {
		return nil, err
	}

	return &HeaderProxy{
		proxy: httputil.NewSingleHostReverseProxy(dstURL),
		headers: map[string]string{
			"NOTICE_PROXY_UTILIZED": "YES",
		},
	}, nil
}

type HeaderProxy struct {
	proxy   *httputil.ReverseProxy
	headers map[string]string
}

func (h *HeaderProxy) Set(key, val string) {
	h.headers[key] = val
}

func (h *HeaderProxy) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	for key, val := range h.headers {
		req.Header.Add(key, val)
	}

	h.proxy.ServeHTTP(res, req)
}

func (proxy *HeaderProxy) StartProxy(ctx context.Context, port *string) error {
	portStr := ":0"
	if port != nil {
		portStr = fmt.Sprintf(":%s", *port)
	}

	listener, err := net.Listen("tcp", portStr)
	if err != nil {
		return err
	}

	listenerPort := listener.Addr().(*net.TCPAddr).Port
	fmt.Printf("starting authentication proxy on http://0.0.0.0:%d\n", listenerPort)

	err = http.Serve(listener, proxy)
	if err != nil {
		return err
	}

	return nil
}
