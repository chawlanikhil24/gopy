# Step1: Lets Import the Library
import goPyServer as GP

# Step2: Create a sample "ADD" which returns the sum of 2 numbers passed in parameters
def ADD(arg1,arg2):
    return arg1+arg2

# Step3: Initialise the object of goPyServer
obj = GP.pyServ(TCP_IP="127.0.0.1",TCP_PORT=9001,Listen=4,buff=1024)

# Step4: Register the ADD method to the RPC Server
obj.register_function(ADD)

# Step5: connect to the Server
obj.connect()

# Step6: Finally Run the JSON RPCServer
print "Server Running on \"127.0.0.1:9001 ...\""
obj.RPCServer()
