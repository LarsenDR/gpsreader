package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"unicode"

	"go.bug.st/serial"
)

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

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	buff := make([]byte, 100)
	for {
		// Reads up to 100 bytes
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
		}

		data := string(buff[:n])
		//key := string(buff[:6])

		switch {
		case strings.Contains(data, "$GNGGA"):
			fmt.Printf("Fields are: %q\n", strings.FieldsFunc(data, f))
			//fmt.Printf("%s", data)
			// case strings.Contains(data, "$GNRMC"):
			// 	fmt.Printf("%s", data)
			// case strings.Contains(data, "$GNGLL"):
			// 	fmt.Printf("%s", data)
		}

	}
}
