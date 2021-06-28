# Golang

## 1. Go 语言可以做什么

- 区块链研发工程师
- Go 服务器端
- 游戏、软件工程师
- Golang 分布式
- 云计算

## 2. Go 语言的特点

- 既有静态语言的安全和性能，又有动态语言的开发维护的高效率
- 保留了 C 语言的编译方式和弱化指针
- 垃圾回收
- 天然并发
    - 从语言层面支持
    - goroutine 协程
    - 基于 CPS（Communicating Sequential Processes）并发模型
- 管道通信机制（通过 Channel 进行通信）
- 支持返回多个值
- 切片 slice，延迟执行 defer

## 3. Go 的数据类型

__基本数据类型__ 

- 数值型
    - 整数：int，int8，int16，int32(rune 表示一个 Unicode 的码点)，int64，uint，uint8，uint16，uint32，uint64，byte(uint8)
    - 浮点：float32，float64
    - 复数：complex32，complex64
- 字符型（没有专门的字符类型，通过 byte 保存单个字母）
- 布尔型：bool
- 字符串：string（官方归类到基本类型）

__派生类型__ 

- 指针 Pointer
- 数组
- 结构体 struct
- 管道 Channel
- 函数
- 切片 slice
- 接口 interface
- map
- ......


