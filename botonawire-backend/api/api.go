package api

import (
	"botonawire/server"
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetRobots(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(server.GetRobots())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ForwardRobot(w http.ResponseWriter, r *http.Request) {
	robot := new(server.Robot)

	r.ParseForm()
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.Unmarshal(bytes, robot)

	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, 1600)

	server.SendPacketToClient(robot.Uuid, 2, buf)
}

func BackwardRobot(w http.ResponseWriter, r *http.Request) {
	robot := new(server.Robot)

	r.ParseForm()
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.Unmarshal(bytes, robot)

	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, 1400)

	server.SendPacketToClient(robot.Uuid, 2, buf)
}

func StopRobot(w http.ResponseWriter, r *http.Request) {
	robot := new(server.Robot)

	r.ParseForm()
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.Unmarshal(bytes, robot)

	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, 1500)

	server.SendPacketToClient(robot.Uuid, 2, buf)
}
