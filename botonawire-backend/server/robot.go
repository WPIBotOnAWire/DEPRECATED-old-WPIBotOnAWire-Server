package server

import (
	"net"
	"sync"
	"time"
)

type Robot struct {
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	Charge    int    `json:"charge"`
	Status    string `json:"status"`
	Connected bool   `json:"connected"`
}

type Client struct {
	Uuid                string
	Name                string
	Charge              int
	Status              string
	Connected           bool
	Connection          net.Conn
	Ip                  string
	Port                int
	AutoMode            bool
	TimeOfLastHeartbeat time.Time
	Mutex               sync.Mutex
}

func client_to_robot(c *Client) Robot {
	var r Robot
	r.Name = c.Name
	r.Charge = c.Charge
	r.Status = c.Status
	r.Connected = c.Connected
	r.Uuid = c.Uuid
	return r
}
