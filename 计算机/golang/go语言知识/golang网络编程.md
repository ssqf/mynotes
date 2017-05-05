# golang网络编程

## tcp/udp连接

### 监听连接端口
1. 创建监听
    listener, err := net.Listen("tcp", "localhost:50000")
2. 接受连接
    conn, err := listener.Accept()
### 连接端口
1. 发起连接
    conn, err := net.Dial("tcp", "localhost:50000")
### 收发数据
1. 发送数据
    `_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))`
2. 接受数据
    `_, err := conn.Read(buf)`
### 关闭连接
`conn.Close()`


## web服务器建立

### 监听接口
    http.ListenAndServe("localhost:8080", nil)
### 请求处理
```
    func HFunc(w http.ResponseWriter, req *http.Request) {...} 
    type HandlerFunc func(ResponseWriter, *Request)
    http.HandleFunc("/", HelloServer)  //helloServer为HandlerFucn类型的回调函数，用于处理对"/"路径的请求
```
