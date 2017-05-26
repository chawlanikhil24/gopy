package goPy

import (
	"encoding/json"
	"fmt"
	"net"
)

type TCProute struct {
	Host string
	Port string
}

type data struct {
	Method string        `json:"method"`
	Args   []interface{} `json:"args"`
	Time   string        `json:"time"`
}

type dataResponse struct {
	Response string `json:"response"`
}

type connection struct {
	conn net.Conn
}

func (obj *TCProute) connect() (*connection, error) {
	addr := obj.Host + ":" + obj.Port
	fmt.Println("host and port to connect", addr)
	client, err := net.Dial("tcp", addr)
	ret := connection{client}
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (obj *connection) sendData(dataIN *data) (bool, error) {
	jsonData, err := json.Marshal(*dataIN)
	if err != nil {
		return false, err
	}
	obj.conn.Write(jsonData)
	return true, nil
}

func (obj *connection) recvData() (dataResponse, error) {
	r := make([]byte, 4096)
	n, _ := obj.conn.Read(r)
	response := r[:n]
	var v dataResponse
	err := json.Unmarshal(response, &v)
	if err != nil {
		return v, err
	}
	return v, nil
}
