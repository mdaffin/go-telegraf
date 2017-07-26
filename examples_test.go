package telegraf_test

import (
	"log"
	"net"

	"github.com/mdaffin/go-telegraf"
)

func ExampleMeasurement() {
	conn, err := net.Dial("udp", "127.0.0.1:8094")
	if err != nil {
		log.Fatalf("failed to connect: %s", err)
	}
	defer conn.Close()

	client := telegraf.New(conn)

	measurement := telegraf.NewMeasurement("cpu", nil)
	measurement.AddFloat64("load_avg", 1.4)
	measurement.AddInt("counter", 1)

	if err := client.Write(measurement); err != nil {
		log.Fatalf("failed to write metric: %s", err)
	}

	// Output:
}
