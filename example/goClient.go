package main

// Step1: Import the package
import (
	"fmt"

	"github.com/chawlanikhil24/gopy"
)

func main() {
	// Step2: Initialise the go client for the Host and Port of running Python Server
	client := &goPy.TCPRoute{"localhost", "9001"}

	// Step3: Connect to Python Server
	conn, _ := client.Connect()

	// Step4: Initialise the Data which mentions the method name and required parameters
	data := &goPy.Data{"ADD", []interface{}{21, 22}}

	// Step5: Make the RPC call to Python server to process the input and return an output
	conn.MakeRPC(data)

	// Step6: Receive the ouput of the function call
	recvData, _ := conn.RecvData()

	// Step7: Finally Print the Receive data for now
	fmt.Println(recvData)
}
