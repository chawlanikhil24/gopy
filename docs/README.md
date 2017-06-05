#**Description of "goPy" JSON-RPC Client Library**
---

This description involves the detailed information about the "Structures" and "Methods" provided in this library.

---

  ## Structures:
  * Connection: Connection is used to store the object of "net.Conn" from "net" libary of Golang. The "net.Conn" object is created as the result of "net.Dial".
  Example:
  ```
    client, err := net.Dial("tcp", addr)
    ret := &Connection{client}
  ```

  * Data: is a user-defined Data structure used to store method name to be invoked and its parameters, which futher used to make the RPC-CALL to Python Server. "Data" stores the "method name"(string) to be called and "method parameters"([]interface{}). Without this, one can't make the request to python Server.
  Example:
  ```
    data := Data{"method name ",[]interface{}{"parameter1","parameter2"}}
    response := MakeRPC(data)
  ```

  * RandomJSON: is a user-defined JSON data structure, which has a "key" named "data" and stores its "value" in "interface{}" when initialised. I made it for the developmental testing. Not any specific purpose.
  ```
    json := RandomJSON{"I am a key"}
    // can be assumed as :
    // json = {"data":"I am a key"}
  ```

* DataResponse: is used to store the JSON response coming from Python server. It has key named "response" whose value is the corresponding "result" of the RPC Call.
```
  r := MakeRPC(data)
  //r = {"response":data}
  fmt.Println(r.Response)
```

* TCPRoute: is used to store the "Host" name and "Port" number under one struct. Both of the members of the struct are of string data type.
```
  addr := TCPRoute{"localhost","9001"}
```
* Parameters: is used to provide a compact structure pass arguments by appending them in an "interface{}" array, which makes it flexible enough to pass any type of arguments, like, string,float, int in one array.
```
  Args := goPy.Parameters{"para1",22,234.5}

```

---
## Methods:
* **Connect():** Connect() function is a member function of "TCPRoute struct" and used to establish a TCP Connection across the "Host" and "Port" Stored by TCPRoute object. Connect() function after successful connection to Python Server returns a *Connection Object, else error. This *Connection is used for futher actions.

  *Function Prototype:*

  ```func (obj *TCPRoute) Connect() (*Connection, error)```
* **Close():** Close() function is a member function of "Connection struct" and used to close the TCP Port after use. It is necessary to Close the Port after use to save resources and uninterrupted enhanced performance.

  *Function Prototype*

  ```func (obj *Connection) Close()```

* **MakeRPC():** MakeRPC() function is a member function of "Connection struct", which accepts a "*Data" object pointer as a parameter and makes the RPC call to Python Server. If the call is successful, then it returns the "bool" response "true", else "false" and "error".

  *Function Prototype*

  ```func (obj *Connection) MakeRPC(dataIN *Data)(bool,error)```

* **SendRandomJSON():** SendRandomJSON() function is a member function of "Connection struct", which I developed for testing and only sends random test data over TCP. It accepts *RandomJSON object as parameter and returns error, if there is any.

  *Function Prototype*

  ```func (obj *Connection) SendRandomJSON(dataIN *RandomJSON)error```

* **RecvData():** RecvData() function is a member function of "Connection struct", which is used to Receive incoming data from Python Server, after making call either from "MakeRPC()" or "SendRandomJSON()". It returns a JSON of "*DataResponse struct".

  *Function Prototype*

  ```func (obj *Connection) RecvData() (*DataResponse,err)```
