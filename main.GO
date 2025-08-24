package main

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/belguitR/goChatApp/server"
)

func main() {
	hub := server.NewHub()

	http.HandleFunc("/online", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
    w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    w.Header().Set("Content-Type", "application/json")

    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusNoContent)
        return
    }

    count := hub.OnlineCount()
    _ = json.NewEncoder(w).Encode(map[string]int{"online": count})
})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		server.HandleWebSocket(hub, w, r)
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
