package server

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

type Packet struct {
	Code uint32
	Data []byte
}

var clients map[string]*Client
var uuids map[net.Addr]string

func get_uuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

func GetRobots() []Robot {
	robots := make([]Robot, len(clients))

	i := 0
	for _, r := range clients {
		robots[i] = client_to_robot(r)
		i++
	}

	return robots
}

func fmt_packet(code int, data []byte) []byte {
	var buf = make([]byte, 8+len(data))

	binary.LittleEndian.PutUint32(buf[0:], uint32(code))
	binary.LittleEndian.PutUint32(buf[4:], uint32(len(data)))
	copy(buf[8:], data)

	return buf
}

func recv(conn net.Conn) (Packet, error) {
	var p Packet

	var codeBytes [4]byte
	for n := 0; n < len(codeBytes); {
		c, err := conn.Read(codeBytes[n:])
		if err != nil {
			return p, err
		}
		n += c
	}
	p.Code = binary.LittleEndian.Uint32(codeBytes[:])

	var lengthBytes [4]byte
	for n := 0; n < len(lengthBytes); {
		c, err := conn.Read(lengthBytes[n:])
		if err != nil {
			return p, err
		}
		n += c
	}
	length := binary.LittleEndian.Uint32(lengthBytes[:])

	data := make([]byte, length)
	for n := 0; n < len(data); {
		c, err := conn.Read(data[n:])
		if err != nil {
			return p, err
		}
		n += c
	}

	p.Data = data

	return p, nil
}

func HandleRobot(conn net.Conn, client *Client) {

	for {
		p, err := recv(conn)
		if err != nil {
			conn.Close()
			return
		}

		switch p.Code {
		case 0:
			conn.Write(fmt_packet(0, []byte("")))
			client.Connected = false
			conn.Close()
			return

		case 1:
			conn.Write(fmt_packet(1, []byte(uuids[conn.RemoteAddr()])))
			client.Mutex.Lock()
			client.TimeOfLastHeartbeat = time.Now()
			client.Mutex.Unlock()
		case 2: // GPS Data
			var gps GPS
			json.Unmarshal(p.Data, &gps)
			client.Mutex.Lock()
			client.Latitude = gps.Latitude
			client.Longitude = gps.Longitude
			client.Altitude = gps.Altitude
			client.Speed = gps.Speed
		}
	}
}

func SendPacketToClient(id string, code int, data []byte) error {
	print(id)
	client, ok := clients[id]
	if !ok || (client == nil) || (client.Connection == nil) {
		print("could not find client with given id")
		return errors.New("could not find client with given id")
	}

	fmt.Printf("Sending to client: IP: %s | UUID: %s", client.Ip, client.Uuid)
	_, err := client.Connection.Write(fmt_packet(code, data))
	return err
}

func CheckHeartbeat() {
	for {
		for _, c := range clients {
			c.Mutex.Lock()
			if time.Since(c.TimeOfLastHeartbeat) > time.Duration(time.Second)*60 {
				c.Connected = false
			}
			c.Mutex.Unlock()
		}
		time.Sleep(time.Duration(time.Second) * 5)
	}
}

func Start() {
	l, err := net.Listen("tcp", ":5555")
	if err != nil {
		println("Error opening tcp listener: %s", err.Error())
		os.Exit(1)
	}

	clients = make(map[string]*Client)
	uuids = make(map[net.Addr]string)

	go CheckHeartbeat()

	for {
		conn, err := l.Accept()
		if err != nil {
			println("Error accepting request: %s", err.Error())
			break
		}

		p, err := recv(conn)
		if err != nil {
			println("Error accepting request: %s", err.Error())
			break
		}

		id := string(p.Data)

		if len(id) == 0 {
			id = get_uuid()
		}

		uuid, ok := uuids[conn.RemoteAddr()]
		if !ok || (uuid == "") {
			uuids[conn.RemoteAddr()] = id
			uuid = id
		}

		client, ok := clients[uuid]
		if !ok || (client == nil) {
			clients[uuid] = new(Client)
			client = clients[uuid]
			client.Uuid = uuid
		}

		client.Connection = conn

		conn.Write(fmt_packet(1, []byte(uuids[conn.RemoteAddr()])))

		fmt.Printf("Connection From: %s | uuid: %s\n", conn.RemoteAddr(), uuid)

		client.TimeOfLastHeartbeat = time.Now()
		client.Connected = true
		go HandleRobot(conn, client)
	}
}
