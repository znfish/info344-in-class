package main

import "fmt"
import "net/http"
import "log"
import "runtime"
import "encoding/json"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	w.Header().Add("Content-Type", "text/plain")
	//w.Write([]byte("Hello, World"))
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Hello %s!", name)
}

func memoryHandler(w http.ResponseWriter, r *http.Request) {
	runtime.GC()                 // force the garbage collector to run, don't really have to do it in a GO server
	stats := &runtime.MemStats{} // creating an instant of the structure (a pointer)
	runtime.ReadMemStats(stats)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats) // encoded into json

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/memory", memoryHandler)
	fmt.Printf("server is listening at http://localhost:4000\n")
	log.Fatal(http.ListenAndServe("localhost:4000", mux))
}
