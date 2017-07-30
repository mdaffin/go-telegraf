package telegraf

import (
	"fmt"
	"net"
)

type Client struct {
	conn net.Conn
}

func NewUDP(addr string) (Client, error) {
	conn, err := net.Dial("udp", addr)
	return Client{conn: conn}, err
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) Write(m Measurement) error {
	_, err := fmt.Fprintln(c.conn, m.ToLineProtocal())
	return err
}

func (c *Client) WriteAll(m []Measurement) error {
	// TODO write a benchmark to see if multiple writes or a string join and a single write is faster
	for _, m := range m {
		if err := c.Write(m); err != nil {
			return err
		}
	}
	return nil
}
