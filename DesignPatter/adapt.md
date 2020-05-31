## 适配器模式
适配器模式，即定义一个包装类，用于包装不兼容接口的对象

包装类 = 适配器Adapter；
被包装对象 = 适配者Adaptee = 被适配的类

把一个类的接口变换成客户端所期待的另一种接口，从而使原本接口不匹配而无法一起工作的两个类能够在一起工作。
![avatar]https://blog.csdn.net/qibin0506/article/details/50598359
用 go 语言实现一个适配器模式可以如下：
```
package adaptor

import "fmt"

// 我们的接口（新接口）——音乐播放
type MusicPlayer interface {
    play(fileType string, fileName string)
}

// 在网上找的已实现好的库 音乐播放
// ( 旧接口）
type ExistPlayer struct {
}

func (*ExistPlayer) playMp3(fileName string) {
    fmt.Println("play mp3 :", fileName)
}
func (*ExistPlayer) playWma(fileName string) {
    fmt.Println("play wma :", fileName)
}

// 适配器
type PlayerAdaptor struct {
    // 持有一个旧接口
    existPlayer ExistPlayer
}

// 实现新接口
func (player *PlayerAdaptor) play(fileType string, fileName string) {
    switch fileType {
    case "mp3":
        player.existPlayer.playMp3(fileName)
    case "wma":
        player.existPlayer.playWma(fileName)
    default:
        fmt.Println("暂时不支持此类型文件播放")
    }
}


```

## 命令模式
命令模式(Command)，将一个请求封装为一个对象，从而使你可用不同的请求对客户进行参数化；对请求排队或记录请求日志，以及支持可撤销的操作

命令模式可以将发送者和接收者完全解耦，发送者与接收者之间没有直接引用关系，发送请求的对象只需要知道如何发送请求，而不必知道如何完成请求。
![avatar]https://blog.csdn.net/cloudUncle/article/details/83685495

用 Go 语言实现一下：
首先让所有的命令对象实现相同的包含一个方法的接口：

```
type command interface {
    execute()
}
                                                                                                    
//开灯命令
type lightOnCommand struct {
    mLight *light //命令对象包含的特定接收者
}
                                                                                                   
//返回一个开灯命令的实例对象
func NewLightOnCommand(light *light) command {
    return &lightOnCommand{mLight: light}
}
                                                                                                   
//实现接口方法捆绑接收者的动作
func (this *lightOnCommand) execute() {
    if !this.mLight.isOn() {
        this.mLight.setOn(true) //开灯
    }
}
                                                                                                   
//关灯命令
type lightOffCommand struct {
    mLight *light
}
          
func NewLightOffCommand(light *light) command {
    return &lightOffCommand{mLight: light}
}
                                                                                                   
func (this *lightOffCommand) execute() {
    if this.mLight.isOn() {
        this.mLight.setOn(false) //关灯
    }
}
```

我们应当考虑面向接口编程，大部分接收者都有简单的开关命令，故上述的代码可改为：
```
type receiver interface {
    setOn(bool) //true：开/false：关
    isOn() bool
}
                                 
//打开命令                                                 
type onCommand struct {
    receiver Receiver
}
                                                         
//创建打开命令的实例，为该实例捆绑接收者                          
func NewOnCommand(receiver Receiver) command {
    return &onCommand{receiver}
}
                                         
//被封装的“请求”                                          
func (this *onCommand) execute() {
    if !this.receiver.isOn() {
        this.receiver.setOn(true) //打开
    }
}
                                           
//关闭命令                                           
type offCommand struct {
    receiver Receiver
}
                                                                                      
func NewOffCommand(receiver Receiver) command {
    return &offCommand{receiver}
}
                                                                                      
func (this *offCommand) execute() {
    if !this.receiver.isOn() {
        this.receiver.setOn(false) //关闭
    }
}

```
最后，再来看看客户端的代码：
```
type RemoteController struct {
    slot command
}
                                                                 
func (this *RemoteController) SetCommand(command command) {
    this.slot = command
}
                                                                 
func (this *RemoteController) ButtonPressed() {
    if this.slot == nil {
        panic("Do not assign command to Controller's slot!")
    }
    this.slot.execute()
}
```

看看接收者们：
```
const (
    LIGHT = " light"
    DOOR  = " door"
)
                                  
//接收者接口        
type Receiver interface {
    setOn(bool)
    isOn() bool
}
                                             
type light struct {
    name string
    on   bool
}
                                             
func (this *light) setOn(b bool) {
    if b {
        fmt.Println(this.name + LIGHT + " is on.")
    } else {
        fmt.Println(this.name + LIGHT + " is off.")
    }
    this.on = b
}
                                             
func (this *light) isOn() bool {
    return this.on
}
                                             
func NewRoomLight() Receiver {
    return &light{"Room", false}
}
                                             
func NewTableLampLight() Receiver {
    return &light{"Table Lamp", false}
}
                                             
type door struct {
    name string
    on   bool
}
                                             
func (this *door) setOn(b bool) {
    if b {
        fmt.Println(this.name + DOOR + " is opened.")
    } else {
        fmt.Println(this.name + DOOR + " is closed.")
    }
    this.on = b
}
                                             
func (this *door) isOn() bool {
    return this.on
}
                                             
func NewGarageDoor() Receiver {
    return &door{"Garage", false}
}
                                             
func NewKitchenDoor() Receiver {
    return &door{"Kitchen", false}
}
```

来测试下吧~：
```
func main() {
        ctrl := new(command.RemoteController)
    var roomLight, garageDoor command.Receiver
              
    roomLight = command.NewRoomLight()
    garageDoor = command.NewGarageDoor()
              
    cmdOn := command.NewOnCommand(roomLight)
    cmdOff := command.NewOffCommand(garageDoor)
              
    ctrl.SetCommand(cmdOn)
    ctrl.ButtonPressed()
    ctrl.SetCommand(cmdOff)
    ctrl.ButtonPressed()
}
```
