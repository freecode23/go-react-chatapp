
# go-react-chatapp

# Prerequisites
node v20.5.1  
npm v9.8.0  


# Backend
redis installed
go get -u github.com/gorilla/mux
go get -u github.com/gorilla/websocket
go get github.com/rs/cors
go mod init 
go mod tidy

## To start backend: 
Step 1: start redis server
```redis-server```

Step 2: cd in to backend folder then run backend:  
```go run main.go```

### redis cheatsheet:
Go into redis terminal and do:
```redis-cli```  
1. Check the Contents of the chatHistory List:  
```LRANGE chatHistory 0 -1```
2. Empty the chatHistory List: 
```DEL chatHistory```

to kill port:
```kill -9 $(lsof -t -i :8080)```
