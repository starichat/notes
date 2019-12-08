# 「浅」谈认证授权

> 简介：我们乘坐火车时，一般需要先凭身份证买票取票才能持票乘车，为了防止逃票，列车员还会进行查票验票，对我们的身份进行验证。我们访问互联网服务的时候，也是类似的过程，凭用户名密码登录（买票取票），访问服务接口享受互联网服务（持票乘车）。
我们可以将凭用户名密码登录（买票取票）称为认证，经过认证后授权乘坐某趟列车的权限，我们就可以凭借这个权限去享受这趟旅途了。
所以，我们登录过程其实涉及到两个过程：
* 认证：
指的是当前用户的身份，当用户登陆过后，系统便能追踪到他的身份作出符合相应的业务逻辑的操作。即是用户没有登陆，大多数系统也能追踪到他的身份，只是当作来宾或者匿名用户来处理。认证技术解决的是“我是谁”的问题.
* 授权：
指的是什么样的身份被允许访问某些资源，在获取到用户身份后继续检查用户的权限。来解决用户能访问什么样的资源。

接下来，我们针对下面一个小的需求带你了解认证授权过程。

> 需求：为了保证接口调用的安全性，只有经过认证的系统才能调用我们的接口，没有认证的接口会被拒绝。

对于这样一个简单需求，其实我们需要处理两件事情：认证和授权。

> 因为http协议是无状态的，当我们需要获得用户是否已经登录时，我们需要检查用户的登录状态。一般来说登录成功后，服务端会颁发一个登录凭证，这个凭证一般会存在两个地方，客户端和服务端，当然也只有这两个地方能够存放。在每次请求的时候必须携带该凭证才能请求api以确保api的安全。就像我们乘车上车的时候，列车员会验票，防止我们逃票。
{补充http协议的无状态性}

微服务架构的安全访问图如下：
http://img.blog.itpub.net/blog/2019/08/21/8bba1b8558339304.jpeg?x-oss-process=style/bb

## 传统的认证授权措施
通过用户名和密码来做认证。我们允许访问我们服务的调用放，派发一个应用名和一个对应的密码。调用方每次进行接口与请求的时候，都携带自己的APPid和密码，服务端接受到接口调用请求之后，解析出用户名和密码，并与存储在服务端的appid和密码进行比对。如果一致，则说明认证成功，否则拒绝请求。

简单来说就是以下方式：



这样验证，有一个明显的缺点：每次都是明文传输密码，很容易被中间人截获，因此，我们在此基础上做一些改进，使用加密算法，对密码进行加密再传输到服务端进行验证。

「重放攻击」
{csrf攻击}

以上做法仍然会被中间人截获，黑客携带该密码仍然可以请求其他资源。

## jwt
1. 用户使用用户名和口令到认证服务器上请求认证
2. 认证服务器验证用户名和口令后，以服务器端生成JWT Token，这个token的生成过程如下
 * 认证服务器还会生成一个 Secret Key（密钥）
 * 对JWT Header和 JWT Payload分别求Base64。在Payload可能包括了用户的抽象ID和的过期时间。
 * 用密钥对JWT签名 HMAC-SHA256(SecertKey, Base64UrlEncode(JWT-Header)+'.'+Base64UrlEncode(JWT-Payload));
3. 然后把 base64(header).base64(payload).signature 作为 JWT token返回客户端。
4. 客户端使用JWT Token向应用服务器发送相关的请求。这个JWT Token就像一个临时用户权证一样。 

客户端每次携带该token进行请求即可，应用服务器每次会验证该token是否过期，是否可用。为了保证安全，这里会缩短token有效期，但是有效期太短又会影响用户体验，所以这里一般有两种解决方案：
1. 提供access_token 该token有效期较长，refresh_token 该token用来刷新access_token
2. 短期的后台token刷新策略，token有效期很短，当token过期后，必须用当前token去获取新的token。

下面给出一个关于JWT的认证授权示例：
```
package main

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct {
	Name     string `json:"username" form:"username" query: "username"`
	Password string `json:"password" form:"password" query:"password"`
}

func login(c echo.Context) error {
	u := new(User)
	// Throws unauthorized error
	if err := c.Bind(u); err != nil {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// // Set claims
	// claims := token.Claims.(jwt.MapClaims)
	// claims["name"] = "Jon Snow"
	// claims["admin"] = true
	// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func main() {
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Login route
	e.POST("/login", login)

	// Unauthenticated route
	e.GET("/", accessible)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", restricted)

	e.Logger.Fatal(e.Start(":8080"))
}

```

运行该服务：
使用postman请求接口：如果没有携带token或者token错误则返回。。，只有对正确的token才会给予结果。

我们看看jwt的token是如何生成以及，服务端如何检验客户端传递的token数据的。

## oauth策略
针对上文提出的传统的认证策略，我们发现仍然不安全，主要原因客户端传递的认证数据是静态的不变的，无法分辨请求方到底是谁。我们提出以下解决方案：
客户端认证的时候，认证成功后，服务端给客户端加密后的token，该token代表着用户认证信息，以后客户端持有该token去访问其他api，服务端只需要验证该token即可。
这种策略为oauth。

oauth提出了新的解决策略，调用方将请求接口的URL和APPID、密码拼接在一起，然后进行加密，生成一个token。调用方在进行接口请求的时候，将这个token及appid，随url一块传递给服务端。服务端接受到这些数据之后，根据appid从数据库中取出相应的密码，并通过同样的token生成算法，生成另一个token。用这个新生成的token和调用方传递来的token对比。如果一致，则允许接口调用请求，否则，拒绝接口调用。

具体流程如下：
1. 生成token
2. 生成新的url
3. 请求服务端
4. 服务端解析出url，appid，token
5. 从数据库中取出数据进行对比
6. 生成token
7. 比对服务端token和客户端token


上述这种设计由于token生成是固定的，所以仍然不安全，被截获后，仍然可以被认证。那么就对生成token的算法进行优化，生成token加入时间戳，调用方在进行接口请求的时候，将token，appid，时间戳，一起传给服务端。

服务端接受到这些数据之后，会验证时间戳，验证一定时间内的token是否过期。

流程如下：
1. 生成token
2. 生成url
3. 请求服务端
4. 服务端解析url，appid，token，ts
5. 验证token是否失效
6. 从数据库中取出数据验证
7. 生成服务端token
8. 校验服务端token和客户端token


## 应用
1. 一般服务的登录校验
jwt 足够
2. 外部系统访问微服务系统
设计api请求授权系统，为了减少对数据的查询，一般来说使用refresh_token 策略更为有效


 > 总结：安全是相对的，没有觉得的安全，我们能做的不过是增加破解的成本而已。当然同时也增加了自己的成本。上述安全措施已经在接口响应效率和整体token算法之间取得了很好的权衡。但是我们会发现每次请求都要验证数据库，这也是一项不小的开机，因此，对于一般的安全验证没有那么高的api，可以降低一些需求，采用jwt模式，由服务端颁发token，客户端携带该token去完成请求即可。

