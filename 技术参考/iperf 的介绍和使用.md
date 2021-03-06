# iperf 的介绍和使用

发表于 2017-06-06 | 分类于 [知识整理/总结 ](http://www.enkichen.com/categories/知识整理-总结/)| | 阅读次数 4091

`iperf` 是一个网络性能测试工具，做服务开发或者测试的同学，接触的可能比较多。因为最近有用到这个工具，并且这个工具做的非常不错，这里记录一下工具的使用方法。`iperf` 是个开源并且跨平台的软件，代码托管在 [GitHub](https://github.com/esnet/iperf) 上，可以从 [Releases](https://github.com/esnet/iperf/releases) 找到各个发行版本，也可以去 [官网](https://iperf.fr/iperf-download.php) 下载各个平台的版本。 使用 `iperf` 时，需要分别运行服务端和客户端，在测试是最好保证两个端的软件版本一致，这样会免去一些没必要的麻烦。

下载好后，可以先在本机做一个简单的回环测试，结果如下：

```
# 运行服务端
$ ./iperf -s
------------------------------------------------------------
Server listening on TCP port 5001
TCP window size:  128 KByte (default)
------------------------------------------------------------
[  4] local 127.0.0.1 port 5001 connected with 127.0.0.1 port 54817
[ ID] Interval       Transfer     Bandwidth
[  4]  0.0-10.0 sec  35.1 GBytes  30.1 Gbits/sec

# 运行客户端
$ ./iperf -c 127.0.0.1
------------------------------------------------------------
Client connecting to 127.0.0.1, TCP port 5001
TCP window size:  144 KByte (default)
------------------------------------------------------------
[  4] local 127.0.0.1 port 54817 connected with 127.0.0.1 port 5001
[ ID] Interval       Transfer     Bandwidth
[  4]  0.0-10.0 sec  35.1 GBytes  30.1 Gbits/sec
```

默认情况下，会使用 TCP 连接，绑定在 5001 端口上，可以从上述结果看到，当前本机的带宽为 `30.1 Gbits/sec` 。

### 主要参数信息

**适用于 服务端/客户端**

```
-f 指定数据显示格式 [k|m|K|M] 分别表示 Kbits、Mbits、KBytes、MBytes，默认是 Mbits
-l 读写缓冲区的大小，默认是 8K
-u 使用 udp 协议
-i 以秒为单位统计带宽值
-m 显示最大的 TCP 数据段大小
-p 指定服务端或者客户端的端口号
-w 指定 TCP 窗口大小
-B 绑定道指定的主机地址或接口
-C 兼容旧版本
-M 设置 TCP 数据包的最大 MTU 值
-V 传输 IPV6 包
```

**适用于 服务端**

```
-s 以服务器模式启动
-U 单线程 UDP 模式
-D 以守护进程模式运行
```

**适用于 客服端**

```
-c 以客户端模式运行，并指定服务端的地址
-b 指定客户端通过 UDP 协议发送信息的带宽，默认为 1Mbit/s
-d 同时进行双向传输测试
-n 指定传输的字节数
-r 单独进行双向传输测试
-t 指定 iperf 测试的时间，默认 10s
-F 指定要传输的文件
-L 指定一个端口，服务利用这端口与客户端连接
-P 指定客户端到服务器的连接数，默认是 1
-T 指定 ttl 值
```

> - 用 -u 参数来指定使用 UDP 协议，需要在 -p 参数之前指定
> - 测试之前确保防火墙为关闭状态

### 网络性能测试

**TCP 协议测试带宽**

```
# 运行服务端
$ iperf -s

# 运行客户端
$ iperf -c 172.18.142.62 -i 1 -t 10
------------------------------------------------------------
Client connecting to 172.18.142.62, TCP port 5001
TCP window size:  129 KByte (default)
------------------------------------------------------------
[  4] local 172.18.98.209 port 57809 connected with 172.18.142.62 port 28756
[ ID] Interval       Transfer     Bandwidth
[  4]  0.0- 1.0 sec   384 KBytes  3.15 Mbits/sec
[  4]  1.0- 2.0 sec   256 KBytes  2.10 Mbits/sec
[  4]  2.0- 3.0 sec   256 KBytes  2.10 Mbits/sec
[  4]  3.0- 4.0 sec   256 KBytes  2.10 Mbits/sec
[  4]  4.0- 5.0 sec   512 KBytes  4.19 Mbits/sec
[  4]  5.0- 6.0 sec  1.12 MBytes  9.44 Mbits/sec
[  4]  6.0- 7.0 sec  1.12 MBytes  9.44 Mbits/sec
[  4]  7.0- 8.0 sec  1.12 MBytes  9.44 Mbits/sec
[  4]  8.0- 9.0 sec  1.25 MBytes  10.5 Mbits/sec
[  4]  9.0-10.0 sec  1.12 MBytes  9.44 Mbits/sec
[  4]  0.0-10.1 sec  7.50 MBytes  6.25 Mbits/sec
```

使用 TCP 协议进行测试时，需要注意的就是 TCP 窗口大小，可以使用 `-w` 参数指定，网络通道的容量 `capacity = bandwidth * round-trip time`，而理论 TCP 窗口大小就是网络通道的容量。例如，网络带宽为 `40Mbit/s`，回环路径消耗时间是 2ms，那么 TCP 的窗口大小不小于 `40Mbit/s×2ms = 80kbit = 10Kbytes` 。

**UDP 协议测试带宽**

```
# 运行服务端
$ iperf -u -s

# 运行客户端
$ iperf -c 172.18.142.62 -u -i 1 -t 10 -b 30M
------------------------------------------------------------
Client connecting to 172.18.142.62, UDP port 5001
Sending 1470 byte datagrams
UDP buffer size: 9.00 KByte (default)
------------------------------------------------------------
[  4] local 172.18.98.209 port 53220 connected with 172.18.142.62 port 28756
[ ID] Interval       Transfer     Bandwidth
[  4]  0.0- 1.0 sec  3.58 MBytes  30.0 Mbits/sec
[  4]  1.0- 2.0 sec  3.58 MBytes  30.0 Mbits/sec
[  4]  2.0- 3.0 sec  3.58 MBytes  30.0 Mbits/sec
[  4]  3.0- 4.0 sec  3.58 MBytes  30.0 Mbits/sec
[  4]  4.0- 5.0 sec  3.58 MBytes  30.0 Mbits/sec
[  4]  5.0- 6.0 sec  3.57 MBytes  30.0 Mbits/sec
[  4]  6.0- 7.0 sec  3.58 MBytes  30.0 Mbits/sec
[  4]  7.0- 8.0 sec  3.58 MBytes  30.0 Mbits/sec
[  4]  8.0- 9.0 sec  3.58 MBytes  30.0 Mbits/sec
[  4]  9.0-10.0 sec  3.58 MBytes  30.0 Mbits/sec
[  4]  0.0-10.0 sec  35.8 MBytes  30.0 Mbits/sec
[  4] Sent 25511 datagrams
[  4] Server Report:
[  4]  0.0-11.6 sec  13.6 MBytes  9.83 Mbits/sec   1.971 ms 15786/25497 (62%)
[  4]  0.0-11.6 sec  140 datagrams received out-of-order
```

上述命令指定了客户端以 `30Mbit/s` 速度发送数据，由于 UDP 协议是无连接不可靠的，并且只管发包，不确保包在服务端是否接收到，所以需要查看服务报告才能确定当前网络性能数据。如果在不知道当前网络带宽的情况下，需要不断的调整参数值，并且查看丢包率，来确定当前网络性能情况。如果你当前是远程登录到服务器上进行测试的，可以从小到大的方式进行测试，否则很容易导致服务当前带宽被占满。

### 参考资料

- [使用 iperf 测试网络性能](http://blog.csdn.net/evenness/article/details/7371845)
- [iperf 测试带宽](http://www.52os.net/articles/iperf-check-bandwidth.html)