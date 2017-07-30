package telegraf_test

import (
	"log"
	"time"

	"github.com/mdaffin/go-telegraf"
)

func ExampleMeasurement() {
	client, err := telegraf.NewUDP("127.0.0.1:8094")
	if err != nil {
		log.Fatal("could not connect:", err)
	}
	defer client.Close()

	measurement := telegraf.MeasureInt("app", "request_size", 5042).AddTag("path", "/api/testing")
	start := time.Now()
	time.Sleep(time.Millisecond * 100)
	measurement = measurement.AddMillisecondsSince("request_time", start).AddInt("response_size", 3045)

	if err := client.Write(measurement); err != nil {
		log.Fatal("failed to write metric:", err)
	}

	// Output:
}
