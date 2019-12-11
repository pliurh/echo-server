package main

import (
    "fmt"
    "net/http"
    "os"
    "log"
)

// DefaultPort is the default port to use if once is not specified by the SERVER_PORT environment variable
const DefaultPort = "7893";

func getServerPort() (string) {
    port := os.Getenv("SERVER_PORT");
    if port != "" {
        return port;
    }

    return DefaultPort;
}

// EchoHandler echos back the request as a response
func EchoHandler(writer http.ResponseWriter, request *http.Request) {

    log.Println("Echoing back request made to " + request.URL.Path + " to client (" + request.RemoteAddr + ")")
    podIp := os.Getenv("POD_IP")
    podName := os.Getenv("POD_NAME")
    nodeName := os.Getenv("NODE_NAME")
    fmt.Fprintf(writer, "Pod %s is serving with IP %s on Node %s!\n", podName, podIp, nodeName)
}

func main() {

    log.Println("starting server, listening on port " + getServerPort())

    http.HandleFunc("/", EchoHandler)
    http.ListenAndServe(":" + getServerPort(), nil)
}

