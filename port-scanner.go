package main

import (
	"fmt"
	"net"
	"time"
)

var (
	ip           string
	fport, lport int
	done         = make(chan bool, 1)
)

func main() {
	fmt.Println("Enter an IP address and a port range to scan:")
	fmt.Println("IP address:")
	_, err := fmt.Scanf("%s", &ip)
	handleErr(err)

	fmt.Println("Initial port:")
	_, err = fmt.Scanf("%d", &fport)
	handleErr(err)

	fmt.Println("Last port:")
	_, err = fmt.Scanf("%d", &lport)
	handleErr(err)

	for port := fport; port <= lport; port++ {
		go makeReq(ip, port, lport)
	}

	<-done
	time.Sleep(2 * time.Millisecond)
}

func makeReq(ip string, port int, lastport int) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), time.Millisecond)
	if err != nil {
		// fmt.Printf("The port %d is closed\n", port)
	} else {
		fmt.Printf("The port %d is open\n", port)
		defer conn.Close()
	}

	if port == lastport {
		done <- true
	}
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
