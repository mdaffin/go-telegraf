package telegraf

import (
	"fmt"
	"net"
	"time"
)

type Measurement struct {
	name      string
	tagSet    map[string]string
	fieldSet  map[string]interface{}
	timestamp time.Time
}

type Client struct {
	conn   net.Conn
	tagSet map[string]string
}

func New(conn net.Conn) Client {
	return Client{
		conn: conn,
	}
}

func (c *Client) Write(m Measurement) error {
	_, err := fmt.Fprintf(c.conn, "%s\n", m)
	return err
}

func NewMeasurement(name string, tags map[string]string) Measurement {
	if tags == nil {
		tags = map[string]string{}
	}
	return Measurement{
		name:      name,
		tagSet:    tags,
		fieldSet:  map[string]interface{}{},
		timestamp: time.Now(),
	}
}

func (m *Measurement) AddBool(name string, value bool) {
	m.fieldSet[name] = value
}

func (m *Measurement) AddInt(name string, value int) {
	m.fieldSet[name] = value
}

func (m *Measurement) AddInt8(name string, value int8) {
	m.fieldSet[name] = value
}

func (m *Measurement) AddInt16(name string, value int16) {
	m.fieldSet[name] = value
}

func (m *Measurement) AddInt32(name string, value int32) {
	m.fieldSet[name] = value
}

func (m *Measurement) AddInt64(name string, value int64) {
	m.fieldSet[name] = value
}

func (m *Measurement) AddUInt(name string, value uint) {
	m.fieldSet[name] = value
}

func (m *Measurement) AddUInt8(name string, value uint8) {
	m.fieldSet[name] = value
}

func (m *Measurement) AddUInt16(name string, value uint16) {
	m.fieldSet[name] = value
}

func (m *Measurement) AddUInt32(name string, value uint32) {
	m.fieldSet[name] = value
}

func (m *Measurement) AddUInt64(name string, value uint64) {
	m.fieldSet[name] = value
}

func (m *Measurement) AddFloat32(name string, value float32) {
	m.fieldSet[name] = value
}

func (m *Measurement) AddFloat64(name string, value float64) {
	m.fieldSet[name] = value
}

func (m *Measurement) AddString(name string, value string) {
	m.fieldSet[name] = value
}
