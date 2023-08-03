[写给后续参加字节跳动青训营的同学](https://github.com/ACking-you/byte_douyin_project/issues/10)
# 抖音极简版
<!-- PROJECT SHIELDS -->

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- links -->
[your-project-path]:ACking-you/byte_douyin_project
[contributors-shield]: https://img.shields.io/github/contributors/ACking-you/byte_douyin_project.svg?style=flat-square
[contributors-url]: https://github.com/ACking-you/byte_douyin_project/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/ACking-you/byte_douyin_project.svg?style=flat-square
[forks-url]: https://github.com/ACking-you/byte_douyin_project/network/members
[stars-shield]: https://img.shields.io/github/stars/ACking-you/byte_douyin_project.svg?style=flat-square
[stars-url]: https://github.com/ACking-you/byte_douyin_project/stargazers
[issues-shield]: https://img.shields.io/github/issues/ACking-you/byte_douyin_project.svg?style=flat-square
[issues-url]: https://img.shields.io/github/issues/ACking-you/byte_douyin_project.svg
[license-shield]: https://img.shields.io/github/license/ACking-you/byte_douyin_project?style=flat-square
[license-url]: https://github.com/ACking-you/byte_douyin_project/blob/master/LICENSE



* [数据库说明](#数据库说明)
    * [数据库关系说明](#数据库关系说明)
    * [数据库建立说明](#数据库建立说明)
* [架构说明](#架构说明)
    * [各模块代码详细说明](#各模块代码详细说明)
        * [Handlers](#handlers)
        * [Service](#service)
        * [Models](#models)
* [遇到的问题及对应解决方案](#遇到的问题及对应解决方案)
    * [返回json数据的完整性和前端要求的一致性](#返回json数据的完整性和前端要求的一致性)
    * [is\_favorite和is\_follow字段的更新](#is_favorite和is_follow字段的更新)
    * [视频的保存和封面的切片](#视频的保存和封面的切片)
        * [视频的保存](#视频的保存)
        * [封面的截取](#封面的截取)
* [可改进的地方](#可改进的地方)
* [项目运行](#项目运行)

## 数据库说明


![database.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/892fbbe46695467ebe4fb4a12ebd412e~tplv-k3u1fbpfcp-watermark.image?)

> 单纯看上面的图会感觉很混乱，现在我们来将关系拆解。

### 数据库关系说明

**关系图如下：**


![database_relation.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/f08918db1ea84126bc21d23fe9401a75~tplv-k3u1fbpfcp-watermark.image?)

> 所有的表都有自己的id主键为唯一的标识。

user_logins：存下用户的用户名和密码

user_infos：存下用户的基本信息

videos：存下视频的基本信息

comment：存下每个评论的基本信息

**具体的关系索引：**

所有的一对一和一对多关系，只需要在一个表中加入对方的id索引。

* 比如user_infos和user_logins的一对一关系，在user_logins中加入user_id字段设为外键存储user_infos中对应的行的id信息。
* 比如user_infos和和videos的一对多关系，在videos中加入user_id字段设为外键存储user_infos中对应的行的id信息。

所有的多对多关系，需要多建立一张表，用该表作为媒介存储两个对象的id作为关系的产生，而它们各自表中不需要再存下额外的字段。

* 比如user_infos和videos的多对多关系，创建一张user_favor_videos中间表，然后将该表的字段均设为外键，分别存下user_infos和videos对应行的id。如id为1的用户对id为2的视频点了个赞，那么就把这个1和2存入中间表user_favor_videos即可。

### 数据库建立说明

数据库各表的建立无需自己实现额外的建表操作，一切都由gorm框架自动建表，具体逻辑在models层的代码中。

> gorm官方文档链接：[链接](https://gorm.io/zh_CN/docs/index.html)

建表和初始化操作由init_db.go来执行：

```go
func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.DBConnectString()), &gorm.Config{
		PrepareStmt:            true, //缓存预编译命令
		SkipDefaultTransaction: true, //禁用默认事务操作
		//Logger:                 logger.Default.LogMode(logger.Info), //打印sql语句
	})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&UserInfo{}, &Video{}, &Comment{}, &UserLogin{})
	if err != nil {
		panic(err)
	}
}
```

## 架构说明


![architecture.png](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/ae11d82b8de74787a258ef36f4cf2557~tplv-k3u1fbpfcp-watermark.image?)

> 以用户登录为例共需要经过以下过程：

1. 进入中间件SHAMiddleWare内的函数逻辑，得到password明文加密后再设置password。具体需要调用gin.Context的Set方法设置password。随后调用next()方法继续下层路由。
2. 进入UserLoginHandler函数逻辑，获取username，并调用gin.Context的Get方法得到中间件设置的password。再调用service层的QueryUserLogin函数。
3. 进入QueryUserLogin函数逻辑，执行三个过程：checkNum，prepareData，packData。也就是检查参数、准备数据、打包数据，准备数据的过程中会调用models层的UserLoginDAO。
4. 进入UserLoginDAO的逻辑，执行最终的数据库请求过程，返回给上层。

### 各模块代码详细说明

我开发的过程中是以单个函数为单个文件进行开发，所以代码会比较长，故我根据数据库内的模型对函数文件进行了如下分包：


![handlers.png](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/6dc222793d6f4038b1bf2435053bfee4~tplv-k3u1fbpfcp-watermark.image?)

service层的分包也是一样的。

#### Handlers

对于handlers层级的所有函数实现有如下规范：

所有的逻辑由代理对象进行，完成以下两个逻辑

1. 解析得到参数。
2. 开始调用下层逻辑。

例如一个关注动作触发的逻辑：

```go
NewProxyPostFollowAction().Do()
//其中Do主要包含以下两个逻辑，对应两个方法
p.parseNum() //解析参数
p.startAction() //开始调用下层逻辑
```

#### Service

对于service层级的函数实现由如下规范：

同样由一个代理对象进行，完成以下三个或两个逻辑

当上层需要返回数据信息，则进行三个逻辑：

1. 检查参数。
2. 准备数据。
3. 打包数据。

当上层不需要返回数据信息，则进行两个逻辑：

1. 检查参数。
2. 执行上层指定的动作。

例如关注动作在service层的逻辑属于第二类：

```go
NewPostFollowActionFlow(...).Do()
//Do中包含以下两个逻辑
p.checkNum() //检查参数
p.publish() //执行动作
```

#### Models

对于models层的各个操作，没有像service和handler层针对前端发来的请求就行对应的处理，models层是面向于数据库的增删改查，不需要考虑和上层的交互。

而service层根据上层的需要来调用models层的不同代码请求数据库内的内容。

## 遇到的问题及对应解决方案

### 返回json数据的完整性和前端要求的一致性

由于数据库内的一对一、一对多、多对多关系是根据id进行映射，所以models层请求得到的字段并不包含前端所需要的直接数据，比如前端要求Comment结构中需要包含UserInfo，而我的Comment结构如下：

```go
type Comment struct {
	Id         int64     `json:"id"`
	UserInfoId int64     `json:"-"` //用于一对多关系的id
	VideoId    int64     `json:"-"` //一对多，视频对评论
	User       UserInfo  `json:"user" gorm:"-"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"-"`
	CreateDate string    `json:"create_date" gorm:"-"`
}
```

很明显，为了与数据库中设计的表一一对应，在原数据的基础上加了几个字段，且在gorm屏蔽了User字段，所以service调用models层得到是Comment数据中User字段还未被填充，还需再填充这部分内容，好在由对应的UserId，故可以正确填充该字段。

为了重用以及不破坏代码的一致性，将填充逻辑写入util包内，比如以上的字段填充函数，同时前端要求的日期格式也能够按要求设置：

```go
func FillCommentListFields(comments *[]*models.Comment) error {
	size := len(*comments)
	if comments == nil || size == 0 {
		return errors.New("util.FillCommentListFields comments为空")
	}
	dao := models.NewUserInfoDAO()
	for _, v := range *comments {
		_ = dao.QueryUserInfoById(v.UserInfoId, &v.User) //填充这条评论的作者信息
		v.CreateDate = v.CreatedAt.Format("1-2")         //转为前端要求的日期格式
	}
	return nil
}
```

这里举了Comment这一个例子，其他的Video也是同理。

### is_favorite和is_follow字段的更新

每次为视频点赞都会在数据库的user_favor_videos表中加入用户的id和视频的id，很明显is_favorite字段是针对每个用户来判断的，而我所设计的数据库中的videos表也是包含这个字段的，但这个字段很明显不能直接进行复用，而是需要每次判断用户和此视频的关系来重新更新。

这个更新过程放入util包的填充函数中即可，为了点赞过程的迅速响应，我采取了nosql的方式存储了这个点赞的映射，也就是userId和videoId的映射，也就是用nosql代替了这个中间表的功效。

具体代码逻辑在cache包内。

### 视频的保存和封面的切片

#### 视频的保存

在本地建立static文件夹存储视频和封面图片。

具体逻辑如下：

1. 检查视频格式
2. 根据userId和该作者发布的视频数量产生唯一的名称，如id为1的用户发布了0个视频，那么本次上传的名称为1-0.mp4
3. 截取第一帧画面作为封面
4. 保存视频基本信息到数据库（包括视频链接和封面链接

#### 封面的截取

使用ffmpeg调用命令行对视频进行截取。

设计ffmpeg请求类Video2Image，通过对它内部的参数设置来构造对应的命令行字符串。具体请看util包内的ffmpeg.go的实现。

由于我设计的命令请求字符串是直接的一行字符串，而go语言exec包里面的Command函数执行所需的仅仅是一个个参数。

所以此处我想到用cgo直接调用 system(args)来解决。

代码如下：

```go
//#include <stdlib.h>
//int startCmd(const char* cmd){
//	  return system(cmd);
//}
import "C"

func (v *Video2Image) ExecCommand(cmd string) error {
	if v.debug {
		log.Println(cmd)
	}
	cCmd := C.CString(cmd)
	defer C.free(unsafe.Pointer(cCmd))
	status := C.startCmd(cCmd)
	if status != 0 {
		return errors.New("视频切截图失败")
	}
	return nil
}
```

## 可改进的地方

1. 写到后面发现很多mysql的数据可以用redis优化。
2. 很多执行逻辑可以通过并行优化。
3. 路由分组可以更为详实。
4. ...

## 项目运行

> 本项目运行不需要手动建表，项目启动后会自动建表。

**运行所需环境**：

* mysql 5.7及以上
* redis 5.0.14及以上
* ffmepg（已放入lib自带，用于对视频切片得到封面
* 需要gcc环境（主要用于cgo，windows请将mingw-w64设置到环境变量

**运行需要更改配置**：

> 进入config目录更改对应的mysql、redis、server、path信息。

* mysql：mysql相关的配置信息
* redis：redis相关配置信息
* server：当前服务器（当前启动的机器）的配置信息，用于生成对应的视频和图片链接
* path：其中ffmpeg_path为lib里的文件路径，static_source_path为本项目的static目录，这里请根据本地的绝对路径进行更改

> 完成config配置文件的更改后，需要再更改conf.go里的解析文件路径为config.toml文件的绝对路径，内容如下：
>
> ```go
> if _, err := toml.DecodeFile("你的绝对路径\\config.toml", &Info); err != nil {
> 		panic(err)
> 	}
> ```
>
>

**运行所需命令**：

```shell
cd .\byte_douyin_project\
go run main.go
```
