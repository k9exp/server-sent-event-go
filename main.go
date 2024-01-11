package main

import (
    "fmt"
    "strconv"
	"net/http"
	"os"
	"time"
)

func getSSEData(eventName string, data any, id, retry uint64) string {
    return fmt.Sprintf("event: %s\ndata: %s\nid: %d\nretry: %d\n\n", eventName, data, id, retry)
}


func main(){
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        html, err := os.ReadFile("./index.html")
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte("Something went wrong"))
            return
        }

        w.Header().Set("Content-Type", "text/html")
        w.Write(html)
    })

    http.HandleFunc("/sse", func (w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/event-stream")
        w.Header().Set("Cache-Control", "no-cache")
        w.Header().Set("Connection", "keep-alive")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")


        flusher := w.(http.Flusher) // type casting
    
        i := 0
        for {
            fmt.Fprintf(w, getSSEData("message", strconv.Itoa(i), uint64(i), 10000))
            i += 1
            flusher.Flush()
            time.Sleep(1000 * time.Millisecond)
        }
    })

    println("Listening at http://localhost:3000")
    http.ListenAndServe(":3000", nil)
}

