package telegraf

import (
	"net"
	"time"
)

type Metric struct {
	measurement string
	tagSet      map[string]string
	fieldSet    map[string]interface{}
	timestamp   time.Time
}

type Client struct {
	conn net.Conn
}

func New(conn net.Conn) Client {
	return Client{
		conn: conn,
	}
}

func NewMetric(measurement string) Metric {
	return Metric{
		measurement: measurement,
		timestamp:   time.Now(),
	}

}
