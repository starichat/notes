# 工厂模式
> 在面向对象的编程语言中（如java，C++）设计模式的概念广为人知, 应用的也非常广泛。设计模式让我们的代码变得灵活起来，具有很强的扩展性。但在与C语言比肩的Go语言中，设计模式的概念并没有十分突出，甚至很少听到。在Go的开发中，借鉴design pattern的理念同样回味无穷我们的开发带来极大的便利。

## 简单工厂
简单工厂模式的工厂类一般是使用静态方法，通过接收的参数的不同来返回不同的对象实例。

## 工厂方法
工厂方法是针对每一种产品提供一个工厂类。通过不同的工厂实例来创建不同的产品实例。
在同一等级结构中，支持增加任意产品。

## 抽象工厂
抽象工厂是应对产品族概念的。比如说，每个汽车公司可能要同时生产轿车，货车，客车，那么每一个工厂都要有创建轿车，货车和客车的方法。

下面结合一个例子看看抽象工厂如何应用的。

```
package main

import "fmt"

type GirlFriend struct {
    nationality string
    eyesColor   string
    language    string
}

type AbstractFactory interface {
    CreateMyLove() GirlFriend
}

type IndianGirlFriendFactory struct {
}

type KoreanGirlFriendFactory struct {
}

func (a IndianGirlFriendFactory) CreateMyLove() GirlFriend {
    return GirlFriend{"Indian", "Black", "Hindi"}
}

func (a KoreanGirlFriendFactory) CreateMyLove() GirlFriend {
    return GirlFriend{"Korean", "Brown", "Korean"}
}

func getGirlFriend(typeGf string) GirlFriend {

    var gffact AbstractFactory
    switch typeGf {
    case "Indian":
        gffact = IndianGirlFriendFactory{}
        return gffact.CreateMyLove()
    case "Korean":
        gffact = KoreanGirlFriendFactory{}
        return gffact.CreateMyLove()
    }
    return GirlFriend{}
}

func main() {

    a := getGirlFriend("Indian")

    fmt.Println(a.eyesColor)
}
```