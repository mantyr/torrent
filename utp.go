package torrent

import (
	"net"
	"time"
)

// Abstracts the utp Socket, so the implementation can be selected from
// different packages.
type utpSocket interface {
	Accept() (net.Conn, error)
	Addr() net.Addr
	Close() error
	LocalAddr() net.Addr
	ReadFrom([]byte) (int, net.Addr, error)
	SetDeadline(time.Time) error
	SetWriteDeadline(time.Time) error
	SetReadDeadline(time.Time) error
	WriteTo([]byte, net.Addr) (int, error)
	DialTimeout(string, time.Duration) (net.Conn, error)
	Dial(string) (net.Conn, error)
}
