package socket

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"google.golang.org/protobuf/proto"
)

// Codec
type Codec struct {
	Marshal   func(v interface{}) (data []byte, err error)
	Unmarshal func(data []byte, v interface{}) (err error)
}

// Codec.Send
func (c Codec) Send(conn *Conn, v interface{}) error {
	data, err := c.Marshal(v)
	if err != nil {
		return sockError(err)
	}
	_, err = conn.Send(data)
	return sockError(err)
}

// Codec.Receive
func (c Codec) Receive(conn *Conn, v interface{}) error {
	bufr, err := conn.Receive()
	if err != nil {
		return sockError(err)
	}
	data, err := ioutil.ReadAll(bufr)
	if err != nil {
		return sockError(err)
	}
	return sockError(c.Unmarshal(data, v))
}

// messageMarshal
func messageMarshal(v interface{}) ([]byte, error) {
	switch data := v.(type) {
	case string:
		return []byte(data), nil
	case []byte:
		return data, nil
	}
	return nil, sockError(fmt.Errorf("not supported"))
}

// messageUnmarshal
func messageUnmarshal(msg []byte, v interface{}) error {
	switch data := v.(type) {
	case *string:
		*data = string(msg)
		return nil
	case *[]byte:
		*data = msg
		return nil
	}
	return sockError(fmt.Errorf("not supported"))
}

/*
Message is a codec to send/receive text/binary data in a frame on connection.
To send/receive text frame, use string type.
To send/receive binary frame, use []byte type.

Trivial usage:

	import "github.com/zmrnet/findgo/socket"

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
*/
var Message = Codec{messageMarshal, messageUnmarshal}

// jsonMarshal
func jsonMarshal(v interface{}) (msg []byte, err error) {
	msg, err = json.Marshal(v)
	return msg, sockError(err)
}

// jsonUnmarshal
func jsonUnmarshal(msg []byte, v interface{}) error {
	return sockError(json.Unmarshal(msg, v))
}

/*
JSON is a codec to send/receive JSON data in a frame from a connection.

Trivial usage:

	import "github.com/zmrnet/findgo/socket"

	type T struct {
		Msg string
		Count int
	}

	// receive JSON type T
	var data T
	socket.JSON.Receive(conn, &data)

	// send JSON type T
	socket.JSON.Send(conn, data)
*/
var JSON = Codec{jsonMarshal, jsonUnmarshal}

// protoMarshal
func protoMarshal(v interface{}) (msg []byte, err error) {
	msg, err = proto.Marshal(v.(proto.Message))
	return msg, sockError(err)
}

// protoUnmarshal
func protoUnmarshal(msg []byte, v interface{}) error {
	return sockError(proto.Unmarshal(msg, v.(proto.Message)))
}

/*
Protobuf is a codec to send/receive Protobuf data in a frame from a connection.

Trivial usage:

	import "github.com/zmrnet/findgo/socket"

	// receive Protobuf type T
	var data pb.T
	socket.Protobuf.Receive(conn, &data)

	// send Protobuf type T
	socket.Protobuf.Send(conn, &data)
*/
var Protobuf = Codec{protoMarshal, protoUnmarshal}
