
# go-react-chatapp

# Prerequisites
node v20.5.1  
npm v9.8.0  
redis-stack-server

# Infrastructure
0. cd into infrastucture folder
1. Have aws cli and terraform installed
2. ```aws configure```
3. Enter Access key and Secret Access key
4. ```terraform init```
5. ```terraform plan```
6. ```apply```



# Backend
0. cd into backend/ folder
1. Have redis installed
2. ```go mod init``` 
3. ```go mod tidy```


## To start backend: 
Step 1: start redis server
```redis-stack-server```

Step 2: cd in to backend folder then run backend:  
```go run main.go```

# Cheatsheet shortcut:
Go into redis terminal and do:
```redis-cli```  
1. Check the Contents of the chatHistory List:  
```LRANGE chatHistory 0 -1```
2. Empty the chatHistory List: 
```DEL chatHistory```
3. dellte all keys:
```redis-cli KEYS * | xargs redis-cli DEL```

to kill port:
```kill -9 $(lsof -t -i :8080)```
