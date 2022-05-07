# 代理模式

代理模式（Proxy Design Pattern）是在实际开发中经常被用到的一种设计模式。

代理模式是一种结构型设计模式， 让你能够提供对象的替代品或其占位符。 代理控制着对于原对象的访问， 并允许在将请求提交给对象前后进行一些处理。

代理模式建议新建一个与原服务对象接口相同的代理类，代理类接收到客户端请求后会创建实际的服务对象， 并将所有工作委派给它。

![image](https://user-images.githubusercontent.com/65383410/165267686-569d80d5-4b89-49ab-8f2b-5a84e7c7f2f9.png)

## 解析

代理模式在不改变原始类（或叫被代理类）代码的情况下，通过引入代理类来给原始类附加功能。

## 实现

```go
package proxy

// 服务接口：代理必须遵循该接口才能伪装成服务对象。
type server interface {
	handleRequest(string, string) (int, string)
}

// 代理(Proxy)类：包含一个指向服务对象的引用成员变量。
type nginx struct {
	application       *application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func newNginxServer() *nginx {
	return &nginx{
		application:       &application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *nginx) handleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	// 验证通过后，把请求发给真实服务应用
	return n.application.handleRequest(url, method)
}

func (n *nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}

// 服务(Service)：提供了一些实用的业务逻辑
type application struct {
}

func (a *application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "Ok"
	}

	if url == "/create/user" && method == "POST" {
		return 201, "User Created"
	}
	return 404, "Not Found"
}

```

## 用法

```go
nginxServer := newNginxServer()
appStatusURL := "/app/status"
createUserURL := "/create/user"

httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

httpCode, body = nginxServer.handleRequest(createUserURL, "POST")
fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

httpCode, body = nginxServer.handleRequest(createUserURL, "GET")
fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

Output：
Url: /app/status
HttpCode: 200
Body: Ok

Url: /app/status
HttpCode: 200
Body: Ok

Url: /app/status
HttpCode: 403
Body: Not Allowed

Url: /app/status
HttpCode: 201
Body: User Created

Url: /app/status
HttpCode: 404
Body: Not Ok
```

## 应用场景

##### 1.业务系统的非功能性需求开发

代理模式最常用的一个应用场景就是，在业务系统中开发一些非功能性需求。 比如：监控、统计、鉴权、限流、事务、幂等、日志。

将这些附加功能与业务功能解耦，放到代理类中统一处理，让业务人员只需要关注业务方面的开发。

##### 2.访问控制 （保护代理）

如果你只希望特定客户端使用服务对象。代理可仅在客户端凭据满足要求时将请求传递给服务对象。

##### 3.代理模式在RPC中的应用

实际上，RPC 框架可以看作一种代理模式。通过远程代理，将网络通信、数据编解码等细节隐藏起来。

客户端在使用 RPC 服务的时候，就像使用本地函数一样，无需了解跟服务器交互的细节。

RPC 服务的开发者也只需要开发业务逻辑，就像开发本地使用的函数一样，不需要关注跟客户端的交互细节。
