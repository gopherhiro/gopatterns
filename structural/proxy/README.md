# 代理模式
代理模式是一种结构型设计模式， 让你能够提供对象的替代品或其占位符。 代理控制着对于原对象的访问， 并允许在将请求提交给对象前后进行一些处理。

代理模式建议新建一个与原服务对象接口相同的代理类，代理类接收到客户端请求后会创建实际的服务对象， 并将所有工作委派给它。

![image](https://user-images.githubusercontent.com/65383410/165267686-569d80d5-4b89-49ab-8f2b-5a84e7c7f2f9.png)


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
- 延迟初始化 （虚拟代理）。 如果你有一个偶尔使用的重量级服务对象， 一直保持该对象运行会消耗系统资源时， 可使用代理模式。
    - 你无需在程序启动时就创建该对象， 可将对象的初始化延迟到真正有需要的时候。
- 访问控制 （保护代理）。 如果你只希望特定客户端使用服务对象。
    - 代理可仅在客户端凭据满足要求时将请求传递给服务对象。
- 记录日志请求 （日志记录代理）。 适用于当你需要保存对于服务对象的请求历史记录时。
    - 代理可以在向服务传递请求前进行记录。
- 缓存请求结果 （缓存代理）。 适用于需要缓存客户请求结果并对缓存生命周期进行管理时， 特别是当返回结果的体积非常大时。
    - 代理可对重复请求所需的相同结果进行缓存， 还可使用请求参数作为索引缓存的键值。
