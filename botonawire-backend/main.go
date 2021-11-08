package main

import (
	"botonawire/api"
	"botonawire/server"
	"net/http"
)

func startServer() {
	http.Handle("/", http.FileServer(http.Dir("static")))

	http.HandleFunc("/robots", api.GetRobots)
	http.HandleFunc("/forward", api.ForwardRobot)
	http.HandleFunc("/backward", api.BackwardRobot)
	http.HandleFunc("/stop", api.StopRobot)
}

func main() {
	startServer()

	go server.Start()

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		print(err.Error())
	}
}
