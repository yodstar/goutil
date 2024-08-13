package socket_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"time"

	"github.com/yodstar/goutil/socket"
)

// Context
type Context struct {
	socket.Context
}

// Context.Handshake
func (c *Context) Handshake(conn *socket.Conn) (err error) {
	if c.IsClient() {
		time.Sleep(time.Second)
		err = socket.Message.Send(conn, "hello server!")
	} else {
		err = socket.Message.Send(conn, "hello client!")
	}
	return
}

// Context.Handle
func (c *Context) Handle(conn *socket.Conn, bufr io.Reader) error {
	if data, err := ioutil.ReadAll(bufr); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(data))
	}
	time.Sleep(2 * time.Second)
	return io.EOF
}

// ExampleListenTCP
func ExampleListenTCP() {
	log.SetFlags(log.Lshortfile)

	ctx := &Context{}

	go func(ctx *Context) {
		time.Sleep(time.Second)

		addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:18080")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		conn, err := net.DialTCP("tcp", nil, addr)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		rwc := &socket.Conn{conn, ctx}

		err = socket.Message.Send(rwc, "hello server!")
		if err != nil {
			fmt.Println(err.Error())
		}

		time.Sleep(time.Second)

		var msg string
		for {
			if err := socket.Message.Receive(rwc, &msg); err != nil {
				fmt.Println(err.Error())
				break
			}
			fmt.Println(msg)
		}

		conn.Close()

		time.Sleep(time.Second)

		_ = ctx.Shutdown()
	}(ctx)

	if err := socket.ListenTCP(":18080", ctx); err != nil {
		log.Println(err.Error())
	}

	// Output:
	// hello server!
	// hello client!
	// EOF
}

// ExampleDialTCP
func ExampleDialTCP() {
	log.SetFlags(log.Lshortfile)

	ctx := &Context{}

	go func() {
		if err := socket.ListenTCP(":18080", ctx); err != nil {
			log.Println(err.Error())
		}
	}()

	time.Sleep(time.Second)

	err := socket.DialTCP("127.0.0.1:18080", &Context{})
	if err != nil && err != io.EOF {
		log.Println(err.Error())
	}

	_ = ctx.Shutdown()

	// Output:
	// hello client!
	// hello server!
}
