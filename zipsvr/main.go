package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/znfish/info344-in-class/zipsvr/handlers"
	"github.com/znfish/info344-in-class/zipsvr/models"
)

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
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	tlskey := os.Getenv("TLSKEY")
	tlsscert := os.Getenv("TLSCERT")
	if len(tlskey) == 0 || len(tlsscert) == 0 {
		log.Fatal("please set TLSKEY and TLSCERT")
	}
	zips, err := models.LoadZips("zips.csv")
	if err != nil {
		log.Fatalf("error loading zips: %v", err) // it will terminate the server, DO NOT use it in http handler
	}
	log.Printf("loaded %d zips", len(zips))

	cityIndex := models.ZipIndex{}
	for _, z := range zips { //needs the _ thing even if we dont care about it
		cityLower := strings.ToLower(z.City)
		cityIndex[cityLower] = append(cityIndex[cityLower], z)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/memory", memoryHandler)

	cityHandler := &handlers.CityHandler{
		Index:      cityIndex,
		PathPrefix: "/zips/",
	}
	mux.Handle("/zips/", cityHandler)
	fmt.Printf("server is listening at https://%s\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsscert, tlskey, mux))
}
