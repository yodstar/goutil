package socket

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
	"strings"
	"time"
)

// Conn
type Conn struct {
	net.Conn
	Ctx IContext
}

// Conn.Receive
func (c *Conn) Receive() (io.Reader, error) {
	return c.Ctx.FrameUnpack(c)
}

// Conn.Send
func (c *Conn) Send(data []byte) (int, error) {
	pack, err := c.Ctx.FramePack(data)
	if err != nil {
		return 0, sockError(err)
	}
	return c.Write(pack)
}

// tcpserve
func tcpserve(listener net.Listener, ctx IContext) (err error) {
	var conn net.Conn
	defer func() {
		if ctx.IsRunning() {
			if err = listener.Close(); err != nil {
				log.Println(sockError(err))
			}
		}
	}()
	for ctx.IsRunning() {
		conn, err = listener.Accept()
		if err != nil {
			if ctx.IsRunning() {
				log.Println(sockError(err))
			}
			time.Sleep(5 * time.Millisecond)
			continue
		}
		go func(c net.Conn) {
			err = ctx.ServeFunc(c)
		}(conn)
	}
	return
}

// ListenTCP
func ListenTCP(address string, ctx IContext) error {
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return sockError(err)
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return sockError(err)
	}
	err = ctx.init(ctx, listener, false)
	if err != nil {
		return sockError(err)
	}
	return tcpserve(listener, ctx)
}

// DialTCP
func DialTCP(address string, ctx IContext) error {
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return sockError(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return sockError(err)
	}
	err = ctx.init(ctx, nil, true)
	if err != nil {
		return sockError(err)
	}
	return ctx.ServeFunc(conn)
}

// ListenTLS
func ListenTLS(addr, certFile, keyFile string, ctx IContext) error {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return sockError(err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	listener, err := tls.Listen("tcp", addr, config)
	if err != nil {
		return sockError(err)
	}
	err = ctx.init(ctx, listener, false)
	if err != nil {
		return sockError(err)
	}
	return tcpserve(listener, ctx)
}

// DialTLS
func DialTLS(addr string, config *tls.Config, ctx IContext) error {
	conn, err := tls.Dial("tcp", addr, config)
	if err != nil {
		return sockError(err)
	}
	err = ctx.init(ctx, nil, true)
	if err != nil {
		return sockError(err)
	}
	return ctx.ServeFunc(conn)
}

// ListenUDP
func ListenUDP(address string, ctx IContext) error {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return sockError(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return sockError(err)
	}
	err = ctx.init(ctx, nil, false)
	if err != nil {
		return sockError(err)
	}
	return ctx.ServeFunc(conn)
}

// DialUDP
func DialUDP(address string, ctx IContext) error {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return sockError(err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return sockError(err)
	}
	err = ctx.init(ctx, nil, true)
	if err != nil {
		return sockError(err)
	}
	return ctx.ServeFunc(conn)
}

func sockError(err error) error {
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if !ok {
			file = "???"
			line = 0
		}
		file = file[strings.LastIndex(file, "/goutil")+1:]
		err = fmt.Errorf("%s (%s:%d)", err, file, line)
	}
	return err
}
