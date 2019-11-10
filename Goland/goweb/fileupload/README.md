# 基于原生go语言的http，实现一个简单的文件上传服务器

> 简介：缘起，工作原因不得不看很多源代码，于是引发思考。平时使用框架进行开发，思考为什么这样设计的机会很少，甚至，觉得学不完，因为一个框架又一个框架的层出不穷，但是细想之下，计算机的理论基础好像并未发生什么大变，但是为什么总对新技术产生恐惧感呢？因为我们对原理掌握的不扎实，这里尝试用最原生的web原理实现功能，试图深层次了解底层原理。单纯为了学而学，显得有些无聊，于是采用实现功能的方式进行开发。今天就结合实现一个文件上传服务器的功能，来学习goweb原理

## 模块划分
接口设计：
login
```
api/login?
```

1. 用vue.js实现两个页面
 * 登陆/注册
 * 文件上传功能

2. go后端实现
show me your code
```

```
## 运行效果

## 总结
基于go实现的web服务器主要有以下三个关键点需要解析：
1. http.HandleFunc
2. http.HandleFunc 的参数
3. http.ListenAndServe 启动服务器