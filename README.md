<div align="center">

<h1 align="center">go-bind-mp</h1>

[![Auth](https://img.shields.io/badge/Auth-eryajf-ff69b4)](https://github.com/eryajf)
[![Go Version](https://img.shields.io/github/go-mod/go-version/eryajf/xirang)](https://github.com/eryajf/xirang)
[![Gin Version](https://img.shields.io/badge/Gin-1.6.3-brightgreen)](https://github.com/eryajf/xirang)
[![Gorm Version](https://img.shields.io/badge/Gorm-1.20.12-brightgreen)](https://github.com/eryajf/xirang)
[![GitHub Issues](https://img.shields.io/github/issues/eryajf/xirang.svg)](https://github.com/eryajf/xirang/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/eryajf/xirang)](https://github.com/eryajf/xirang/pulls)
[![GitHub Pull Requests](https://img.shields.io/github/stars/eryajf/xirang)](https://github.com/eryajf/xirang/stargazers)
[![HitCount](https://views.whatilearened.today/views/github/eryajf/xirang.svg)](https://github.com/eryajf/xirang)
[![GitHub license](https://img.shields.io/github/license/eryajf/xirang)](https://github.com/eryajf/xirang/blob/main/LICENSE)

<p> 🐉 简单好用，不缠不绕，直接上手的go-backend框架 </p>

<img src="https://camo.githubusercontent.com/82291b0fe831bfc6781e07fc5090cbd0a8b912bb8b8d4fec0696c881834f81ac/68747470733a2f2f70726f626f742e6d656469612f394575424971676170492e676966" width="800"  height="3">
</div><br>

<p align="center">
  <a href="" rel="noopener">
 <img src="https://cdn.staticaly.com/gh/eryajf/tu/main/img/image_20220826_101156.png" alt="Project logo"></a>
</p>


## 🥸 项目介绍

`go-bind-mp` 是一个非常简单的 `gin+gorm` 框架的基础架构，你只需要修改简单的代码，即可开始上手编写你的接口。

只需要根据情况修改配置`config.yml`，然后配置里边的数据库配置信息，即可开始开发。

数据表会自动创建，也可以通过docs下的sql自行创建。

## 👨‍💻 项目地址

| 分类 |                        GitHub                       
| :--: | :--------------------------------------------------
| 后端 |  https://github.com/yin-zt/go-bind-mp
|
## 📖 目录结构

```
go-bind-mp
├── cmd ----------------程序启动脚本
    ├── doce----------------介绍文档服务【待完成】
    ├── server---------------主服务
        ├── config.yml  ---------------- 主配置文件
        ├── main.go---------------- 主服务启动脚本
├── pkg---------------- 逻辑代码目录
    ├── config----------------配置文件读取
    ├── controller------------控制层
    ├── middleware------------中间件
    ├── model-----------------对象定义
    ├── routers---------------路由
    ├── service---------------服务层
        ├── logic ---------------逻辑层
        ├── isql --------------- 数据库交互层
    ├── util----------------一些公共组件与工具


```


## 🚀 快速开始

go-bind-mp项目的基础依赖项只有MySQL，本地准备好mysql服务之后，就可以启动项目，进行调试。

### 拉取代码

```sh
# 后端代码
$ git clone https://github.com/yin-zt/go-bind-mp.git

### 更改配置

```sh
# 修改后端配置
$ cd go-bind-mp
# 文件路径 config.yml, 根据自己本地的情况，调整数据库等配置信息。
$ vim config.yml
```

### 启动服务

```sh
# 启动后端
$ cd go-bind-mp
$ go mod tidy
$ go run main.go

```

