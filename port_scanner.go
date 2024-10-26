package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

func main() {

	var (
		ports      = flag.Int("p", 0, "Ports to be scanned")
		goRoutines = flag.Int("t", 1, "GoRoutines")
		ip_address = flag.String("ip-address", "127.0.0.1", "ip_address")
		b_port     = flag.Bool("a", false, "Scan all 65535 ports (when this flag is used the value of -p flag is ignored)")
		count      = 1
		wg         sync.WaitGroup
	)

	flag.Parse()

	if *ip_address == "" {
		log.Fatal("Plaase provide --ip-address")
	}

	if !*b_port {
		if *ports == 0 {

			fmt.Println("Scanning the first 1024 ports (use ./port_scanner -h for help)")

			*ports = 1024

		} else if *ports > 65535 {
			fmt.Println("Number of ports over 65535 exit code 1")
			return
		}
	} else {
		fmt.Println("Scanning all 65535 ports")
		*ports = 65535
	}

	if *goRoutines > *ports {
		log.Fatal("You Can not Have more go routines than ports to scan.")
	}

	portsPerRoutine := *ports / *goRoutines
	restOfPorts := *ports % *goRoutines

	for i := 0; i < *goRoutines; i++ {
		wg.Add(1)
		go func(count int) {
			defer wg.Done()
			for j := count; j < portsPerRoutine+count; j++ {
				result, err := net.DialTimeout("tcp", net.JoinHostPort(*ip_address, strconv.Itoa(j)), time.Duration(1*time.Second))

				if err == nil {
					fmt.Println("[+] Port", j, "open")
					result.Close()
				}
			}
		}(count)

		count += portsPerRoutine
	}

	for i := count; i < count+restOfPorts; i++ {
		result, err := net.DialTimeout("tcp", net.JoinHostPort(*ip_address, strconv.Itoa(i)), time.Duration(1*time.Second))
		if err == nil {
			fmt.Println("[+] Port", i, "open")
			result.Close()
		}
	}

	count += restOfPorts - 1

	wg.Wait()
	fmt.Println(restOfPorts)
	fmt.Println("Done closed or filtered Ports:", count)
}
