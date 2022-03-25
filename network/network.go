package network

import (
	"bytes"
	"fmt"
	"github.com/mark2185/pomogoro/timer"
	"net"
	"os"
)

const SockAddr = "/tmp/pomogoro.sock"

func Connect() net.Conn {
	conn, err := net.Dial("unix", SockAddr)
	if err != nil {
		fmt.Println("Cannot dial because...")
		panic(err.Error())
	}
	return conn
}

func Listen(t *timer.Timer) {
	os.Remove(SockAddr)
	l, err := net.Listen("unix", SockAddr)
	defer l.Close()
	if err != nil {
		fmt.Println("Listen error: ", err.Error())
		panic(err.Error())
	}
	buf := make([]byte, 4096)
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err.Error())
		}
		length, err := conn.Read(buf)
		switch {
		case bytes.Compare(buf[:length], []byte("pause")) == 0:
			t.Pause()
		case bytes.Compare(buf[:length], []byte("reset")) == 0:
			t.Reset()
		case bytes.Compare(buf[:length], []byte("toggle")) == 0:
			t.Toggle()
		case bytes.Compare(buf[:length], []byte("resume")) == 0:
			t.Resume()
		case bytes.Compare(buf[:length], []byte("stop")) == 0:
			t.Stop()
		case bytes.Compare(buf[:length], []byte("switch")) == 0:
			t.Switch()
		case bytes.Compare(buf[:length], []byte("increaseTime")) == 0:
			t.UpdateTime(1, 0)
		case bytes.Compare(buf[:length], []byte("decreaseTime")) == 0:
			t.UpdateTime(-1, 0)
		default:
			panic("Message I received is:" + string(buf[:length]))
		}
	}
}
