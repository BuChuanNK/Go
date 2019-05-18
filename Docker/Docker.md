# 自己动手写 Docker

## 1. 容器与开发语言

### 1.1 Docker
#### 1.1.1 简介
Docker 是一个开源工具，可以将应用打包成一个标准格式的镜像，并且以容器的方式运行。 
Docker 容器将一系列软件包装在一个完整的文件系统中，文件系统包含应用程序运行所需要的一切：代码、运行时工具、系统工具、系统依赖，几乎有任何可以安装在服务器上的东西。 
#### 1.1.2 容器和虚拟机比较
容器和虚拟机同样有着资源隔离和分配的优点，但容器更加便携和高效。
虚拟机需要各自的 GUEST OS, 容器包含用户的程序和所有的依赖, 但是容器之间是共享 Kernel 的。 各个容器在 Host 上互相隔离, 并且在用户态下运行. 
#### 1.1.3 容器加速开发效率
#### 1.1.4 利用容器合作开发
#### 1.1.5 利用容器快速扩容
#### 1.1.6 安装使用 Docker

### 1.2 Go

## 2. 基础技术

### 2.1 Linux Namespace 介绍
**Docker 是一个使用了 Linux Namespace 和 Cgroups 的虚拟化工具。**
#### 2.1.1 概念
Linux Namespace 是 Kernel 的一个功能，可以隔离一系列的系统资源，比如PID, User ID, Network 等。
使用 Linux Namespace 可以做到 UID 级别的隔离, 以 UID 为 n 的用户, 虚拟化出一个 Namespace, 在这个 Namespace 中, 用户是具有 root 权限的. 
PID也是可以通过 Linux Namespace 虚拟化的。从用户的角度来看，每一个命名空间就是一个独立的 Linux, 有自己的 init 进程 (PID 为 1), 其他进程的 PID 依次递增. 
Namespace 的 API 主要使用3个系统调用:
1. clone() 创建新进程. 根据系统调用参数来判断哪些类型的 Namespace 被创建, 而且它们的子进程也会被包含到这些 Namespace . 
2. unshare() 将进程移出某个 Namespace 中. 
3. setns() 将进程加入到 Namespace 中. 

#### 2.1.2 UTS Namespace
UTS Namespace 主要用来隔离 nodename 和 domainname 两个系统标识. 在 UTS Namespace 里面, 每个 Namespace 允许有自己的 hostname. 
