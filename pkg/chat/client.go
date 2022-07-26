package chat

import (
	"net"
	"net/http"
	"strings"
)

type Client struct {
	conn net.Conn
	name string
	room *room
}

func (c *Client) readInput(w http.ResponseWriter, r *http.Request) {
	for {
		msg := r.FormValue("message")
		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])
	}
}

func (c *Client) mesg(msg string) {
	c.conn.Write([]byte(msg))
}
