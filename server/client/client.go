package client

import (
	"io"
	"net"

	"encoding/json"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type Data struct {
	Id      int
	conn    net.Conn
	reader  *wsutil.Reader
	decoder *json.Decoder
	writer  *wsutil.Writer
	encoder *json.Encoder
}

func (client *Data) SendMessage(message interface{}) error {
	if err := client.encoder.Encode(&message); err != nil {
		return err
	}
	if err := client.writer.Flush(); err != nil {
		return err
	}
	return nil
}

func (client *Data) WaitMessage() error {
	hdr, err := client.reader.NextFrame()
	if err != nil {
		return err
	}
	if hdr.OpCode == ws.OpClose {
		return io.EOF
	}
	return nil
}

func (client *Data) DecodeMessageReceived(message interface{}) error {
	if err := client.decoder.Decode(&message); err != nil {
		return err
	}
	return nil
}

func (client *Data) Close() {
	client.conn.Close()
}

func New(id int, conn net.Conn) *Data {
	var (
		reader  = wsutil.NewReader(conn, ws.StateServerSide)
		decoder = json.NewDecoder(reader)
		writer  = wsutil.NewWriter(conn, ws.StateServerSide, ws.OpText)
		encoder = json.NewEncoder(writer)
	)
	client := Data{id, conn, reader, decoder, writer, encoder}
	return &client
}
