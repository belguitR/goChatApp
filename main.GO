package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"

    "github.com/belguitR/goChatApp/server"
)

func main() {
    hub := server.NewHub()

    http.HandleFunc("/online", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
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

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        p := "./public/browser" + r.URL.Path
        if info, err := os.Stat(p); err == nil && !info.IsDir() {
            http.ServeFile(w, r, p)
            return
        }
        http.ServeFile(w, r, "./public/browser/index.html")
    })

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}