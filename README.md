This is a starting point for Go solutions to the
["Build Your Own Redis" Challenge](https://codecrafters.io/challenges/redis).

Run:

* start server 
   ```bash  
   go run ./cmd/main 
   ```
* send command
   ```bash  
    telnet localhost 6379
    #eg
    SET 1 1 1000 # SET KEY VAL EXP
    GET 1 #GET KEY
    DEL 1 # DEL KEY
    PING 
   ```