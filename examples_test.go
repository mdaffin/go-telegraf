package telegraf_test

import (
	"log"

	"github.com/mdaffin/go-telegraf"
)

func ExampleMeasurement() {
	client, err := telegraf.NewUDP("127.0.0.1:8094", nil)
	if err != nil {
		log.Fatal("could not connect:", err)
	}
	defer client.Close()

	measurement := telegraf.NewMeasurement("cpu", map[string]string{"region": "europe-west1"})
	measurement.AddFloat64("load_avg", 1.4)
	measurement.AddInt("counter", 1)

	if err := client.Write(measurement); err != nil {
		log.Fatal("failed to write metric:", err)
	}

	// Output:
}
