// Writes metrics to telegraf using the socket_listener. It supports the udp, tcp and unix socket connections.
package telegraf

import (
	"fmt"
	"net"
)

// Client connects and writes measurements to telegraf.
type Client struct {
	conn net.Conn
}

// NewUDP client that connects to the telegraf socket_listener plugin with a udp address.
//
// Example telegraf configuration.
//
//   [[inputs.socket_listener]]
//     service_address = "udp://127.0.0.1:8094"
func NewUDP(addr string) (Client, error) {
	conn, err := net.Dial("udp", addr)
	return Client{conn: conn}, err
}

// NewTCP client that connects to the telegraf socket_listener plugin with a tcp address.
//
// Example telegraf configuration.
//
//   [[inputs.socket_listener]]
//     service_address = "tcp://127.0.0.1:8094"
func NewTCP(addr string) (Client, error) {
	conn, err := net.Dial("tcp", addr)
	return Client{conn: conn}, err
}

// NewUnix client that connects to the telegraf socket_listener plugin with a unix socket.
//
// Example telegraf configuration.
//
//   [[inputs.socket_listener]]
//     service_address = "unix:///var/run/telegraf.sock"
func NewUnix(addr string) (Client, error) {
	conn, err := net.Dial("unix", addr)
	return Client{conn: conn}, err
}

// Close the connection to telegraf.
func (c *Client) Close() error {
	return c.conn.Close()
}

// Write a metric to telegraf.
func (c *Client) Write(m Measurement) error {
	_, err := fmt.Fprintln(c.conn, m.ToLineProtocal())
	return err
}

// Write a list of metrics to telegraf.
func (c *Client) WriteAll(m []Measurement) error {
	// TODO write a benchmark to see if multiple writes or a string join and a single write is faster
	for _, m := range m {
		if err := c.Write(m); err != nil {
			return err
		}
	}
	return nil
}
