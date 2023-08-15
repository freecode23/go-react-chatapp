
# go-react-chatapp

# PREREQUISITES
node v20.5.1  
npm v9.8.0  


# BACKEND
redis installed
go get -u github.com/gorilla/mux
go get -u github.com/gorilla/websocket
go mod init 
go mod tidy

## To start backend: 
Step 1: start redis server
```redis-server```

Step 2: cd in to backend folder then run backend:  
```go run main.go```