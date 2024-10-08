package main

import (
	"context"
	"live-wave/internals/chat"
	"live-wave/internals/video"
	"log"
	"net/http"
)

func main() {
	setupApp()

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func setupApp() {
	http.Handle("/", http.FileServer(http.Dir("./frontend/views")))

	output := http.FileServer(http.Dir("./frontend/static"))
	http.Handle("/static/", http.StripPrefix("/static/", output))

	ctx := context.Background()

	video.AllRooms.Init()

	//create new manager for websocket traffic
	manager := chat.NewManager(ctx)
	http.HandleFunc("/ws", manager.ServeWebSocket)
	http.HandleFunc("/create-room", video.CreateRoomRequestHandler)
	http.HandleFunc("/join-room", video.JoinRoomRequestHandler)
}
