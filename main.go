package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"strings"

	"go.bug.st/serial"
)

type Rmcpacket = struct {
	Type     string
	Utc      string
	Status   string
	Lat      string
	LatDR    string
	Lon      string
	LonDR    string
	Speed    string
	Track    string
	Date     string
	Magvar   string
	Vardir   string
	Modeind  string
	Checksum string
}

func main() {
	pFlag := flag.Int("p", 0, "Port Number to read.")

	flag.Parse()

	// Retrieve the port list
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}

	var id int
	id = 0
	for _, port := range ports {
		fmt.Printf("Found port: %d %v\n", id, port)
		id++
	}

	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open(ports[*pFlag], mode)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(port)
	for scanner.Scan() {
		data := scanner.Text()

		switch {
		case strings.Contains(data, "$GNGSA"):
			// fmt.Printf("%s\n", data)
		case strings.Contains(data, "$GNGGA"):
			// fmt.Printf("Fields are: %q\n", strings.Split(data, ",\n"))
			// fmt.Printf("%s\n", data)
		case strings.Contains(data, "$GNRMC"):
			fmt.Printf("Fields are: %q\n", strings.Split(data, ","))
			fmt.Printf("%q %q %q\n", strings.Split(data, ",")[1][0:2], strings.Split(data, ",")[1][2:4], strings.Split(data, ",")[1][4:])
			// fmt.Printf("\n")
		case strings.Contains(data, "$GNGLL"):
			// fmt.Printf("\n")
		case strings.Contains(data, "$GNVTG"):
			// fmt.Printf("\n")
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
