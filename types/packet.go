package types

import (
	"net"
	"strings"
)

type Packet []byte

func (p *Packet) Read(conn net.Conn) (msg string, err error) {
	buffer := *p
	n, err := conn.Read(buffer)
	if err != nil {
		return
	}

	if n == 0 {
		return "", nil
	}

	if string(buffer[:n]) == string([]byte{13, 0}) {
		return "", nil
	}

	in := strings.TrimSpace(string(buffer[:n]))
	if in == "" {
		return "", nil
	}

	return in, nil
}
