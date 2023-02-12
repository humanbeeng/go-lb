package client

import (
	"fmt"
	"net"
)

type Client struct {
	Conn net.Conn
}

func New(loadbalanceraddr string) (*Client, error) {
	conn, err := net.Dial("tcp", loadbalanceraddr)
	if err != nil {
		return nil, fmt.Errorf("failed to establish connection with %v", loadbalanceraddr)
	}
	return &Client{Conn: conn}, nil
}
