package client

import (
	"fmt"

	"gopkg.ilharper.com/koi/core/god/proto"
	"gopkg.ilharper.com/koi/core/util/net"
)

func Ping(conn *Options) error {
	var err error

	ws, err := Connect(conn)
	if err != nil {
		return fmt.Errorf("failed to connect to daemon: %w", err)
	}

	request := proto.NewRequest("ping", nil)

	err = net.JSON.Send(ws, request)
	if err != nil {
		return fmt.Errorf("websocket send error: %w", err)
	}

	var resp proto.Response
	err = net.JSON.Receive(ws, &resp)
	if err != nil {
		return fmt.Errorf("websocket receive error: %w", err)
	}
	if resp.Type != "pong" {
		return fmt.Errorf("pingpong failed: response not 'pong'")
	}

	return nil
}
