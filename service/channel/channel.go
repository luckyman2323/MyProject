package channel

import (
	"fmt"
	"sync"
	"time"
)

type Connection struct {
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte
	isClosed  bool
	mutex     sync.Mutex
}

func InitConnection() (conn *Connection, err error) {
	conn = &Connection{
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
		isClosed:  false,
	}

	go conn.readLoop()

	go conn.writeLoop()
	return
}

// 内部实现
func (conn *Connection) readLoop() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("000000")

		select {
		case <-conn.inChan:
			fmt.Println("========readin")
		case <-conn.closeChan:
			goto ERR
		}
	}
ERR:
	fmt.Println("=========1111111")
}

func (conn *Connection) writeLoop() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("11111")
		select {
		case <-conn.outChan:
		case <-conn.closeChan:
			goto ERR
		}
	}
ERR:
	fmt.Println("=========222222")
}

func (conn *Connection) Close() {
	conn.mutex.Lock()
	if !conn.isClosed {
		fmt.Println("=====start close")
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}
