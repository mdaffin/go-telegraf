package telegraf_test

import (
	"log"
	"time"

	"github.com/mdaffin/go-telegraf"
)

func ExampleClient() {
	client, err := telegraf.NewTCP("127.0.0.1:8094")
	if err != nil {
		log.Fatal("could not connect:", err)
	}
	defer func() { _ = client.Close() }()

	m := telegraf.MeasureFloat64("cpu", "load_avg", 0.5)

	if err := client.Write(m); err != nil {
		log.Fatal("failed to write metric:", err)
	}
}

func ExampleMeasurement() {
	_ = telegraf.MeasureInt("app", "request_size", 5042).AddString("error", "somethine went wrong").AddTag("path", "/api/testing")
}

func ExampleMeasurement_complex() {
	m := telegraf.MeasureInt("app", "request_size", 5042).AddTag("path", "/api/testing")
	start := time.Now()
	time.Sleep(time.Millisecond * 100)
	m = m.AddMillisecondsSince("request_time", start).AddInt("response_size", 3045).AddTag("status_code", "200")
}
