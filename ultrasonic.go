package main

import (
	"fmt"
	"net"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

const (
	// Server address and port
	serverAddr = "192.168.1.100:8002"

	// GPIO pins (BCM numbering)
	triggerPin = 23
	echoPin    = 24

	// Speed of sound
	soundSpeed = 34300 // cm/s
)

func main() {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		return
	}
	// Unmap gpio memory when done
	defer rpio.Close()

	// Setup pins
	trigger := rpio.Pin(triggerPin)
	echo := rpio.Pin(echoPin)

	trigger.Output()
	echo.Input()

	// Connect to server
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	for {
		// Measure distance
		distance := measure(trigger, echo)

		// Send data to server
		fmt.Fprintf(conn, "%.1f\n", distance)

		// Wait before next measurement
		time.Sleep(500 * time.Millisecond)
	}
}

func measure(trigger, echo rpio.Pin) float64 {
	// Trigger pulse
	trigger.High()
	time.Sleep(10 * time.Microsecond)
	trigger.Low()

	// Measure pulse duration
	start := time.Now()
	for echo.Read() == rpio.Low {
	}
	risingEdge := time.Now()
	for echo.Read() == rpio.High {
	}
	pulseDuration := risingEdge.Sub(start)

	// Calculate distance
	return soundSpeed * pulseDuration.Seconds() / 2
}
