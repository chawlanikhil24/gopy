package goPy

import (
	"encoding/json"
	"net"
)

// Connection  interface from net package
type Connection struct {
	Conn net.Conn
}

// Data structure to be used to make RPC call
type Data struct {
	Method string        `json:"method"`
	Args   []interface{} `json:"args"`
	// Time   string        `json:"time"`
}

// DataResponse of data from RPC CALL
type DataResponse struct {
	Response string `json:"response"`
}

// RandomJSON is used to send random data as string
// {"data":"USER INPUT STRING DATA"}
type RandomJSON struct {
	StringData interface{} `json:"data"`
}

//TCPRoute is used to set the Host and Port pair
type TCPRoute struct {
	Host string
	Port string
}

// Connect is used to connect with host and port on a TCP stream
func (obj *TCPRoute) Connect() (*Connection, error) {
	addr := obj.Host + ":" + obj.Port
	// fmt.Println("host and port to connect", addr)
	client, err := net.Dial("tcp", addr)
	ret := &Connection{client}
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// Close Function can be used to close the TCP stream
func (obj *Connection) Close() {
	conn := *obj
	x := conn.Conn
	x.Close()
}

// MakeRPC is used to send *Data  over TCP stream
func (obj *Connection) MakeRPC(dataIN *Data) (bool, error) {
	jsonData, err := json.Marshal(*dataIN)
	if err != nil {
		return false, err
	}
	obj.Conn.Write(jsonData)
	return true, nil
}

// SendRandomJSON method can be used to send any sort of StringData
func (obj *Connection) SendRandomJSON(dataIN *RandomJSON) error {
	jsonData, err := json.Marshal(*dataIN)
	if err != nil {
		return err
	}
	obj.Conn.Write(jsonData)
	return nil
}

// RecvData is used to RecvData data over TCP stream
func (obj *Connection) RecvData() (*DataResponse, error) {
	r := make([]byte, 4096)
	n, _ := obj.Conn.Read(r)
	response := r[:n]
	var v DataResponse
	err := json.Unmarshal(response, &v)
	if err != nil {
		return &v, err
	}
	return &v, nil
}
