# Adb常用指令

adb环境配置就不说了(将sdk中的adb添加到系统环境变量中)。

## 1.查看连接的设备：

adb devices

结果：List of devices attached

emulator-5554 device

SH0A6PL00243 device

## 2.安装apk：

adb install D:\app-debug.apk

如果有多个device，如第一个命令情况，则这样使用：通过adb -s cmd向设备发送adb命令，

adb -s SH0A6PL00243 install D:\app-debug.apk

执行其他命令多设备情况也是如此

## 3.保留数据和缓存文件，重新安装apk，也称强制安装：

adb install -r D:\app-debug.apk

## 4.安装apk到sd卡：

adb install -s D:\app-debug.apk

## 5.直接卸载：

adb uninstall

## 6.卸载 app 但保留数据和缓存文件:

adb uninstall -k

## 7.列出手机装的所有app的包名:

adb shell pm list packages

## 8.列出系统应用的所有包名:

adb shell pm list packages -s

## 9.清除应用数据与缓存:

adb shell pm clear

## 10.启动应用：

adb shell am start -n com.mvvm.demo/.ui.SplashActivity

## 11.强制停止应用:

adb shell am force-stop

## 12.查看日志：

adb logcat

加上过滤会更丝滑：

adb logcat | grep -i "^E.MyApp"

含义： grep白名单模式，“xxx” 匹配的tag，

"^Exxx" 是Error级别，"^.MyApp"结尾匹配MyApp的tag。-i 忽略大小写。

grep -v黑名单模式，其他同上。

比如：

adb logcat | grep -v "^..MyApp|^..MyActivity" adb logcat | grep -vE "^..MyApp|^..MyActivity" #后面第二个使用 egrep 无须转义符

## 13.重启：

adb reboot

## 14.获取序列号：

$adb get-serialno

emulator-5554

## 15.查看设备型号：

$adb shell getprop ro.product.model

Android SDK built for x86

## 16.查看 Android 系统版本：

$adb shell getprop ro.build.version.release

7.0

## 17.查看屏幕分辨率：

$adb shell wm size

Physical size: 480x800

## 18.查看屏幕密度:

$adb shell wm density

Physical density: 480

## 19.从电脑上传文件到手机：

adb push <本地路径> <远程路径>

## 20.把手机上的文件复制到电脑:

adb pull <远程路径> <本地路径>

## 21. 解决adb连接不稳定:

adb start/kill server

掉线时, 可以先kill-server, 然后start-server来确保Server进程启动. 往往可以解决问题。

## 22.截屏：

adb shell screencap /sdcard/screen1.png

## 23.查看很多系统信息(window,activity,stack,wifi等信息)：

adb shell dumpsys -h

比较实用的一个命令是，

adb shell dumpsys meminfo 包名 -d，

可以查看出当前项目是否发生内存泄漏，

实用方法，进入app，然后正常退出，

执行上面命令，如果有组件发生被引用数不为零，就是泄漏，可以dump内存片段文件hprof分析。至于完整查找定位内存泄漏问题，我将在下个文章分析出来。

## 24.获取电量消耗信息：

获取整个设备的电量消耗信息：

adb shell dumpsys batterystats | more

获取某个apk的电量消耗信息：

adb shell dumpsys batterystats | more

可以使用一个python脚本historian.py来形成可视化html来分析耗电量问题。整理怎么使用我将会在未来的性能优化中分享。

如果想要看更多adb命令，推荐一个《一份超全超详细的 ADB 用法大全》

# 文章末尾记录一些常用出现的坑和解决办法：

## 1.adb (5037)端口被占用

解决办法：

a. 第一步，看看服务端口

adb nodaemon server

cannot bind 'tcp:5037'

b . 第二部，查找5037端口的使用情况

netstat -ano | findstr "5037"

TCP 127.0.0.1:5037 0.0.0.0:0 LISTENING 8516

TCP 127.0.0.1:5037 127.0.0.1:59163 TIME_WAIT 0

TCP 127.0.0.1:5037 127.0.0.1:59164 TIME_WAIT 0

TCP 127.0.0.1:5037 127.0.0.1:59167 TIME_WAIT 0

c. 第三步，查看线程id8586的进程名称

tasklist | findstr "8516"

sjk_daemon 8516 Console 13,071 K

哦，原来是sjk_daemon进程占了adb的端口。

d. 第四部，查看进程名对应的进程id

tasklist

Image Name PID Session Name Session# Mem Usage

========================= ======== ================ =========== ============

System Idle Process 0 Services 0 24 K

System 4 Services 0 1,128 K

sjk_daemon 963 Console 1 3,071 K

tasklist.exe 1260 Console 1 5,856 K

e， 第五步， 将这个进程kill掉：

taskkill /f /pid 963

如果这个命令提示无权限，那么，可以去windows的“任务管理器”中“进程”那个窗口，找到这个进程，将它杀掉。

再运行adb devices，就没有问题了。

adb devices

9dk7f482396a371j device

或者重启：

adb kill-server

adb start-server