package server

import (
	"net"
	"sync"
	"time"
)

type Robot struct {
	Uuid      string  `json:"uuid"`
	Name      string  `json:"name"`
	Charge    int     `json:"charge"`
	Status    string  `json:"status"`
	Connected bool    `json:"connected"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
	Speed     float64 `json:"speed"`
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
	Latitude            float64
	Longitude           float64
	Altitude            float64
	Speed               float64
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
	r.Latitude = c.Latitude
	r.Longitude = c.Longitude
	r.Altitude = c.Altitude
	r.Speed = c.Speed
	return r
}
