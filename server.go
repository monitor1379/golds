package golds

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-25 13:00:46
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-27 23:24:17
 */

import (
	"errors"
	"io"
	"net"

	"github.com/monitor1379/golds/goldscore"
	"github.com/monitor1379/golds/handlers"
	"github.com/syndtr/goleveldb/leveldb"
	"go.uber.org/zap"
)

const (
	defaultNetwork = "tcp"
)

var (
	ErrInvalidRequestPacketValueLength = errors.New("invalid request packet value length")
)

type Server struct {
	db     *leveldb.DB
	router *goldscore.Router
}

func NewServer(db *leveldb.DB) *Server {
	server := new(Server)
	server.db = db
	server.router = goldscore.NewRouter()
	server.router.AddHandleFunc("set", handlers.Set)
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

		logger.Debug("accept connection", zap.String("ip", conn.RemoteAddr().String()))
		go this.handleConn(conn)
	}
}

func (this *Server) handleConn(conn net.Conn) {
	defer func() {
		conn.Close()
		logger.Debug("connection close")
	}()

	packetEncoder := NewPacketEncoder(conn)
	packetDecoder := NewPacketDecoder(conn)

	for {
		packet, err := packetDecoder.Decode()
		if err == io.EOF {
			break
		}

		if err != nil {
			logger.Error("decode packet failed", zap.Error(err))
			break
		}

		logger.Debug("accept packet", zap.String("packet", packet.String()))

		err = this.route(packet, packetEncoder, packetDecoder)

		if err != nil {
			logger.Error("handle request failed", zap.Error(err))
			break
		}
	}
}

func (this *Server) route(packet *Packet, packetEncoder *PacketEncoder, packetDecoder *PacketDecoder) error {
	if len(packet.Value) == 0 {
		return ErrInvalidRequestPacketValueLength
	}

	// route
	if len(packet.Value) == 3 && string(packet.Value[0]) == "SET" {
		// set <key> <value>

	} else if len(packet.Value) == 2 && string(packet.Value[0]) == "GET" {
		// get <key>

	}

	err := packetEncoder.Encode(OKPacket)
	if err != nil {
		return err
	}

	return nil
}
