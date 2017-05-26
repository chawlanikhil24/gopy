package goPy

import (
	"encoding/json"
	"fmt"
	"net"
)

//TCProute is the first struct in this program
type TCProute struct {
	Host string
	Port string
}

// the data structure to be used to make RPC call
type data struct {
	Method string        `json:"method"`
	Args   []interface{} `json:"args"`
	Time   string        `json:"time"`
}

// The Response of data from RPC CALL
type dataResponse struct {
	Response string `json:"response"`
}

// Using the simple connection interface from net package
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

// func main() {
// 	cli := &TCProute{"localhost", "9000"}
// 	conn, _ := cli.connect()
// 	t0 := time.Now()
// 	tIni := (strconv.Itoa(t0.Hour())) + ":" + (strconv.Itoa(t0.Minute())) + ":" + strconv.Itoa(t0.Second()) + "." + strconv.Itoa(t0.Nanosecond())[:6]
// 	d := &data{"show", []interface{}{10, 20.2, 30, 40, "i am nikhil"}, tIni}
// 	conn.sendData(d)
// 	recv, _ := conn.recvData()
// 	fmt.Println(recv)
//
// }
