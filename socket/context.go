package socket

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

// FramePrefix - Frame:<[Prefix:<[]byte>] [Length:<uint32>] [Content:<[]byte>]>
var FramePrefix = []byte{'Y', 'O', 'D', 'S', 'T', 'A', 'R'}

// IContext
type IContext interface {
	init(IContext, net.Listener, bool) error
	Initialize() error
	IsRunning() bool
	IsClient() bool
	Shutdown() error
	FrameUnpack(*Conn) (io.Reader, error)
	FramePack([]byte) ([]byte, error)
	ServeFunc(net.Conn) error
	Handshake(*Conn) error
	Handle(*Conn, io.Reader) error
	Destroy(*Conn) error
}

// Context
type Context struct {
	instance IContext
	listener net.Listener
	isclient bool
	running  bool
}

// Context.init
func (c *Context) init(ctx IContext, lis net.Listener, isc bool) (err error) {
	c.instance = ctx
	c.listener = lis
	c.isclient = isc
	if err = ctx.Initialize(); err == nil {
		c.running = true
	}
	return
}

// Context.Initialize
func (c *Context) Initialize() error {
	return nil
}

// Context.IsRunning
func (c *Context) IsRunning() bool {
	return c.running
}

// Context.IsClient
func (c *Context) IsClient() bool {
	return c.isclient
}

// Context.Handshake
func (c *Context) Shutdown() (err error) {
	c.running = false
	if c.listener != nil {
		err = sockError(c.listener.Close())
	}
	return
}

// Context.FrameUnpack
func (c *Context) FrameUnpack(conn *Conn) (io.Reader, error) {
	var err error
	var b byte

	bufr := bufio.NewReader(conn.Conn)
	// Prefix
	pos := 0
	for {
		b, err = bufr.ReadByte()
		if err != nil {
			return nil, sockError(err)
		}
		if b == FramePrefix[pos] {
			pos++
		} else {
			pos = 0
		}
		if pos >= len(FramePrefix) {
			break
		}
	}
	// Length
	var buf = make([]byte, 4)
	for i := 0; i < 4; i++ {
		b, err = bufr.ReadByte()
		if err != nil {
			return nil, sockError(err)
		}
		buf[i] = b
	}
	n := int64(binary.BigEndian.Uint32(buf))
	if n > (1 << 24) {
		return nil, fmt.Errorf("frame too large: %d", n)
	}
	return io.LimitReader(bufr, n), nil
}

// Context.FramePack
func (c *Context) FramePack(data []byte) ([]byte, error) {
	n := len(data)
	if n > (1 << 24) {
		return nil, sockError(fmt.Errorf("frame too large: %d", n))
	}
	// Length
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(n))
	frame := bytes.Join([][]byte{FramePrefix, buf, data}, []byte{})
	return frame, nil
}

// Context.ServeFunc
func (c *Context) ServeFunc(conn net.Conn) (err error) {
	var bufr io.Reader

	ctx := c.instance
	rwc := &Conn{conn, ctx}
	// Destroy
	defer func() {
		ctx.Destroy(rwc)
		if err = conn.Close(); err != nil {
			log.Println(sockError(err))
		}
	}()
	// Handshake
	if err = ctx.Handshake(rwc); err != nil {
		return
	}

	for {
		if bufr, err = rwc.Receive(); err != nil {
			if err != io.EOF && err != io.ErrUnexpectedEOF {
				log.Println(sockError(err))
			}
			break
		}
		if err = ctx.Handle(rwc, bufr); err != nil {
			if err != io.EOF && err != io.ErrUnexpectedEOF {
				log.Println(sockError(err))
			}
			break
		}
	}
	return err
}

// Context.Handshake
func (c *Context) Handshake(conn *Conn) error {
	return nil
}

// Context.Handle
func (c *Context) Handle(conn *Conn, bufr io.Reader) error {
	return nil
}

// Context.Destroy
func (c *Context) Destroy(conn *Conn) error {
	return nil
}
