package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"io/ioutil"
)

// Function to read the port from a configuration file
func getPort(defaultPort string) string {
	configFile := "webserv.conf"
	if _, err := os.Stat(configFile); err == nil {
		// Read the config file
		data, err := ioutil.ReadFile(configFile)
		if err == nil {
			port := strings.TrimSpace(string(data))
			if port != "" {
				return port
			}
		}
	}
	return defaultPort
}

// Function to save the selected port to the configuration file
func savePort(port string) {
	configFile := "webserv.conf"
	err := ioutil.WriteFile(configFile, []byte(port), 0644)
	if err != nil {
		fmt.Println("Error saving port to config file:", err)
	}
}

// Get the server's IP address
func getIPAddress() string {
	// Connect to a non-routable IP to determine local IP
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "127.0.0.1" // Fallback to localhost
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

func main() {
	// Define a flag for the port with a default value
	defaultPort := "8080"
	portPtr := flag.String("port", getPort(defaultPort), "Port to serve on")

	// Parse the flags
	flag.Parse()

	// Save the selected port to a config file for future runs
	savePort(*portPtr)

	server := &http.Server{Addr: ":" + *portPtr}

	// Handle signals for graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		fmt.Println("\nShutting down server...")
		server.Close()
		os.Exit(0) // Exit gracefully
	}()

	// Display the server IP and port
	ipAddress := getIPAddress()
	fmt.Printf("Serving HTTP on %s port %s (http://%s:%s/)\n", ipAddress, *portPtr, ipAddress, *portPtr)

	// Display the additional message right after starting the server
	fmt.Println("\nRun webserv with -port <port> to select any port.")
	fmt.Println("(Close with CTRL+C)")

	// Start the server
	err := http.ListenAndServe(":"+*portPtr, http.FileServer(http.Dir(".")))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
