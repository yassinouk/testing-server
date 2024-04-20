## this a server for testing ultrasonic sensor, it will return the distance in cm, and connect to the server using websocket.

### The code establishes a TCP connection to a server at a specified address ("192.168.1.100:8002") and transmits the measured distance data at regular intervals.
## How to use
1. Install the required binaries (specfically for raspberry pi 3 B+ running linux)
```bash
wget https://github.com/yassinouk/testing-server/releases/download/v0.1.0-alpha/testing-server.zip
unzip testing-server.zip
```
2. Run the server
```bash
./testing-server
```



