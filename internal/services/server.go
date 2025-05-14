package services

import "net"

type Server struct {
	ListenAddr string
	ln         net.Listener
}
