# **goPy**
---
#### A Go Client library for JSON-RPC based on Python. With this library, you can cross-communicate with Python RPC Server in a flick.

#### For a quick demo for this library, please visit this [LINK](http://chawlanikhil24.blogspot.com/2017/06/lets-combine-features-of-python-and.html).

---
**Why I developed this library?**

Since, I've working on Python for 2 years and many of my current projects are based on Python. Now, since just for variety and enthusiasm I wanted to adapt myself to a compiler based language and so I heard about [Golang](chawlanikhil24.blogspot.com). It was new and I kinda felt fun doing it. Now, I wanted to merge my Python projects with Golang. And as said,


### *"Necessity is the mother of invention(innovation)"*

And here I developed a library for those who are keen to experiment with both **Python** & **Golang** .

---

**How to use this library?**

As, this is a Client Library, that means there has to be server first as well.Here's the [LINK](https://github.com/chawlanikhil24/goPyServer/blob/master/README.md) to setup the Python JSON-RPC Server using the library named [goPyServer](https://github.com/chawlanikhil24/goPyServer).

Now, I assume that you are all set with a running Python JSON-RPC Server.

* **Step 1:** Pull this library in your machine.
```
go get github.com/chawlanikhil24/gopy
```

* **Step 2:** import this library in your go script
```
import "github.com/chawlanikhil24/gopy"
```

* **Step 3:** Initialise the go client for the Host and Port of running Python Server
```
client := &goPy.TCPRoute{"localhost", "9001"}
```

* **Step 4:** Connect to Python Server
```
conn, _ := client.Connect()
```

* **Step 5:** Initialise the **goPy.Data** which mentions the **METHOD NAME** and **PARAMETERS REQUIRED**
```
data := &goPy.Data{"ADD", goPy.Parameters{21, 22}}
```

* **Step 6:** Make the RPC call to Python server to process the input and return an output
```
conn.MakeRPC(data)
```

* **Step 7:** Receive the ouput of the function call
```
recvData, _ := conn.RecvData()
```

* **Step 8:** Finally Print the Receive data for now
```
fmt.Println(recvData)
```

For the Detailed description of this library, click [Here](https://github.com/chawlanikhil24/gopy/tree/master/docs)
