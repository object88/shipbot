package client

import "github.com/object88/shipbot/log"

type Option func(c *Client) error

func SetLogger(l *log.Log) Option {
	return func(c *Client) error {
		c.Log = l
		return nil
	}
}
