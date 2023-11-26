package godb

import (
	"encoding/json"
	"net"
)

type Client struct {
	conn   net.Conn
	reader *json.Decoder
	writer *json.Encoder
}

func NewClient(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	client := &Client{
		conn:   conn,
		reader: json.NewDecoder(conn),
		writer: json.NewEncoder(conn),
	}

	return client, nil
}

func (c *Client) Get(key string) (*Response, error) {
	request := Request{
		Method: "get",
		Params: []string{key},
	}

	err := c.writer.Encode(request)
	if err != nil {
		return nil, err
	}

	response := Response{}
	err = c.reader.Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) Set(key string, value string) (*Response, error) {
	request := Request{
		Method: "set",
		Params: []string{key, value},
	}

	err := c.writer.Encode(request)
	if err != nil {
		return nil, err
	}

	response := Response{}
	err = c.reader.Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) Del(key string) (*Response, error) {
	request := Request{
		Method: "del",
		Params: []string{key},
	}

	err := c.writer.Encode(request)
	if err != nil {
		return nil, err
	}

	response := Response{}
	err = c.reader.Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
