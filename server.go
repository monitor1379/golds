package golds

import (
	"fmt"
	"net"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-25 13:00:46
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-26 16:28:42
 */

const (
	defaultNetwork = "tcp"
)

type Server struct {
}

func NewServer() *Server {
	s := new(Server)
	return s
}

func (s *Server) Listen(address string) error {
	logger.Sugar().Infof("Server listening %s", address)
	listener, err := net.Listen(defaultNetwork, address)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			// TODO(monitor1379): log this error
			fmt.Println(err)
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	for {
		// Read request

		// process

		// Write response

		break
	}
}
