package api

import (
	"net"
)

func ApiUpload() (net.Conn, error) {

	conn, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		return nil, err
	}

	return conn, nil
}
