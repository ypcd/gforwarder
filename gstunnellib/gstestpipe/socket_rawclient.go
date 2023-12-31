package gstestpipe

import (
	"errors"
	"io"
	"net"
	"sync"
	"sync/atomic"

	"github.com/ypcd/gforwarder/gstunnellib"
)

type RawClientSocket struct {
	close           atomic.Bool
	serverAddr      string
	connListIn      []net.Conn
	lock_connListIn sync.Mutex
}

func NewRawClientSocket(serverAddr string) *RawClientSocket {
	return &RawClientSocket{serverAddr: serverAddr}
}

func (ss *RawClientSocket) createConn() net.Conn {
	client, err := net.Dial("tcp4", ss.serverAddr)
	checkError_panic(err)
	ss.lock_connListIn.Lock()
	ss.connListIn = append(ss.connListIn, client)
	ss.lock_connListIn.Unlock()
	return client
}

func (ss *RawClientSocket) Run() net.Conn {
	return ss.createConn()
}

func (ss *RawClientSocket) Close() {
	ss.close.Swap(true)

	ss.lock_connListIn.Lock()
	defer ss.lock_connListIn.Unlock()
	for _, v := range ss.connListIn {
		v.Close()
	}
}

type RawClientSocketEcho struct {
	RawClientSocket
}

func NewRawClientSocketEcho(serverAddr string) *RawClientSocketEcho {
	return &RawClientSocketEcho{RawClientSocket: RawClientSocket{serverAddr: serverAddr}}
}

func (ss *RawClientSocketEcho) createEchoHandler(client net.Conn) {
	if gstunnellib.G_ants_pool != nil {
		gstunnellib.G_ants_pool.Submit(func() {
			defer client.Close()

			_, err := io.Copy(client, client)
			if errors.Is(err, net.ErrClosed) {
				checkError_info(err)
				return
			}
			if err != nil {
				checkError(err)
				return
			}

		})
	} else {
		go func() {
			defer client.Close()

			_, err := io.Copy(client, client)
			if errors.Is(err, net.ErrClosed) {
				checkError_info(err)
				return
			}
			if err != nil {
				checkError(err)
				return
			}

		}()
	}
}

func (ss *RawClientSocketEcho) Run() {
	ss.createEchoHandler(ss.createConn())
}
