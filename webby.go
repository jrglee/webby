package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func iphandler(w http.ResponseWriter, r *http.Request) {
	logIt(w, r)
	fmt.Fprintf(w, GetLocalIP())
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	logIt(w, r)
	fmt.Fprintf(w, time.Now().String())
}

func logIt(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, r.RemoteAddr, r.UserAgent())
}

func main() {
	http.HandleFunc("/", iphandler)
	http.HandleFunc("/time", timeHandler)

	log.Println("Started Webby!")
	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
