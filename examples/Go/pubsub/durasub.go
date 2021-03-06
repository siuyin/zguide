//
// Durable subscriber.
//
// Author: Alec Thomas <alec@swapoff.org>
// Requires: http://github.com/alecthomas/gozmq
//
package main

import (
	zmq "github.com/alecthomas/gozmq"
)

func main() {
	context, _ := zmq.NewContext()
	defer context.Close()

	//  Connect our subscriber socket
	subscriber, _ := context.NewSocket(zmq.SUB)
	defer subscriber.Close()
	subscriber.SetSockOptString(zmq.IDENTITY, "Hello")
	subscriber.SetSockOptString(zmq.SUBSCRIBE, "")
	subscriber.Connect("tcp://localhost:5565")

	//  Synchronize with publisher
	sync, _ := context.NewSocket(zmq.PUSH)
	defer sync.Close()
	sync.Connect("tcp://localhost:5564")
	sync.Send([]byte{}, 0)

	//  Get updates, expect random Ctrl-C death
	for {
		data, _ := subscriber.Recv(0)
		str := string(data)
		println(str)
		if str == "END" {
			break
		}
	}
}
