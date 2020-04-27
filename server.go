package golds

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-25 13:00:46
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-28 00:57:16
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
	ErrRequestPacketRoutePathNotFound  = errors.New("requet packet route path is not found")
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
	server.router.AddHandleFunc("get", handlers.Get)

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

	packetEncoder := goldscore.NewPacketEncoder(conn)
	packetDecoder := goldscore.NewPacketDecoder(conn)

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
			continue
		}
	}
}

func (this *Server) route(packet *goldscore.Packet, packetEncoder *goldscore.PacketEncoder, packetDecoder *goldscore.PacketDecoder) error {
	if len(packet.Array) == 0 {
		return ErrInvalidRequestPacketValueLength
	}

	routePath := string(packet.Array[0].Value)
	handleFunc, ok := this.router.Route(routePath)
	if !ok {
		return ErrRequestPacketRoutePathNotFound
	}

	// TODO(monitor1379): 如何设计这个handleFunc, context, 以及request/response
	ctx := goldscore.NewContext()
	handleFunc(ctx)

	err := packetEncoder.Encode(goldscore.OKPacket)
	if err != nil {
		return err
	}

	return nil
}
