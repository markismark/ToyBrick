# ToyBrick--Free splicing API
Some programming languages only have one thread.For example,JavaScript and PHP.Sometimes, we need to call HTTP request more than one. For performance consideration, merge multi requests into one.

RequestProxy open multi goruntime to send the requests and avoid one by one

# ToyBrick--自由组合你的ＡＰＩ
一些程序语言是单线程，例如JavaScript，PHP．有时候，我们需要调用多个ＡＰＩ．为了性能考虑，将多个请求合并成一个，通过后端多线程的方式来请求处理，最后统一返回给调用方．

## FUNCTION
* Forwarding request
* Monitor
* WebSocket
* Refererr White List
* Request White List
* Request Pool
* PHP AND JS SDK

##HOW TO USE
1. Build and Run. You can use the command 

    > `go build -o toybrick main.go ; ./toybrick`
    
    or
     
    > `./run.sh`
    
2. The proxy url : http://localhost:8123/proxy

3. The example url : http://localhost:8123/static/js/example/main.html


> UNDER DEVELOPMENT




