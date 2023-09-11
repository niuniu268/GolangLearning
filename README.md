# Instant Message System
## Golang Practice

### Build up a server
- go build -o server main.go ser.go
- nc 127.0.0.1 8888

### Build a user function

the user function is accomplished: whe a user becomes online, others would receive online message
![broadcast](https://github.com/niuniu268/GolangLearning/blob/master/img/Broadcast.png?raw=true)

when one user sends a message, the message should be broadcasted into all online users.
![conversation](https://github.com/niuniu268/GolangLearning/blob/master/img/conversation.png?raw=true)

