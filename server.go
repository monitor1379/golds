package golds

import (
	"net"

	"github.com/syndtr/goleveldb/leveldb"
	"go.uber.org/zap"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-25 13:00:46
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-26 19:28:27
 */

const (
	defaultNetwork = "tcp"
)

type Server struct {
	db *leveldb.DB
}

func NewServer(db *leveldb.DB) *Server {
	server := new(Server)
	server.db = db
	return server
}

func (this *Server) Listen(address string) error {
	logger.Sugar().Infof("Server listening %s", address)

	listener, err := net.Listen(defaultNetwork, address)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("listener accept failed", zap.Error(err))
			continue
		}

		go this.handleConn(conn)
	}
}

func (this *Server) handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		// Read request

		// process

		// Write response

		break
	}
}
