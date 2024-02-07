package chat

import "github.com/gorilla/websocket"

type Client struct {
	conn *websocket.Conn
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		conn: conn,
	}
}

func (c *Client) SendMessage(msg []byte) error {
	return c.conn.WriteMessage(websocket.TextMessage, msg)
}
