package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {

	count := 0

	ip_address := ""

	var ports = flag.Int("p", 0, "Ports to be scanned")
	var b_port = flag.Bool("a", false, "Scan all 65535 ports (when this flag is used the value of -p flag is ignored)")

	flag.Parse()

	if !*b_port {
		if *ports == 0 {

			fmt.Println("Scanning the first 1024 ports (use ./port_scanner -h for help)")

			*ports = 1023

		} else if *ports > 65535 {
			fmt.Println("Number of ports over 65535 exit code 1")
			return
		}
	} else {
		fmt.Println("Scanning all ports (./post_scanner -h for help): ")
		*ports = 65535
	}

	fmt.Println("Give IP address: ")
	fmt.Scan(&ip_address)
	fmt.Println()

	for i := 1; i <= *ports; i++ {

		result, err := net.DialTimeout("tcp", net.JoinHostPort(ip_address, strconv.Itoa(i)), time.Duration(1*time.Second))

		if err == nil {
			fmt.Println("[+] Port", i, "open")
			result.Close()
		} else {
			count++
		}
	}

	fmt.Println("Done closed or filtered Ports:", count)
}
