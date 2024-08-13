

# socket
`import "github.com/yodstar/goutil/socket"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func DialTCP(address string, ctx IContext) error](#DialTCP)
* [func DialTLS(addr string, config *tls.Config, ctx IContext) error](#DialTLS)
* [func DialUDP(address string, ctx IContext) error](#DialUDP)
* [func ListenTCP(address string, ctx IContext) error](#ListenTCP)
* [func ListenTLS(addr, certFile, keyFile string, ctx IContext) error](#ListenTLS)
* [func ListenUDP(address string, ctx IContext) error](#ListenUDP)
* [type Codec](#Codec)
  * [func (c Codec) Receive(conn *Conn, v interface{}) error](#Codec.Receive)
  * [func (c Codec) Send(conn *Conn, v interface{}) error](#Codec.Send)
* [type Conn](#Conn)
  * [func (c *Conn) Receive() (io.Reader, error)](#Conn.Receive)
  * [func (c *Conn) Send(data []byte) (int, error)](#Conn.Send)
* [type Context](#Context)
  * [func (c *Context) Destroy(conn *Conn) error](#Context.Destroy)
  * [func (c *Context) FramePack(data []byte) ([]byte, error)](#Context.FramePack)
  * [func (c *Context) FrameUnpack(conn *Conn) (io.Reader, error)](#Context.FrameUnpack)
  * [func (c *Context) Handle(conn *Conn, bufr io.Reader) error](#Context.Handle)
  * [func (c *Context) Handshake(conn *Conn) error](#Context.Handshake)
  * [func (c *Context) Initialize() error](#Context.Initialize)
  * [func (c *Context) IsClient() bool](#Context.IsClient)
  * [func (c *Context) IsRunning() bool](#Context.IsRunning)
  * [func (c *Context) ServeFunc(conn net.Conn) (err error)](#Context.ServeFunc)
  * [func (c *Context) Shutdown() (err error)](#Context.Shutdown)
* [type IContext](#IContext)

#### <a name="pkg-examples">Examples</a>
* [DialTCP](#example_DialTCP)
* [ListenTCP](#example_ListenTCP)

#### <a name="pkg-files">Package files</a>
[codec.go](/src/github.com/yodstar/goutil/socket/codec.go) [context.go](/src/github.com/yodstar/goutil/socket/context.go) [socket.go](/src/github.com/yodstar/goutil/socket/socket.go) 



## <a name="pkg-variables">Variables</a>
``` go
var FramePrefix = []byte{'F', 'I', 'N', 'D', 'G', 'O', '0', '1'}
```
FramePrefix - Frame:<[Prefix:<[]byte>] [Length:<uint32>] [Content:<[]byte>]>

``` go
var JSON = Codec{jsonMarshal, jsonUnmarshal}
```
JSON is a codec to send/receive JSON data in a frame from a connection.

Trivial usage:


	import "github.com/yodstar/goutil/socket"
	
	type T struct {
		Msg string
		Count int
	}
	
	// receive JSON type T
	var data T
	socket.JSON.Receive(conn, &data)
	
	// send JSON type T
	socket.JSON.Send(conn, data)

``` go
var Message = Codec{messageMarshal, messageUnmarshal}
```
Message is a codec to send/receive text/binary data in a frame on connection.
To send/receive text frame, use string type.
To send/receive binary frame, use []byte type.

Trivial usage:


	import "github.com/yodstar/goutil/socket"
	
	// receive text frame
	var message string
	socket.Message.Receive(conn, &message)
	
	// send text frame
	message = "hello"
	socket.Message.Send(conn, message)
	
	// receive binary frame
	var data []byte
	socket.Message.Receive(conn, &data)
	
	// send binary frame
	data = []byte{0, 1, 2}
	socket.Message.Send(conn, data)

``` go
var Protobuf = Codec{protoMarshal, protoUnmarshal}
```
Protobuf is a codec to send/receive Protobuf data in a frame from a connection.

Trivial usage:


	import "github.com/yodstar/goutil/socket"
	
	// receive Protobuf type T
	var data pb.T
	socket.Protobuf.Receive(conn, &data)
	
	// send Protobuf type T
	socket.Protobuf.Send(conn, &data)



## <a name="DialTCP">func</a> [DialTCP](/src/target/socket.go?s=1276:1324#L78)
``` go
func DialTCP(address string, ctx IContext) error
```
DialTCP



## <a name="DialTLS">func</a> [DialTLS](/src/target/socket.go?s=2078:2143#L113)
``` go
func DialTLS(addr string, config *tls.Config, ctx IContext) error
```
DialTLS



## <a name="DialUDP">func</a> [DialUDP](/src/target/socket.go?s=2704:2752#L143)
``` go
func DialUDP(address string, ctx IContext) error
```
DialUDP



## <a name="ListenTCP">func</a> [ListenTCP](/src/target/socket.go?s=913:963#L61)
``` go
func ListenTCP(address string, ctx IContext) error
```
ListenTCP



## <a name="ListenTLS">func</a> [ListenTLS](/src/target/socket.go?s=1628:1694#L95)
``` go
func ListenTLS(addr, certFile, keyFile string, ctx IContext) error
```
ListenTLS



## <a name="ListenUDP">func</a> [ListenUDP](/src/target/socket.go?s=2354:2404#L126)
``` go
func ListenUDP(address string, ctx IContext) error
```
ListenUDP




## <a name="Codec">type</a> [Codec](/src/target/codec.go?s=111:244#L12)
``` go
type Codec struct {
    Marshal   func(v interface{}) (data []byte, err error)
    Unmarshal func(data []byte, v interface{}) (err error)
}

```
Codec










### <a name="Codec.Receive">func</a> (Codec) [Receive](/src/target/codec.go?s=455:510#L28)
``` go
func (c Codec) Receive(conn *Conn, v interface{}) error
```
Codec.Receive




### <a name="Codec.Send">func</a> (Codec) [Send](/src/target/codec.go?s=260:312#L18)
``` go
func (c Codec) Send(conn *Conn, v interface{}) error
```
Codec.Send




## <a name="Conn">type</a> [Conn](/src/target/socket.go?s=107:151#L15)
``` go
type Conn struct {
    net.Conn
    Ctx IContext
}

```
Conn










### <a name="Conn.Receive">func</a> (\*Conn) [Receive](/src/target/socket.go?s=169:212#L21)
``` go
func (c *Conn) Receive() (io.Reader, error)
```
Conn.Receive




### <a name="Conn.Send">func</a> (\*Conn) [Send](/src/target/socket.go?s=260:305#L26)
``` go
func (c *Conn) Send(data []byte) (int, error)
```
Conn.Send




## <a name="Context">type</a> [Context](/src/target/context.go?s=584:679#L32)
``` go
type Context struct {
    // contains filtered or unexported fields
}

```
Context










### <a name="Context.Destroy">func</a> (\*Context) [Destroy](/src/target/context.go?s=3185:3228#L171)
``` go
func (c *Context) Destroy(conn *Conn) error
```
Context.Destroy




### <a name="Context.FramePack">func</a> (\*Context) [FramePack](/src/target/context.go?s=2000:2056#L113)
``` go
func (c *Context) FramePack(data []byte) ([]byte, error)
```
Context.FramePack




### <a name="Context.FrameUnpack">func</a> (\*Context) [FrameUnpack](/src/target/context.go?s=1319:1379#L75)
``` go
func (c *Context) FrameUnpack(conn *Conn) (io.Reader, error)
```
Context.FrameUnpack




### <a name="Context.Handle">func</a> (\*Context) [Handle](/src/target/context.go?s=3090:3148#L166)
``` go
func (c *Context) Handle(conn *Conn, bufr io.Reader) error
```
Context.Handle




### <a name="Context.Handshake">func</a> (\*Context) [Handshake](/src/target/context.go?s=3009:3054#L161)
``` go
func (c *Context) Handshake(conn *Conn) error
```
Context.Handshake




### <a name="Context.Initialize">func</a> (\*Context) [Initialize](/src/target/context.go?s=926:962#L51)
``` go
func (c *Context) Initialize() error
```
Context.Initialize




### <a name="Context.IsClient">func</a> (\*Context) [IsClient](/src/target/context.go?s=1079:1112#L61)
``` go
func (c *Context) IsClient() bool
```
Context.IsClient




### <a name="Context.IsRunning">func</a> (\*Context) [IsRunning](/src/target/context.go?s=1001:1035#L56)
``` go
func (c *Context) IsRunning() bool
```
Context.IsRunning




### <a name="Context.ServeFunc">func</a> (\*Context) [ServeFunc](/src/target/context.go?s=2350:2404#L126)
``` go
func (c *Context) ServeFunc(conn net.Conn) (err error)
```
Context.ServeFunc




### <a name="Context.Shutdown">func</a> (\*Context) [Shutdown](/src/target/context.go?s=1158:1198#L66)
``` go
func (c *Context) Shutdown() (err error)
```
Context.Handshake




## <a name="IContext">type</a> [IContext](/src/target/context.go?s=250:571#L17)
``` go
type IContext interface {
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
    // contains filtered or unexported methods
}
```
IContext














- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
