# socket
Socket Framework for Golang

## Example:

- Example 1

```
import (
	...

	"github.com/yodstar/goutil/socket"
)

// ServerContext
type ServerContext struct {
	socket.Context
}

// ServerContext.Handshake
func (c *ServerContext) Handshake(conn *socket.Conn) error {
	socket.Message.Send(conn, "hello client!")
	return nil
}

// ServerContext.Handle
func (c *ServerContext) Handle(conn *socket.Conn, bufr io.Reader) error {
	data, err := ioutil.ReadAll(bufr)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(data))
	}
	return err
}

func main() {

	if err := socket.ListenTCP(":18080", &ServerContext{}); err != nil {
		panic(err.Error())
	}

}


```

- Example 2

```

// ClientContext
type ClientContext struct {
	socket.Context
}

// ClientContext.Handshake
func (c *ClientContext) Handshake(conn *socket.Conn) error {
	socket.Message.Send(conn, "hello server!")
	return nil
}

// ClientContext.Handle
func (c *ClientContext) Handle(conn *socket.Conn, bufr io.Reader) error {
	data, err := ioutil.ReadAll(bufr)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(data))
	}
	return err
}

func main() {

	if err := socket.DialTCP("127.0.0.1:18080", &ClientContext{}); err != nil {
		panic(err.Error())
	}

}


```